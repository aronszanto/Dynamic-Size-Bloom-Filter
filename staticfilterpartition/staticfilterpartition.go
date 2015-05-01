/*

 This file implements the foundation for our CS51 final project. Scaleable filters are essentially
 a whole bunch of static filters takced onto each other.

 Joseph Kahn    josephkahn@college.harvard.edu
 Grace Lin      glin@college.harvard.edu
 Aron Szanto    aszanto@college.harvard.edu

*/

package StaticFilterPartition

// Including library packages referenced in this file
import (
	"encoding/binary"
	"github.com/willf/bitset"
	"hash"
	"hash/fnv"
	"math"
)

// type definition for standard bloom filter
type FilterBase struct {
	m              uint        // size of bitset
	k              uint        // number of hash functions
	h              hash.Hash64 // hashing generator
	err_bound      float64     // error bound
	fill_ratio     float64     // fill ratio
	partition_size uint        // size partition
	n              uint        // expected number of keys

}

type Filter struct {
	params     *FilterBase      // needed for generation
	partitions []*bitset.BitSet // partitions
	counter    uint             // counts elements
	hashed_is  []uint           // keeps track of the int representations which need to be flipped in bitset
}

/*
 calculates the length of the bitset and the number of
 required hash functions given the size of set being
 stored and the acceptable error bound for the task at hand
*/
func NewFilterBase(num uint, err_bound float64) *FilterBase {
	fill_ratio := 0.5
	k := calcK(err_bound)
	m := calcM(num, err_bound, fill_ratio)
	partition_size := calcPartitionSize(m, k)

	return &FilterBase{
		fill_ratio:     fill_ratio,
		k:              k,
		m:              m,
		partition_size: partition_size,
		err_bound:      err_bound,
		n:              num,
		h:              fnv.New64a(),
	}
}

func calcK(err_bound float64) uint {
	return uint(math.Ceil(math.Log2(1 / err_bound)))
}

func calcM(n uint, err_bound, fill_ratio float64) uint {
	return uint(math.Ceil(float64(n) /
		math.Abs(((math.Log(fill_ratio) * math.Log(1-fill_ratio)) / math.Log(err_bound)))))
}

func calcPartitionSize(m, k uint) uint {
	return uint(math.Ceil(float64(m) / float64(k)))
}

func NewFilter(num uint, err_bound float64) *Filter {
	init_params := NewFilterBase(num, err_bound)
	return &Filter{
		params:     init_params,
		partitions: setParts(init_params.partition_size, init_params.k),
		counter:    uint(0),
		hashed_is:  make([]uint, init_params.k),
	}
}

/*
 For each string in the set, CalcBits returns a slice of the bits to be flipped.

 i.e. Given a string, S, and a number of hash functions, k, a slice of length k
 is returned, where each element in the slice corresponds to the hash of a given hash
 function of S.
*/
func (filter *Filter) CalcBits(d []byte) {
	//fmt.Printf("CalcBits\n")
	filter.params.h.Reset()
	filter.params.h.Write(d)
	hash := filter.params.h.Sum(nil)
	h1 := binary.BigEndian.Uint32(hash[0:4])
	h2 := binary.BigEndian.Uint32(hash[4:8])
	for i := range filter.hashed_is[:filter.params.k] {
		filter.hashed_is[i] = (uint(h2) + uint(h1)*uint(i)) % filter.params.partition_size
	}
}

func (f *Filter) BitsFlipped() uint {
	c := uint(0)
	for i := range f.partitions {
		c += f.partitions[i].Count()
	}
	return c
}

func setParts(partition_size, k uint) []*bitset.BitSet {
	sets := make([]*bitset.BitSet, k)
	for i := range sets {
		sets[i] = bitset.New(partition_size)
	}
	return sets
}

func (filter *Filter) M() uint {
	return filter.params.m
}

func (filter *Filter) K() uint {
	return filter.params.k
}

func (filter *Filter) N() uint {
	return filter.params.n
}

func (filter *Filter) E() float64 {
	return filter.params.err_bound
}

// return approximate fill ratio across the partitions
func (filter *Filter) ApproxP() float64 {
	return 1.0 - math.Exp(-float64(filter.counter)/float64(filter.params.m))
}

/*
 Takes in data (in the case of a dictionary, strings), runs it through CalcBits,
 which returns a slice of bits to be flipper, and flips appropriate bits.
*/
func (filter *Filter) Insert(data []byte) {
	filter.CalcBits(data)
	for i, v := range filter.hashed_is[:filter.params.k] {
		filter.partitions[i].Set(v)
	}
	filter.counter++
}

/*
 Finds the indexes of a given element (as in Insert), but instead of flipping bits,
 checks to see if they have been set to 1s. If they have, it's possible that an element
 which is not a member returns true. However, if an element HAS been inserted, it will
 never come back as a false negative.
*/
func (filter *Filter) Lookup(data []byte) bool {
	filter.CalcBits(data)
	for i, v := range filter.hashed_is[:filter.params.k] {
		if !filter.partitions[i].Test(uint(v)) {
			// definitely not in set
			return false
		}
	}
	// may be in set, with false positive rate given by err_bound
	return true
}

func (filter *Filter) Reset() {
	filter.params = NewFilterBase(filter.params.n, filter.params.err_bound)
	filter.partitions = setParts(filter.params.partition_size, filter.params.k)
	filter.counter = uint(0)
}
