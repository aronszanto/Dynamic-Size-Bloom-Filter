/*
  INTERFACE FOR BLOOM FILTERS. ALL FILTERS MUST IMPLEMENT THIS INTERFACE.

*/
package bloom_i

import "math"

type BLOOMFILTER interface {
	Insert(data []byte)      // inserts a key
	Lookup(data []byte) bool // checks whether a key has been inserted, never returns false negatives
	Count() uint             // number of keys inserted
	GetN() uint              // expected number of keys
	ApproxP() float64        // approximate fill ratio
}

// methods calculate number of hash functions, size of bitset, and size of partition.
func CalcK(err_bound float64) uint {
	return uint(math.Ceil(math.Log2(float64(1) / float64(err_bound))))
}

func CalcM(n uint, err_bound, fill_ratio float64) uint {
	return uint(math.Ceil(float64(n) /
		math.Abs(((math.Log(fill_ratio) * math.Log(1-fill_ratio)) / math.Log(err_bound)))))
}

func CalcPartitionSize(m, k uint) uint {
	return uint(math.Ceil(float64(m) / float64(k)))
}
