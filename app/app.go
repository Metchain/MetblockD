package app

import (
	"flag"
	"github.com/Metchain/Metblock/app/protocol"
	"github.com/Metchain/Metblock/blockchain"
	"github.com/Metchain/Metblock/blockchain/consensus"
	"github.com/Metchain/Metblock/blockchain/mempool"
	"github.com/Metchain/Metblock/blockchainserver"
	"github.com/Metchain/Metblock/commanager"
	"github.com/Metchain/Metblock/db/database"
	"github.com/Metchain/Metblock/domain"
	"github.com/Metchain/Metblock/mconfig/infraconfig"
	"github.com/Metchain/Metblock/mconfig/profiling"
	"github.com/Metchain/Metblock/network/addressmanager"
	"github.com/Metchain/Metblock/protoserver"
	"github.com/Metchain/Metblock/utils/signal"
	"github.com/Metchain/Metblock/version"
	"os"
	"time"
)

func (app *metchainApp) main(startedChan chan<- struct{}) error {
	// Get a channel that will be closed when a shutdown signal has been
	// triggered either from an OS signal such as SIGINT (Ctrl+C) or from
	// another subsystem such as the RPC server.
	interrupt := signal.InterruptListener()
	defer log.Info("Shutdown complete")

	// Show version at startup.
	log.Infof("Version %s", version.Version())

	// Enable http profiling server if requested.
	if app.cfg.Profile != "" {
		profiling.Start(app.cfg.Profile, log)
	}
	profiling.TrackHeap(app.cfg.AppDir, log)

	// Return now if an interrupt signal was triggered.
	if signal.InterruptRequested(interrupt) {
		return nil
	}

	if app.cfg.ResetDatabase {
		err := removeDatabase(app.cfg)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	// Open the database
	databaseContext, err := openDB(app.cfg)
	if err != nil {
		log.Errorf("Loading database failed: %+v", err)
		return err
	}

	defer func() {
		log.Infof("Gracefully shutting down the database...")
		err := databaseContext.Close()
		if err != nil {
			log.Errorf("Failed to close the database: %s", err)
		}
	}()

	// Return now if an interrupt signal was triggered.
	if signal.InterruptRequested(interrupt) {
		return nil
	}

	// Create componentManager and start it.
	componentManager := app.NewComponentManager(app.cfg, databaseContext, interrupt)

	defer func() {
		log.Infof("Gracefully shutting down metchaind...")

		shutdownDone := make(chan struct{})
		go func() {

			shutdownDone <- struct{}{}
		}()

		const shutdownTimeout = 2 * time.Minute

		select {
		case <-shutdownDone:
		case <-time.After(shutdownTimeout):
			log.Criticalf("Graceful shutdown timed out %s. Terminating...", shutdownTimeout)
		}
		log.Infof("Metchaind shutdown complete")
	}()
	componentManager.Start()

	if startedChan != nil {
		startedChan <- struct{}{}
	}

	<-interrupt
	if startedChan != nil {
		startedChan <- struct{}{}
	}

	<-interrupt
	return nil
}

func (MetApp *metchainApp) NewComponentManager(cfg *infraconfig.Config, db database.Database, interrupt chan<- struct{}) *ComponentManager {

	consensusConfig := consensus.Config{
		Params:                          *cfg.ActiveNetParams,
		IsArchival:                      cfg.IsArchivalNode,
		EnableSanityCheckPruningUTXOSet: cfg.EnableSanityCheckPruningUTXOSet,
	}
	mempoolConfig := mempool.DefaultConfig(&consensusConfig.Params)
	mempoolConfig.MaximumOrphanTransactionCount = cfg.MaxOrphanTxs
	mempoolConfig.MinimumRelayTransactionFee = cfg.MinRelayTxFee
	addressManager, err := addressmanager.New(addressmanager.NewConfig(cfg), db)
	if err != nil {
		log.Infof("Shutting down : ", err)

		os.Exit(112)
	}
	domain, err := domain.New(&consensusConfig, mempoolConfig, db)
	if err != nil {
		log.Criticalf("Shutting down :", err)

		os.Exit(112)
	}
	log.Info(domain)

	bc := blockchain.GenesisGenrate()
	msg, nerr := bc.VerifyGenesis()
	if nerr {
		log.Infof("Shutting down : ", msg)

		os.Exit(112)
	}

	log.Infof(msg)
	consensus.Sync(bc.GensisCompile(), db, cfg)

	mbc := blockchain.Start(db)

	port := flag.Uint("port", 5000, "TCP Port Number for Blockchain Server")
	flag.Parse()
	app := blockchainserver.NewBlockchainServer(uint16(*port), mbc)
	go func() {

		app.Run()

	}()

	netAdapter, err := protoserver.NewNetAdapter(cfg, mbc)
	if err != nil {
		log.Infof("Shutting down : ", msg)

		os.Exit(112)
	}

	network, err := commanager.New(cfg, netAdapter, addressManager)
	if err != nil {
		log.Infof("Shutting down : ", err)

		os.Exit(112)
	}

	protocolManager, err := protocol.NewManager(cfg, netAdapter, addressManager, network)

	rpcManager := setupRPC(cfg, domain, netAdapter, protocolManager, network, addressManager, interrupt)

	return &ComponentManager{
		cfg:               cfg,
		protocolManager:   protocolManager,
		rpcManager:        rpcManager,
		connectionManager: network,
		netAdapter:        netAdapter,
		addressManager:    addressManager,
	}

}
