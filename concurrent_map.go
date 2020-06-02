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
func (sm *shardMap) set(key PartitionKey, value interface{}) {
	keyVal := key.Value()
	sm.lock.Lock()
	sm.m[keyVal] = value
	sm.lock.Unlock()
}

// add RLock for get key
func (sm *shardMap) get(key PartitionKey) (interface{}, bool) {
	keyVal := key.Value()
	sm.lock.RLock()
	value, exists := sm.m[keyVal]
	sm.lock.RUnlock()
	return value, exists

}

// add Lock for del key
func (sm *shardMap) del(key PartitionKey) {
	keyVal := key.Value()
	sm.lock.Lock()
	delete(sm.m, keyVal)
	sm.lock.Unlock()
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
	sm := m.getShard(key)
	sm.set(key, value)
}

// Get is to get kv and return whether the key exists
func (m *ConcurrentMap) Get(key PartitionKey) (value interface{}, exists bool) {
	sm := m.getShard(key)
	return sm.get(key)
}

// Del is to delete the key
func (m *ConcurrentMap) Del(key PartitionKey) {
	sm := m.getShard(key)
	sm.del(key)
}
