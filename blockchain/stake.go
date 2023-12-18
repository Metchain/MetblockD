package blockchain

import (
	"fmt"
	"github.com/Metchain/Metblock/db/database"
	"github.com/Metchain/Metblock/mconfig"
	pb "github.com/Metchain/Metblock/protoserver/grpcserver/protowire"
	"google.golang.org/protobuf/proto"

	"math"
	"strconv"
	"time"
)

type NFTStaking struct {
	NFT           int
	MetCoin       float64
	WalletAddress string
	LockTime      int64
	UnLockTime    int64
	Status        bool
	LockingTXHash []byte
}

func (bc *Blockchain) StakeNFTMet(sender string, seed string, value string, nft string, stakeunlock string) bool {
	if bc.verifyNFt(sender, nft) {
		var dur string
		if stakeunlock == "3" {
			dur = fmt.Sprintf("%s", mconfig.Lock3Month)
		} else if stakeunlock == "6" {
			dur = fmt.Sprintf("%s", mconfig.Lock6Month)
		} else if stakeunlock == "9" {
			dur = fmt.Sprintf("%s", mconfig.Lock9Month)
		} else if stakeunlock == "12" {
			dur = fmt.Sprintf("%s", mconfig.Lock12Month)

		}
		Unlocktime, _ := time.ParseDuration(dur + "m")

		f, _ := strconv.ParseFloat(value, 64)

		NFTid, _ := strconv.ParseInt(nft, 10, 64)

		err := bc.verifyBalance(f, sender)
		if err != nil {

			return false
		} else {
			Stake := &NFTStaking{
				MetCoin:       math.Abs(f),
				WalletAddress: sender,
				LockTime:      time.Now().UnixMilli(),
				UnLockTime:    time.Now().Add(Unlocktime).UnixMilli(),
				NFT:           int(NFTid),
			}
			bc.NFTPool = append(bc.NFTPool, Stake)
			return true
		}

	}
	return false
}
func (bc *Blockchain) verifyBalance(f float64, sender string) error {
	walletcheck := GetWalletBalnace(sender, bc.Db)
	walletinfo := new(pb.WalletBalanceRespose)
	proto.Unmarshal(walletcheck, walletinfo)
	bal, _ := strconv.ParseFloat(walletinfo.Amount, 64)
	f = math.Abs(f)
	if bal < f {
		return fmt.Errorf("Error: Ensufficent Balance")
	} else if f < mconfig.MinimumStaking {
		return fmt.Errorf("Error: Doesn't match the minimum staking requirements.")
	}
	return nil
}

var stakedkey = database.MakeBucket([]byte("layer2_nft_staked"))

func (bc *Blockchain) StakedNFT(b *MiniBlock, pool []*NFTStaking) []*Transaction {

	batch, _ := bc.Db.Begin()
	defer batch.RollbackUnlessClosed()
	txs := []*Transaction{}
	for _, z := range pool {

		if z.MetCoin >= mconfig.MinimumStaking {

			tempkey := stakedkey.Key([]byte(fmt.Sprintf("%v", z.NFT)))
			nftid := fmt.Sprintf("%v", z.NFT)
			t := NewTransaction(z.WalletAddress, STAKING_SENDER+"( MET NFT #"+nftid+")", float32(z.MetCoin))
			sn := StakedInfo(t.txhash, b, z)
			staker, err := proto.Marshal(sn)

			if err != nil {
				log.Criticalf("Error. This shouldn't happen: %v", err)

			} else {

				txs = append(txs, t)
				bc.Txpool = append(bc.Txpool, t)
				sn := new(pb.StakeNFT)

				batch.Put(tempkey, staker)
				bc.NFTStake = append(bc.NFTStake, sn)
			}

		}
	}

	err := batch.Commit()

	if err != nil {
		log.Criticalf("Error while adding staking NFT: ", err)

	}
	return txs
}

func StakedInfo(txhash [32]byte, b *MiniBlock, z *NFTStaking) *pb.StakeNFT {
	return &pb.StakeNFT{
		Txhash:      fmt.Sprintf("%x", txhash),
		Blockhash:   fmt.Sprintf("%x", b.currentHash),
		LockTime:    uint64(z.LockTime),
		UnlockTime:  uint64(z.UnLockTime),
		Blockid:     fmt.Sprintf("%v", b.height),
		NFTid:       fmt.Sprintf("%v", z.NFT),
		NFTSender:   fmt.Sprintf("%v", z.WalletAddress),
		StakeAmount: fmt.Sprintf("%v", z.MetCoin),
	}
}

