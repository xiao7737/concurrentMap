package concurrentMap

import "testing"

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
	cm.Set(ConvertStr("first, "), "live")
	cm.Set(ConvertStr("then, "), "die")
	if cm.Count() != 2 {
		t.Error("map should contain two elements")
	}
}
