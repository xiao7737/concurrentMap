package concurrentMap

import "sync"

type ConcurrentMap struct {
	shard      []*shardMap
	numOfShard int
}

// shard map
type shardMap struct {
	m    map[interface{}]interface{}
	lock sync.RWMutex
}

// interface of key
type PartitionKey interface {
	Value() interface{}
	PartitionKey() int64
}

func createShardMap() *shardMap {
	return &shardMap{
		m: make(map[interface{}]interface{}),
	}
}

// CreateConcurrentMap is to create a Map with number entered by the user
func CreateConcurrentMap(numOfShard int) *ConcurrentMap {
	var shard []*shardMap
	for i := 0; i < numOfShard; i++ {
		shard = append(shard, createShardMap())
	}
	return &ConcurrentMap{shard, numOfShard}
}
