package concurrentMap

import (
	"strconv"
	"testing"
)

func TestCreateConcurrentMap(t *testing.T) {
	cm := CreateConcurrentMap(32)
	if cm.numOfShard != 32 {
		t.Error("init concurrentMap fail")
	}
	if cm.Count() != 0 {
		t.Error("map should be empty")
	}
}

func TestConcurrentMap_Set(t *testing.T) {
	cm := CreateConcurrentMap(32)
	cm.Set(ConvertStr("first, "), "hello")
	cm.Set(ConvertStr("then, "), "bye")
	if cm.Count() != 2 {
		t.Error("map should contain two elements")
	}
}

func TestConcurrentMap_Get(t *testing.T) {
	cm := CreateConcurrentMap(32)
	cm.Set(ConvertStr("hello"), "go")
	cm.Set(ConvertInt64(333), 666)
	res1, ok := cm.Get(ConvertStr("hello"))
	if res1 != "go" || !ok {
		t.Error("string key get err")
	}
	res2, ok := cm.Get(ConvertInt64(333))
	if res2 != 666 || !ok {
		t.Error("int64 key get err")
	}
}

func TestConcurrentMap_Del(t *testing.T) {
	cm := CreateConcurrentMap(32)
	cm.Set(ConvertStr("hello"), "go")
	cm.Set(ConvertInt64(333), 666)
	cm.Del(ConvertInt64(333))
	if cm.Count() != 1 {
		t.Error("map should only have one key")
	}
	res, ok := cm.Get(ConvertInt64(666))
	if res != nil || ok {
		t.Error("map should not contain this key")
	}
	// Delete a key not in the map
	cm.Del(ConvertStr("usb"))
}

func TestConcurrentMap_Count(t *testing.T) {
	cm := CreateConcurrentMap(32)
	for i := 0; i < 512; i++ {
		cm.Set(ConvertStr("str"+strconv.Itoa(i)), i)
		cm.Set(ConvertInt64(int64(i)), i)
	}
	if cm.Count() != 1024 {
		t.Error("the number of elements in the map should be 1024")
	}
}

func TestConcurrentMap_Exists(t *testing.T) {
	cm := CreateConcurrentMap(32)
	cm.Set(ConvertStr("hello"), "go")
	cm.Set(ConvertInt64(333), 666)

	if !cm.Exists(ConvertStr("hello")) {
		t.Error("hello should be in the map")
	}

	if cm.Exists(ConvertInt64(444)) {
		t.Error("444 should no is exists in the map")
	}
}
