package app

import (
	"flag"
	"github.com/Metchain/Metblock/blockchain"
	"github.com/Metchain/Metblock/blockchain/consensus"
	"github.com/Metchain/Metblock/blockchainserver"
	"github.com/Metchain/Metblock/db/database"
	"github.com/Metchain/Metblock/mconfig/infraconfig"
	"github.com/Metchain/Metblock/mconfig/profiling"
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
	NewComponentManager(app.cfg, databaseContext, interrupt)

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

	if startedChan != nil {
		startedChan <- struct{}{}
	}

	<-interrupt
	return nil
}

func NewComponentManager(cfg *infraconfig.Config, db database.Database, interrupt chan<- struct{}) {
	bc := blockchain.GenesisGenrate()
	msg, err := bc.VerifyGenesis()
	if err {
		log.Infof(msg)
		log.Infof("Shutting down")

		os.Exit(112)
	}
	log.Infof(msg)
	metch := consensus.Sync(bc.GensisCompile(), db, cfg)

	mbc := blockchain.Start(metch)

	port := flag.Uint("port", 5000, "TCP Port Number for Blockchain Server")
	flag.Parse()
	app := blockchainserver.NewBlockchainServer(uint16(*port), metch, mbc)
	go func() {

		app.Run()

	}()
	blockchain.NewLastMiniBlock(metch)
	protoserver.NewNetAdapter(cfg, metch, mbc)

	/*go func() {
		protoserver.P2PServer(metch, mbc)
	}()

	protoserver.RPCServer(metch, mbc)*/

	blockchain.GetBlockTemplateBC(metch, "")
}
