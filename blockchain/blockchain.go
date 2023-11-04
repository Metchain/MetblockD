package blockchain

import (
	"github.com/Metchain/Metblock/domain"
	pb "github.com/Metchain/Metblock/proto"
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

	MuxNeighbors    sync.Mutex
	LastBlockTime   int64
	LastBlockHeight uint64
	LastBlockHash   [32]byte
	*domain.Metchain
}

type WalletCreated struct {
	WalletAddress string
	BlockHash     string
	LockHash      string
}

func Start(mc *domain.Metchain) *Blockchain {

	bc := new(Blockchain)
	bc.Metchain = mc
	bc.LastBlockHeight, bc.LastBlockHash, bc.LastBlockTime = LastMiniBlockRPC(mc.Dbcon)
	bc.NFTStake = bc.GetStakedNFT()
	return bc
}
