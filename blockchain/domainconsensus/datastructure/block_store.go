package datastructure

import (
	"github.com/Metchain/MetblockD/app/model"
	"github.com/Metchain/MetblockD/blockchain/consensus/database/serialization"
	"github.com/Metchain/MetblockD/external"
	"github.com/Metchain/MetblockD/utils/lrucache"
	"github.com/Metchain/MetblockD/utils/staging"
	"github.com/golang/protobuf/proto"
)

var bucketName = []byte("blocks")

// blockStore represents a store of blocks
type blockStore struct {
	shardID     model.StagingShardID
	cache       *lrucache.LRUCache
	countCached uint64
	bucket      model.DBBucket
	countKey    model.DBKey
}

func (bs *blockStore) Stage(stagingArea *model.StagingArea, blockHash *external.DomainHash, block *external.DomainBlock) {
	//TODO implement me
	panic("implement me")
}

func (bs *blockStore) IsStaged(stagingArea *model.StagingArea) bool {
	//TODO implement me
	panic("implement me")
}

func (bs *blockStore) Block(dbContext model.DBReader, stagingArea *model.StagingArea, blockHash *external.DomainHash) (*external.DomainBlock, error) {
	//TODO implement me
	panic("implement me")
}

func (bs *blockStore) HasBlock(dbContext model.DBReader, stagingArea *model.StagingArea, blockHash *external.DomainHash) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (bs *blockStore) Blocks(dbContext model.DBReader, stagingArea *model.StagingArea, blockHashes []*external.DomainHash) ([]*external.DomainBlock, error) {
	//TODO implement me
	panic("implement me")
}

func (bs *blockStore) Delete(stagingArea *model.StagingArea, blockHash *external.DomainHash) {
	//TODO implement me
	panic("implement me")
}

func (bs *blockStore) Count(stagingArea *model.StagingArea) uint64 {
	//TODO implement me
	panic("implement me")
}

func (bs *blockStore) AllBlockHashesIterator(dbContext model.DBReader) (model.BlockIterator, error) {
	//TODO implement me
	panic("implement me")
}

func New(dbContext model.DBReader, prefixBucket model.DBBucket, cacheSize int, preallocate bool) (model.BlockStore, error) {
	blockStore := &blockStore{
		shardID:  staging.GenerateShardingID(),
		cache:    lrucache.New(cacheSize, preallocate),
		bucket:   prefixBucket.Bucket(bucketName),
		countKey: prefixBucket.Key([]byte("blocks-count")),
	}

	err := blockStore.initializeCount(dbContext)
	if err != nil {
		return nil, err
	}

	return blockStore, nil
}

func (bs *blockStore) initializeCount(dbContext model.DBReader) error {
	count := uint64(0)
	hasCountBytes, err := dbContext.Has(bs.countKey)
	if err != nil {
		return err
	}
	if hasCountBytes {
		countBytes, err := dbContext.Get(bs.countKey)
		if err != nil {
			return err
		}
		count, err = bs.deserializeBlockCount(countBytes)
		if err != nil {
			return err
		}
	}
	bs.countCached = count
	return nil
}

func (bs *blockStore) deserializeBlock(blockBytes []byte) (*external.DomainBlock, error) {
	dbBlock := &serialization.DbBlock{}
	err := proto.Unmarshal(blockBytes, dbBlock)
	if err != nil {
		return nil, err
	}
	return serialization.DbBlockToDomainBlock(dbBlock)
}

func (bs *blockStore) deserializeBlockCount(countBytes []byte) (uint64, error) {
	dbBlockCount := &serialization.DbBlockCount{}
	err := proto.Unmarshal(countBytes, dbBlockCount)
	if err != nil {
		return 0, err
	}
	return dbBlockCount.Count, nil
}
