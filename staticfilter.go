/*

 This file implements the foundation for our CS51 final project.

 Joseph Kahn    josephkahn@college.harvard.edu
 Grace Lin      glin@college.harvard.edu
 Aron Szanto    aszanto@college.harvard.edu

*/

package StaticFilter

// Including library packages referenced in this file
import (
	// "encoding/binary"
	"github.com/willf/bitset"
	"hash"
	// "hash/fnv"
	"math"
)

// type definition for standard bloom filter
type FilterBase struct {
	m uint        // size of bitset
	k uint        // number of hash functions
	h hash.Hash32 // hashing generator
}

type Filter struct {
	params *FilterBase    // needed for generation
	b      *bitset.BitSet // pointer to bitset
}

/*
 calculates the length of the bitset and the number of
 required hash functions given the size of set being
 stored and the acceptable error bound for the task at hand
*/
func NewFilterBase(num uint, eps float64) *FilterBase {
	fb := new(FilterBase)
	// calculating length
	fb.m = uint(math.Ceil(-1 * (float64(num) * math.Log(eps)) / (math.Log(2) * math.Log(2))))
	// calculate num hash functions
	fb.k = uint(math.Ceil((float64(fb.m) / float64(num)) * math.Log(2)))
	return fb
}

func (fb *FilterBase) CalcBits(d []byte) []uint32 {
	fb.h = fnv.New32a()
	fb.h.Reset()
	fb.h.Write(d)
	hash_stream := fb.h.Sum(nil)
	h1 := binary.BigEndian.Uint32(hash_stream[4:8])
	h2 := binary.BigEndian.Uint32(hash_stream[0:4])

	indices := make([]uint32, fb.k)
	for i := uint32(0); i < uint32(fb.k); i++ {
		indices[i] = (h1 + h2*i) % uint32(fb.m)
	}
	return indices
}

func NewFilter(num uint, eps float64) *Filter {
	filter := new(Filter)
	filter.params = NewFilterBase(num, eps)
	filter.b = bitset.New(filter.params.m)
	return filter
}

// Takes in a slice of indexes
func (filter *Filter) Insert(data []byte) {
	indices := filter.params.CalcBits(data)
	for i := uint(0); i < uint(len(indices)); i++ {
		filter.b = filter.b.Set(uint(indices[i]))
	}
}

func (filter *Filter) Lookup(data []byte) bool {
	indices := filter.params.CalcBits(data)
	// might be there unless definitely not in set
	var found bool
	for i := uint(0); i < uint(len(indices)); i++ {
		if filter.b.Test(uint(i)) == false {
			// definitely not in set
			found = false
			break
		} else {
			// might be in the set
			found = true
		}
	}
	return found
}

func (filter *Filter) Reset() {
	filter.b = filter.b.ClearAll()
}