func (bc *Blockchain) verifyNFt(sender string, nft string) bool {
	ni := new(pb.NFTResponse)

	m := bc.VerifyNFT(nft)
	proto.Unmarshal(m, ni)
	if ni.NFTWallet == sender {
		return true
	}
	return false
}

var nftLayerkey = database.MakeBucket([]byte("layer2_nft"))

func (bc *Blockchain) VerifyNFT(wx string) []byte {

	txs, err := bc.Db.Get(nftLayerkey.Key([]byte(wx)))
	if err != nil {
		log.Infof("Error locating the Wallet Address")
		return nil
	}

	return txs
}

var stakednftbucket = database.MakeBucket([]byte("layer2_nft_staked"))

func (bc *Blockchain) GetStakedNFT() []*pb.StakeNFT {
	i := 1
	db := bc.Db
	Staked := []*pb.StakeNFT{}
	for i <= 500 {
		m, err := db.Get(stakednftbucket.Key([]byte(fmt.Sprintf("%v", i))))
		if err == nil {

			dbnft := new(pb.StakeNFT)
			proto.Unmarshal(m, dbnft)
			if dbnft.LockTime == 0 {
				dbnft, err = bc.ConverOldNFTToNew(m, dbnft)
				if err != nil {
					log.Infof("Error while staking nft : %s \n", err)
				}
			}
			Staked = append(Staked, dbnft)

		}

		i = i + 1
	}
	log.Info("Checking for Staked NFT's")

	return Staked
}

func (bc *Blockchain) ConverOldNFTToNew(m []byte, nft *pb.StakeNFT) (*pb.StakeNFT, error) {
	oldstake := new(pb.StakeNFTOldVersion)
	proto.Unmarshal(m, oldstake)
	loctime, err := strconv.ParseInt(oldstake.LockTime, 10, 64)
	if err != nil {
		return nil, err
	} else {
		nft.LockTime = uint64(loctime)

	}
	unloctime, err := strconv.ParseInt(oldstake.UnlockTime, 10, 64)
	if err != nil {
		return nil, err
	} else {
		nft.UnlockTime = uint64(unloctime)

	}
	return nft, nil
}

