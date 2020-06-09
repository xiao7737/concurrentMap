package concurrentMap

import (
	"strconv"
	"sync"
	"testing"
)

func BenchmarkConcurrentMap_Set(b *testing.B) {
	cm := CreateConcurrentMap(32)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go cm.Set(ConvertStr(strconv.Itoa(i)), "test")
	}
}

func BenchmarkSyncMap_Store(b *testing.B) {
	var sm sync.Map
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go sm.Store(strconv.Itoa(i), "test")
	}
}

func BenchmarkConcurrentMap_Get(b *testing.B) {
	cm := CreateConcurrentMap(32)
	for i := 0; i < b.N; i++ {
		go cm.Set(ConvertStr(strconv.Itoa(i)), "test")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go cm.Get(ConvertStr(strconv.Itoa(i)))
	}
}

func BenchmarkSyncMap_Load(b *testing.B) {
	var sm sync.Map
	for i := 0; i < b.N; i++ {
		go sm.Store(strconv.Itoa(i), "test")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go sm.Load(strconv.Itoa(i))
	}
}

func BenchmarkConcurrentMap_Del(b *testing.B) {
	cm := CreateConcurrentMap(32)
	for i := 0; i < b.N; i++ {
		go cm.Set(ConvertStr(strconv.Itoa(i)), "test")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go cm.Del(ConvertStr(strconv.Itoa(i)))
	}
}

func BenchmarkSyncMap_Delete(b *testing.B) {
	var sm sync.Map
	for i := 0; i < b.N; i++ {
		go sm.Store(strconv.Itoa(i), i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go sm.Delete(strconv.Itoa(i))
	}
}
