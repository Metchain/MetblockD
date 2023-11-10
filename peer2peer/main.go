package main

import (
	"context"
	"github.com/Metchain/Metblock/mconfig"
	pb "github.com/Metchain/Metblock/proto"
	"github.com/Metchain/Metblock/utils"
	"github.com/btcsuite/goleveldb/leveldb"
	"github.com/btcsuite/goleveldb/leveldb/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Peer2peer struct {
	Domains []string
	DbCon   *leveldb.DB
}

func main() {
	Peer := VerifyConcensusDomain()
	Peer.PeerDomain()
}

func VerifyConcensusDomain() *Peer2peer {
	Peer2peer := new(Peer2peer)
	DomainConfigure := []string{"wallet.metchain.tech"}
	ConcensusDomain := []string{}

	for _, val := range DomainConfigure {
		ConcensusDomain = append(ConcensusDomain, mconfig.VerifyDomainIp(val))
	}
	Peer2peer.Domains = ConcensusDomain

	Peer2peer.DbCon = GetDB()
	return Peer2peer
}

func (p2p *Peer2peer) PeerDomain() {
	log.Println(p2p.Domains)
	for _, domain := range p2p.Domains {
		//go func() {

		p2p.DomainDail(domain)
		//}()

	}

}

func DomainToDial(domain string) string {
	return domain + ":14031"
}

func (p2p *Peer2peer) DomainDail(domain string) {
	//conn, err := grpc.Dial(DomainToDial(domain), grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("154.53.46.191:14031", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	client := pb.NewP2PClient(conn)

	// Add verification here.
	i := 0
	for {
		i = i + 1
		stream, _ := client.MessageStream(context.Background())

		payload := new(pb.MetchainMessage_P2PBlockWithTrustedDataRequestMessage)
		payload.P2PBlockWithTrustedDataRequestMessage = p2p.ConvertGroupToSecureInfo(p2p.ConvertBlockToGroupBlock(p2p.ConvertBlockToProto()))
		req := &pb.MetchainMessage{
			Payload: payload,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
			//return err
		} else {
			log.Println("Sent Request")
		}
		msg, err := stream.Recv()
		msg.GetP2PLatestBlockHeightRequest()
		if err != nil {
			log.Fatalf("Error while recieving %v", err)
		} else {
			//log.Println("Message recived: ", msg)
		}
		println(i)
		//ime.Sleep(6000 * time.Millisecond)

	}

	//

}

func GetDB() *leveldb.DB {
	d := mconfig.GetDBDirTestnet()
	db, err := leveldb.OpenFile(d, nil)
	if err != nil {
		log.Fatalf("Error while accessing the blockchain db : %v", err)
	}
	return db
}

func (p2p *Peer2peer) GetLastBlock() []byte {
	db := p2p.DbCon
	value := []byte{}
	iter := db.NewIterator(util.BytesPrefix([]byte("bk-")), nil)
	/**/

	ok := iter.Last()
	if ok {
		value = iter.Value()

	}
	ok = iter.Next()
	if ok {
		value = iter.Value()

	}

	iter.Release()
	return value
}

func (p2p *Peer2peer) ConvertBlockToProto() *pb.BlockInfo {
	lb := p2p.GetLastBlock()
	n := new(MBlock)
	n.UnmarshalJSON(lb)

	ProtoBlock := new(pb.BlockInfo)

	ProtoBlock.Blockhash = utils.ConvertDomainByteToProtoByte(n.CurrentHash)
	ProtoBlock.PreviousHash = utils.ConvertDomainByteToProtoByte(n.PreviousHash)
	ProtoBlock.Megablock = utils.ConvertDomainByteToProtoByte(n.Megablock)
	ProtoBlock.Metblock = utils.ConvertDomainByteToProtoByte(n.Metblock)
	ProtoBlock.Nonce = n.Nonce
	ProtoBlock.Blockheight = n.Height
	ProtoBlock.CurrentHash = utils.ConvertDomainByteToProtoByte(n.CurrentHash)
	ProtoBlock.Timestamp = n.Timestamp
	ProtoBlock.Bits = n.Bits
	for _, tx := range n.Transactions {
		ntx := new(pb.BlockTransactions)
		ntx.Sender = tx.senderBlockchainAddress
		ntx.Recipient = tx.recipientBlockchainAddress
		ntx.Txtype = utils.ConvertDomainInt8ToProtoInt64(tx.txtype)
		ntx.Value = utils.ConvertValToString(tx.value)
		ntx.Txhash = utils.ConvertDomainByteToProtoByte(tx.txhash)
		ntx.Timestamp = tx.timestamp
		ntx.Txstatus = utils.ConvertDomainInt8ToProtoInt64(tx.txstatus)
	}
	return ProtoBlock
}

func (p2p *Peer2peer) ConvertBlockToGroupBlock(b *pb.BlockInfo) []*pb.BlockInfo {
	block := make([]*pb.BlockInfo, 0)
	block = append(block, b)
	return block
}

func (p2p *Peer2peer) ConvertGroupToSecureInfo(bks []*pb.BlockInfo) *pb.P2PBlockWithTrustedDataRequestMessage {

	return &pb.P2PBlockWithTrustedDataRequestMessage{
		BlockInfo: bks,
	}

}
