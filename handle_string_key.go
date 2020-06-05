package concurrentMap

import (
	"unsafe"
)

const (
	c1 uint32 = 0xcc9e2d51
	c2 uint32 = 0x1b873593
)

type StringKey struct {
	value string
}

// convert key of type string to type int64
func (s *StringKey) PartitionKey() int64 {
	return int64(murmurHash(s.value))
}

func (s *StringKey) Value() interface{} {
	return s.value
}

func ConvertStr(key string) *StringKey {
	return &StringKey{key}
}

// link：https://github.com/aappleby/smhasher
// Even if the keys entered are regular,
// the algorithm can still give a good random distribution,
// and the calculation speed of the algorithm is also very fast
func murmurHash(str string) uint32 {

	data := ([]byte)(str)
	var h1 uint32 = 37

	blocks := len(data) / 4
	var p uintptr
	if len(data) > 0 {
		p = uintptr(unsafe.Pointer(&data[0]))
	}

	p1 := p + uintptr(4*blocks)
	for ; p < p1; p += 4 {
		k1 := *(*uint32)(unsafe.Pointer(p))
		k1 *= c1
		k1 = (k1 << 15) | (k1 >> 17)
		k1 *= c2

		h1 ^= k1
		h1 = (h1 << 13) | (h1 >> 19)
		h1 = h1*5 + 0xe6546b64
	}

	tail := data[blocks*4:]

	var k1 uint32
	switch len(tail) & 3 {
	case 3:
		k1 ^= uint32(tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint32(tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint32(tail[0])
		k1 *= c1
		k1 = (k1 << 15) | (k1 >> 17)
		k1 *= c2
		h1 ^= k1
	}

	h1 ^= uint32(len(data))

	h1 ^= h1 >> 16
	h1 *= 0x85ebca6b
	h1 ^= h1 >> 13
	h1 *= 0xc2b2ae35
	h1 ^= h1 >> 16

	return (h1 << 24) | (((h1 >> 8) << 16) & 0xFF0000) | (((h1 >> 16) << 8) & 0xFF00) | (h1 >> 24)
}
