package concurrentMap

type Int64Key struct {
	value int64
}

func (i *Int64Key) PartitionKey() int64 {
	return i.value
}

func (i *Int64Key) Value() interface{} {
	return i.value
}

func ConvertInt64(key int64) *Int64Key {
	return &Int64Key{key}
}
