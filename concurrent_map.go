package concurrentMap

import (
	"sync"
)

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

// add Lock for set kv
func (s *shardMap) set(key PartitionKey, value interface{}) {
	keyVal := key.Value()
	s.lock.Lock()
	s.m[keyVal] = value
	s.lock.Unlock()
}

// add RLock for get key
func (s *shardMap) get(key PartitionKey) (value interface{}, exists bool) {
	keyVal := key.Value()
	s.lock.RLock()
	value, exists = s.m[keyVal]
	s.lock.RUnlock()
	return

}

// add Lock for del key
func (s *shardMap) del(key PartitionKey) {
	keyVal := key.Value()
	s.lock.Lock()
	delete(s.m, keyVal)
	s.lock.Unlock()
}

// count the elements of a shard
func (s *shardMap) count() (count int) {
	s.lock.RLock()
	count = len(s.m)
	s.lock.RUnlock()
	return
}

// get the shard for key
// routing algorithm：shard = hash(routing_key) % number_of_shards
// similar to redis-cluster and es
func (m *ConcurrentMap) getShard(key PartitionKey) *shardMap {
	shardIndex := key.PartitionKey() % (int64(m.numOfShard))
	return m.shard[shardIndex]
}

// CreateConcurrentMap is to create a Map with number entered by the user
//todo 切换成New的方式
func CreateConcurrentMap(numOfShard int) *ConcurrentMap {
	var shard []*shardMap
	for i := 0; i < numOfShard; i++ {
		shard = append(shard, createShardMap())
	}
	return &ConcurrentMap{shard, numOfShard}
}

// Set is to set kv
// Set the same key, the new value will overwrite the old value
func (m *ConcurrentMap) Set(key PartitionKey, value interface{}) {
	shardMap := m.getShard(key)
	shardMap.set(key, value)
}

// Get is to get kv and return whether the key exists
func (m *ConcurrentMap) Get(key PartitionKey) (value interface{}, exists bool) {
	shardMap := m.getShard(key)
	return shardMap.get(key)
}

// Del is to delete the key
// There will be no prompt to delete a key that is not in a map
func (m *ConcurrentMap) Del(key PartitionKey) {
	shardMap := m.getShard(key)
	shardMap.del(key)
}

// Count is to count the total elements of the concurrentMap
func (m *ConcurrentMap) Count() (count int) {
	for i := 0; i < m.numOfShard; i++ {
		shardMap := m.shard[i]
		count += shardMap.count()
	}
	return
}

// Exists is to return whether key exists in the map
func (m *ConcurrentMap) Exists(key PartitionKey) (exists bool) {
	shardMap := m.getShard(key)
	_, exists = shardMap.get(key)
	return
}
