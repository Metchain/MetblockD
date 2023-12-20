package blockchain

import (
	"github.com/Metchain/MetblockD/db/database"
	"github.com/Metchain/MetblockD/domain"
	pb "github.com/Metchain/MetblockD/protoserver/grpcserver/protowire"

	"sync"
)

type Blockchain struct {
	Txpool []*Transaction

	Unstake []*int64
	//Add on 10/4/2023 Update for recording the wallets created. Not to be changed
	Wallets        []*WalletCreated
	MiniBlock      []string
	MetBlock       []string
	MegaBlock      []string
	MiniBlockCount int
	MetBlockCount  int
	MegaBlockCount int

	NFTPool  []*NFTStaking
	NFTStake []*pb.StakeNFT
	//added for localnodes
	Port uint16
	//added for mining
	Mux sync.Mutex

	MuxNeighbors sync.Mutex

	LastRPCBlock *LastRPCBlock
	*domain.Metchain
}

type LastRPCBlock struct {
	Height       uint64
	Timestamp    int64
	Nonce        uint64
	PreviousHash [32]byte //As per the Hash size
	Megablock    [32]byte
	Metblock     [32]byte
	Transactions []*Transaction
	CurrentHash  [32]byte
	Bits         uint64
}

type WalletCreated struct {
	WalletAddress string
	BlockHash     string
	LockHash      string
}

func Start(db database.Database) *Blockchain {

	bc := &Blockchain{
		Metchain: &domain.Metchain{
			Db: db,
		},
	}

	err := bc.LastMiniBlockRPC()
	if err != nil {
		log.Criticalf("Error while processing the Latest RPC Block : %s", err)
	}
	bc.NFTStake = bc.GetStakedNFT()
	return bc
}