func CheckNFT(nv []int, nft string) bool {
	nftid, _ := strconv.ParseInt(nft, 10, 64)
	for _, id := range nv {
		if id == int(nftid) {
			return true
		}
	}
	return false
}
func (bc *Blockchain) CheckNFTRewards() []*Transaction {
	ntxpool := []*Transaction{}
	nfts := bc.GetNFT()
	for _, s := range bc.NFTStake {

		locktime := (s.UnlockTime - s.LockTime) / (60 * 1000)
		tn := uint64(time.Now().UnixMilli())
		if tn > s.UnlockTime {

			var reward float32
			var err error
			if CheckNFT(nfts.Superrare, s.NFTid) {
				if locktime == mconfig.Lock3Month {

					reward, err = bc.calculatereward(2, s.StakeAmount)
					if err != nil {
						log.Criticalf("Superrare NFT Reward Err: %s", err)
					}
				} else if locktime == mconfig.Lock6Month {
					reward, err = bc.calculatereward(5, s.StakeAmount)
					if err != nil {
						log.Criticalf("Superrare NFT Reward Err: %s", err)
					}
				} else if locktime == mconfig.Lock9Month {
					reward, err = bc.calculatereward(8, s.StakeAmount)
					if err != nil {
						log.Criticalf("Superrare NFT Reward Err: %s", err)
					}
				} else if locktime == mconfig.Lock12Month {
					reward, err = bc.calculatereward(11, s.StakeAmount)
					if err != nil {
						log.Criticalf("Superrare NFT Reward Err: %s", err)
					}
				} else {
					log.Infof("Unexpect error in %v NFT \n", s.NFTid)
					return nil
				}
			} else if CheckNFT(nfts.Rare, s.NFTid) {
				if locktime == mconfig.Lock3Month {

					reward, err = bc.calculatereward(1.7, s.StakeAmount)
					if err != nil {
						log.Criticalf("Rare NFT Reward Err: %s", err)
					}
				} else if locktime == mconfig.Lock6Month {
					reward, err = bc.calculatereward(4.25, s.StakeAmount)
					if err != nil {
						log.Criticalf("Rare NFT Reward Err: %s", err)
					}
				} else if locktime == mconfig.Lock9Month {
					reward, err = bc.calculatereward(6.8, s.StakeAmount)
					if err != nil {
						log.Criticalf("Rare NFT Reward Err: %s", err)
					}
				} else if locktime == mconfig.Lock12Month {
					reward, err = bc.calculatereward(9.35, s.StakeAmount)
					if err != nil {
						log.Criticalf("Rare NFT Reward Err: %s", err)
					}
				} else {
					log.Infof("Unexpect error in %v NFT \n", s.NFTid)
					return nil
				}
			} else if CheckNFT(nfts.LessCommon, s.NFTid) {
				if locktime == mconfig.Lock3Month {

					reward, err = bc.calculatereward(1.3, s.StakeAmount)
					if err != nil {
						log.Criticalf("LessCommon NFT Reward Err: %s", err)
					}
				} else if locktime == mconfig.Lock6Month {
					reward, err = bc.calculatereward(3.25, s.StakeAmount)
					if err != nil {
						log.Criticalf("LessCommon NFT Reward Err: %s", err)
					}
				} else if locktime == mconfig.Lock9Month {
					reward, err = bc.calculatereward(5.2, s.StakeAmount)
					if err != nil {
						log.Criticalf("LessCommon NFT Reward Err: %s", err)
					}
				} else if locktime == mconfig.Lock12Month {
					reward, err = bc.calculatereward(7.15, s.StakeAmount)
					if err != nil {
						log.Criticalf("LessCommon NFT Reward Err: %s", err)
					}
				} else {
					log.Infof("Unexpect error in %v NFT \n", s.NFTid)
					return nil
				}
			} else if CheckNFT(nfts.Common, s.NFTid) {
				if locktime == mconfig.Lock3Month {

					reward, err = bc.calculatereward(1, s.StakeAmount)
					if err != nil {
						log.Criticalf("Common NFT Reward Err: %s", err)
					}
				} else if locktime == mconfig.Lock6Month {
					reward, err = bc.calculatereward(2.5, s.StakeAmount)
					if err != nil {
						log.Criticalf("Common NFT Reward Err: %s", err)
					}
				} else if locktime == mconfig.Lock9Month {
					reward, err = bc.calculatereward(4, s.StakeAmount)
					if err != nil {
						log.Criticalf("Common NFT Reward Err: %s", err)
					}
				} else if locktime == mconfig.Lock12Month {
					reward, err = bc.calculatereward(5.5, s.StakeAmount)
					if err != nil {
						log.Criticalf("Common NFT Reward Err: %s", err)
					}
				} else {
					log.Infof("Unexpect error in %v NFT \n", s.NFTid)
					return nil
				}
			}

			tx := NewTransaction(MINING_SENDER+"( MET NFT #"+s.NFTid+")", s.NFTSender, reward)
			ntxpool = append(ntxpool, tx)
			am, err := strconv.ParseFloat(s.StakeAmount, 64)
			if err != nil {
				log.Criticalf("Error processing staking amount: %s", err)
			}
			oldtx := NewTransaction(STAKING_SENDER+"( MET NFT #"+s.NFTid+")", s.NFTSender, float32(am))
			ntxpool = append(ntxpool, oldtx)

			kp := fmt.Sprintf("%v", s.NFTid)
			bc.Db.Delete(stakednftbucket.Key([]byte(kp)))

		}

	}
	bc.NFTStake = bc.GetStakedNFT()
	return ntxpool

}

func (bc *Blockchain) calculatereward(lt float32, met string) (float32, error) {
	fm, err := strconv.ParseFloat(met, 64)
	if err != nil {
		log.Criticalf("Error calculating reward: %s", err)
	}
	return ((lt * float32(fm)) / 100), err
}

func (bc *Blockchain) ClearStake(ls []*int64) {
	newpool := []*pb.StakeNFT{}

	for _, nft := range bc.NFTStake {
		err := error(nil)
	NFTCheck:
		for _, v := range ls {
			if nft.NFTid != fmt.Sprintf("%v", *v) {
				newpool = append(newpool, nft)
				err = nil
				continue NFTCheck
			} else {
				err = fmt.Errorf("Error locating NFT")
			}
		}
		if err != nil {
			log.Criticalf("Error in clearing staking pool : %s", err)
		}
	}
	if len(ls) >= 1 {
		bc.NFTStake = newpool
	}
	for _, v := range ls {

		kp := fmt.Sprintf("%v", *v)
		err := bc.Db.Delete(stakednftbucket.Key([]byte(kp)))
		if err != nil {
			log.Criticalf("Error in clearing staking pool : %s", err)
		}
	}
}
