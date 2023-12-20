package convertor

import (
	"fmt"
	"github.com/Metchain/MetblockD/app/rpc/rpccontext"
	"github.com/Metchain/MetblockD/appmessage"
	"github.com/Metchain/MetblockD/blockchain"
	"github.com/Metchain/MetblockD/external"
	"github.com/pkg/errors"
	"time"
)

func RPCBlockToDomainBlock(b *appmessage.RPCBlock, context *rpccontext.Context) ([32]byte, error, float64) {

	nb := new(external.TempBlock)

	bc := context
	var oldtimestamp int64
	nb.Bits = uint64(b.Header.Bits)
	//log.Println(bc.MiniBlock)

	nb.Nonce = b.Header.Nonce
	nb.Height = bc.LastRPCBlock.Header.Blockheight() + 1
	nb.PreviousHash = *bc.LastRPCBlock.Header.BlockHash().ByteArray()
	oldtimestamp = bc.LastRPCBlock.Header.TimeInMilliseconds()

	nb.Timestamp = time.Now().UnixMilli()

	oldntime := oldtimestamp + 6000

	if oldtimestamp > nb.Timestamp || oldntime >= nb.Timestamp {
		//log.Printf("This is trigered")
		nb.Bits = nb.Bits + 1
		return [32]byte{}, errors.New("Block is too early"), 0
	} else if nb.Timestamp-oldntime > 4000 && nb.Bits > 569658475 {
		nb.Bits = nb.Bits - 1
	} else {
		nb.Bits = nb.Bits + 1
	}
	/*
		ntx := bc.Txpool
		nfttx := bc.NFTPool

		nftre := []*Transaction{}
		if len(bc.NFTStake) >= 1 {
			nftre = bc.CheckNFTRewards()

		}
		bc.ClearTransactionPool()*/
	transactions := make([]*external.TempTransaction, 0)

	//Get Last Megablock

	nb.Megablock = blockchain.StringTo32Byte(b.Header.Parents[0].ParentHashes[0])

	// Get last Metblock
	nb.Metblock = blockchain.StringTo32Byte(b.Header.Parents[0].ParentHashes[1])

	nb.CurrentHash = blockchain.Hash([]byte(fmt.Sprintf("%x", nb.PreviousHash)), nb.Timestamp, nb.Nonce)

	transactions = append(transactions, blockchain.NewTransactionMiner(blockchain.MINING_SENDER, b.Header.UTXOCommitment, blockchain.MINING_REWARD))
	Reward := blockchain.MINING_REWARD

	// ADD transaction pool transactions and nft transactions here.

	nb.Transactions = transactions

	block, err := context.Domain.MiningManager().AddMiningBlock(nb)
	if err != nil {
		return [32]byte{}, err, 0
	}
	//Update the latest rpc block

	/*if len(nfttx) >= 1 {

		bc.StakedNFT(nb, nfttx)

	}
	bc.WalletToDB(db, nb)*/

	return block.CurrentHash, nil, Reward
}
