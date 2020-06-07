package concurrentMap

import (
	"strconv"
	"sync"
	"testing"
)

func BenchmarkConcurrentMap_Set(b *testing.B) {
	var sm sync.Map
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Store(strconv.Itoa(i), i)
	}
}

func BenchmarkSyncMap_Store(b *testing.B) {
	cm := CreateConcurrentMap(32)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cm.Set(ConvertStr(strconv.Itoa(i)), i)
	}
}

func BenchmarkConcurrentMap_Get(b *testing.B) {
	var sm sync.Map
	for i := 0; i < b.N; i++ {
		sm.Store(strconv.Itoa(i), i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Load(strconv.Itoa(i))
	}
}

func BenchmarkSyncMap_Load(b *testing.B) {
	cm := CreateConcurrentMap(32)
	for i := 0; i < b.N; i++ {
		cm.Set(ConvertStr(strconv.Itoa(i)), i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cm.Get(ConvertStr(strconv.Itoa(i)))
	}
}
