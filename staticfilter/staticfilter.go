/*

 This file implements the foundation for our CS51 final project.

 Joseph Kahn    josephkahn@college.harvard.edu
 Grace Lin      glin@college.harvard.edu
 Aron Szanto    aszanto@college.harvard.edu

*/

package StaticFilter

// Including library packages referenced in this file
import (
	"fmt"
	"github.com/willf/bitset"
	"hash"
	"hash/fnv"
	"math"
)

// type definition for standard bloom filter
type FilterBase struct {
	m         uint        // size of bitset
	k         uint        // number of hash functions
	h         hash.Hash64 // hashing generator
	err_bound float64     // false-positive error bound
}

type Filter struct {
	params  *FilterBase    // needed for generation
	bset    *bitset.BitSet // pointer to bitset slice
	counter uint           // counts elements
}

/*
 calculates the length of the bitset and the number of
 required hash functions given the size of set being
 stored and the acceptable error bound for the task at hand
*/
// clean this up, make it one statement
func NewFilterBase(num uint, err_bound float64) *FilterBase {
	filter_base := new(FilterBase)
	filter_base.err_bound = err_bound
	// calculating length
	filter_base.m = uint(math.Ceil(-1 * (float64(num) * math.Log(err_bound)) / (math.Log(2) * math.Log(2))))
	// calculate num hash functions
	filter_base.k = uint(math.Ceil(math.Log(err_bound) / math.Log(2)))
	// Pretty sure you can just do this
	// filter_base:= FilterBase{*insert equation for m here, inesrt equation for k here, h, err_bound}
	return filter_base
}

func (filter_base *FilterBase) CalcBits(d []byte) []uint32 {
	filter_base.h = fnv.New64a()
	filter_base.h.Reset()
	filter_base.h.Write(d)
	hash := filter_base.h.Sum64()
	h1 := uint32(hash & ((1 << 32) - 1))
	h2 := uint32(hash >> 32)
	//o := fmt.Sprint("h1 = ", h1, " and h2 = ", h2, "\n")
	//fmt.Printf(o)

	indices := make([]uint32, filter_base.k)
	for i := uint32(0); i <= uint32(filter_base.k); i++ {
		indices[i] = (h1 + h2*i) % uint32(filter_base.m) //changed this line to i-1 and the above to <= and 1
	}
	//op := fmt.Sprint(indices, " : bits set to be flipped\n")
	//fmt.Printf(op)
	return indices
}

func NewFilter(num uint, err_bound float64) *Filter {
	filter := new(Filter)
	filter.params = NewFilterBase(num, err_bound)
	fmt.Printf(fmt.Sprint("m = ", filter.params.m, "\n"))
	filter.bset = bitset.New(filter.params.m)
	return filter
}

func (filter *Filter) M() uint {
	return filter.params.m
}

func (filter *Filter) K() uint {
	return filter.params.k
}

func (filter *Filter) E() float64 {
	return filter.params.err_bound
}

// Takes in a slice of indexes
func (filter *Filter) Insert(data []byte) {
	//p := fmt.Sprint("\nOperating onfilter with k = ", filter.params.k, " and m = ", filter.params.m, "\n\n\n")
	//fmt.Printf(p)
	indices := filter.params.CalcBits(data)
	for i := uint(0); i < uint(len(indices)); i++ {
		filter.bset = filter.bset.Set(uint(indices[i]))
	}
	filter.counter++
}

func (filter *Filter) Lookup(data []byte) bool {
	indices := filter.params.CalcBits(data)
	// might be there unless definitely not in set
	for i := 0; i < len(indices); i++ {
		if !filter.bset.Test(uint(indices[i])) {
			// definitely not in set
			//op := fmt.Sprint("Bit #", i, " = ", indices[i], " would have been \n")
			//fmt.Printf(op)
			return false
		}
	}
	return true
}
