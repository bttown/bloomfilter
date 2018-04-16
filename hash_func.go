package bloomfilter

import (
	"hash/fnv"
)

func hash(v interface{}) uint64 {
	switch v.(type) {
	case string:
		h := fnv.New64a()
		h.Write([]byte(v.(string)))
		return h.Sum64()
	default:
		panic("the type not supported!")
	}
}
