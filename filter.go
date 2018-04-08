package bloomfilter

type Filter struct {
	bitMap []uint8
	bitSize int64
}

const numHashFunctions = 5

func New(bitSize int64) *Filter {
	filter := Filter{
		bitSize: bitSize,
		bitMap: make([]uint8, bitSize),
	}

	return &filter
}

func (filter *Filter) Put(v interface{}) {
	var hash1 = hash(v)
	var hash2 = hash1 << 32

	for i := 1; i < numHashFunctions; i++ {
		nextHash := hash1 + uint64(i) * hash2
		if nextHash < 0 {
			nextHash = -nextHash
		}

		index := nextHash % uint64(filter.bitSize)
		filter.bitMap[index] = 1
	}
}

func (filter *Filter) MightContains(v interface{}) bool {
	var hash1 = hash(v)
	var hash2 = hash1 << 32

	for i := 1; i < numHashFunctions; i++ {
		nextHash := hash1 + uint64(i) * hash2
		if nextHash < 0 {
			nextHash = -nextHash
		}

		index := nextHash % uint64(filter.bitSize)
		if filter.bitMap[index] == 0 {
			return false
		}
	}

	return true
}