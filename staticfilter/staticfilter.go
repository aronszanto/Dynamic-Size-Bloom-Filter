/*

 Before we chose to partition our filter, we built out this plain static filter, which
 forms the foundation for scaleablefilter.go.

 Though the variable names m, k, n, and h are not particularly descriptive, they are consistent
 with the research available on bloom filters, so we have chosen to keep them.

 Joseph Kahn    josephkahn@college.harvard.edu
 Grace Lin      glin@college.harvard.edu
 Aron Szanto    aszanto@college.harvard.edu

*/

package StaticFilter

// Including library packages referenced in this file
import (
	"encoding/binary"
	"github.com/aszanto9/Blumo/bloom_i"
	"github.com/willf/bitset"
	"hash"
	"hash/fnv"
	"math"
)

// Type definition for standard bloom filter
type FilterBase struct {
	m         uint        // size of bitset
	k         uint        // number of hash functions
	h         hash.Hash64 // hashing generator
	err_bound float64     // false-positive error bound
	n         uint
}

type Filter struct {
	params  *FilterBase    // keeps track of underlying filter parameters struct
	bset    *bitset.BitSet // pointer to bitset
	counter uint           // counts elements in filter
}

/*
 Given the set size and an acceptable error bound for false positives
 calculates the size of the bitset, and the number of hash functions
 required for insertion
*/
func NewFilterBase(num uint, err_bound float64) *FilterBase {
	filter_base := new(FilterBase)
	filter_base.err_bound = err_bound
	// calculating required legnth, m, of bitset
	filter_base.m = uint(math.Ceil(-1 * (float64(num) * math.Log(err_bound)) / (math.Log(2) * math.Log(2))))
	// calculate num hash functions
	filter_base.k = uint(math.Ceil(math.Log(err_bound) / math.Log(2)))
	filter_base.h = fnv.New64a()
	filter_base.n = num
	return filter_base
}

/*
 For each string in the set, CalcBits returns a slice of the bits to be flipped.

 i.e. Given a string, S, and a number of hash functions, k, a slice of length k
 is returned, where each element in the slice corresponds to the hash of a given hash
 function of S.
*/
func (filter_base *FilterBase) CalcBits(data []byte) []uint32 {
	filter_base.h.Reset()
	filter_base.h.Write(data)
	hash := filter_base.h.Sum(nil)
	h1 := binary.BigEndian.Uint32(hash[0:4])
	h2 := binary.BigEndian.Uint32(hash[4:8])
	indices := make([]uint32, filter_base.k)
	for i := range indices[:filter_base.k] {
		indices[i] = (uint32(h2) + uint32(h1)*uint32(i)) % uint32(filter_base.m)
	}
	return indices
}

func NewFilter(num uint, err_bound float64) bloom_i.BLOOMFILTER {
	filter := new(Filter)
	filter.params = NewFilterBase(num, err_bound)
	//fmt.Printf(fmt.Sprint("m = ", filter.params.m, "\n"))
	filter.bset = bitset.New(filter.params.m)
	return filter
}

/*
 The following three functions are methods for accessing
 filter parameters externally.
*/
func (filter *Filter) M() uint {
	return filter.params.m
}

func (filter *Filter) GetN() uint {
	return filter.params.n
}

func (filter *Filter) K() uint {
	return filter.params.k
}

func (filter *Filter) E() float64 {
	return filter.params.err_bound
}

func (filter *Filter) Count() uint {
	return filter.counter
}

func (filter *Filter) ApproxP() float64 {
	return float64(filter.counter) / float64(filter.params.m)
}

/*
 Takes in a slice of data, uses CalcBits to determine which bits to flip,
 and then sets the appopriate bits to 1s.
*/
func (filter *Filter) Insert(data []byte) {
	indices := filter.params.CalcBits(data)
	for i := uint(0); i < uint(len(indices)); i++ {
		filter.bset = filter.bset.Set(uint(indices[i]))
	}
	filter.counter++
}

/*
 Follows a similar method to Insert (in terms of calculating indexes of bits),
 but instead of flipping bits, checks to see if they have already been flipped.
*/
func (filter *Filter) Lookup(data []byte) bool {
	indices := filter.params.CalcBits(data)
	for i := 0; i < len(indices); i++ {
		if !filter.bset.Test(uint(indices[i])) {
			// definitely not in set
			return false
		}
	}
	// might be in set, with false-positive rate given by err_bound
	return true
}
