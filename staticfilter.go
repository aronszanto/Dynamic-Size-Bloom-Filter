/*

 This file implements the foundation for our CS51 final project.

 Joseph Kahn    josephkahn@college.harvard.edu
 Grace Lin      glin@college.harvard.edu
 Aron Szanto    aszanto@college.harvard.edu

*/

package StaticFilter

// Including library packages referenced in this file
import (
	"github.com/willf/bitset"
	"hash"
	"hash/fnv"
	"math"
)

// type definition for standard bloom filter
type FilterBase struct {
	m uint64      // size of bitset
	k uint64      // number of hash functions
	h hash.Hash64 // hash function generator
}

type Filter struct {
	params *FilterBase // needed for generation
	b      *BitSet     // pointer to bitset
}

/*
 calculates the length of the bitset and the number of
 required hash functions given the size of set being
 stored and the acceptable error bound for the task at hand
*/
func NewFilterBase(num uint64, eps float64) *FilterBase {
	fb := new(FilterBase)
	// calculating length
	fb.m = uint64(math.Ceil(-1 * (float64(num) * math.Log(eps)) / (math.Log(2) * math.Log(2))))
	// calculate num hash functions
	fb.k = uint64(math.Ceil((float64(m) / float64(num)) * math.Log(2)))
	return &fb
}

func NewFilter(num uint64, eps float64) *Filter {
	filter := new(Filter)
	filter.params = NewFilterBase(num, eps)
	filter.b = BitSet.New(filter.params.m)
	return &filter
}

/*
// Takes in a slice of indexes
func Insert() {

}
*/
