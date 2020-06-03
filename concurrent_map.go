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

// add Lock for set kv
func (s *shardMap) set(key PartitionKey, value interface{}) {
	keyVal := key.Value()
	s.lock.Lock()
	s.m[keyVal] = value
	s.lock.Unlock()
}

// add RLock for get key
func (s *shardMap) get(key PartitionKey) (interface{}, bool) {
	keyVal := key.Value()
	s.lock.RLock()
	value, exists := s.m[keyVal]
	s.lock.RUnlock()
	return value, exists

}

// add Lock for del key
func (s *shardMap) del(key PartitionKey) {
	keyVal := key.Value()
	s.lock.Lock()
	delete(s.m, keyVal)
	s.lock.Unlock()
}

// get the shard for key
func (m *ConcurrentMap) getShard(key PartitionKey) *shardMap {
	shardIndex := key.PartitionKey() % (int64(m.numOfShard))
	return m.shard[shardIndex]
}

// CreateConcurrentMap is to create a Map with number entered by the user
func CreateConcurrentMap(numOfShard int) *ConcurrentMap {
	var shard []*shardMap
	for i := 0; i < numOfShard; i++ {
		shard = append(shard, createShardMap())
	}
	return &ConcurrentMap{shard, numOfShard}
}

// Set is to set kv
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
func (m *ConcurrentMap) Del(key PartitionKey) {
	shardMap := m.getShard(key)
	shardMap.del(key)
}

// 实现count，exist，upsert
