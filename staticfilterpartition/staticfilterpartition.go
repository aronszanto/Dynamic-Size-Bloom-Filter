/*

 This file implements the foundation for our CS51 final project.

 Joseph Kahn    josephkahn@college.harvard.edu
 Grace Lin      glin@college.harvard.edu
 Aron Szanto    aszanto@college.harvard.edu

*/

package StaticFilterPartition

// Including library packages referenced in this file
import (
	"encoding/binary"
	//"fmt"
	"github.com/willf/bitset"
	"hash"
	"hash/fnv"
	"math"
)

// type definition for standard bloom filter
type FilterBase struct {
	m  uint        // size of bitset
	k  uint        // number of hash functions
	h  hash.Hash64 // hashing generator
	e  float64     // error bound
	p  float64     // fill ratio
	sp uint        // size partition
	n  uint        // expected number of keys

}

type Filter struct {
	params   *FilterBase      // needed for generation
	ps       []*bitset.BitSet // partitions
	Counter  uint             // counts elements
	hashedis []uint
}

/*
 calculates the length of the bitset and the number of
 required hash functions given the size of set being
 stored and the acceptable error bound for the task at hand
*/
func NewFilterBase(num uint, eps float64) *FilterBase {
	p := 0.5
	k := calcK(eps)
	m := calcM(num, eps, p)
	sp := calcSP(m, k)

	return &FilterBase{
		p:  p,
		k:  k,
		m:  m,
		sp: sp,
		e:  eps,
		n:  num,
		h:  fnv.New64a(),
	}
}

func calcK(eps float64) uint {
	return uint(math.Ceil(math.Log2(1 / eps)))
}

func calcM(n uint, eps, p float64) uint {
	return uint(math.Ceil(float64(n) /
		math.Abs(((math.Log(p) * math.Log(1-p)) / math.Log(eps)))))
}

func calcSP(m, k uint) uint {
	return uint(math.Ceil(float64(m) / float64(k)))
}

func NewFilter(num uint, eps float64) *Filter {
	init_params := NewFilterBase(num, eps)
	return &Filter{
		params:   init_params,
		ps:       setParts(init_params.sp, init_params.k),
		Counter:  uint(0),
		hashedis: make([]uint, init_params.k),
	}

	/*filter.params = NewFilterBase(num, eps)
	fmt.Printf(fmt.Sprint("m = ", filter.params.m, "\n"))
	filter.ps = setParts(filter.params.sp, filter.params.k)
	return filter*/
}

func (f *Filter) CalcBits(d []byte) {
	//fmt.Printf("CalcBits\n")
	f.params.h.Reset()
	f.params.h.Write(d)
	hash := f.params.h.Sum(nil)
	h1 := binary.BigEndian.Uint32(hash[0:4])
	h2 := binary.BigEndian.Uint32(hash[4:8])
	//o := fmt.Sprint("h1 = ", h1, " and h2 = ", h2, "\n")
	//fmt.Printf(o)
	for i := range f.hashedis[:f.params.k] {
		f.hashedis[i] = (uint(h2) + uint(h1)*uint(i)) % f.params.sp //changed this line to i-1 and the above to <= and 1
	}

	//op := fmt.Sprint(indices, " : bits set to be flipped\n")
	//fmt.Printf(op)
}

func (f *Filter) BitsFlipped() uint {
	c := uint(0)
	for i := range f.ps {
		c += f.ps[i].Count()
	}
	return c
}

func setParts(sp, k uint) []*bitset.BitSet {
	sets := make([]*bitset.BitSet, k)
	for i := range sets {
		sets[i] = bitset.New(sp)
	}
	return sets
}

func (f *Filter) M() uint {
	return f.params.m
}

func (f *Filter) K() uint {
	return f.params.k
}

func (f *Filter) N() uint {
	return f.params.n
}

func (f *Filter) E() float64 {
	return f.params.e
}

// return approximate fill ratio across the k partitions
func (f *Filter) ApproxP() float64 {
	return 1.0 - math.Exp(-float64(f.Counter)/float64(f.params.m))
}

// Takes in a slice of indexes
func (filter *Filter) Insert(data []byte) {
	//p := fmt.Sprint("\nOperating onfilter with k = ", filter.params.k, " and m = ", filter.params.m, "\n\n\n")
	//fmt.Printf(p)
	filter.CalcBits(data)
	//fmt.Printf("Insert after CalcBits\n")
	for i, v := range filter.hashedis[:filter.params.k] {
		//fmt.Printf(fmt.Sprint("Loop iteration ", i, "\n"))
		filter.ps[i].Set(v)
	}
	filter.Counter++
}

func (filter *Filter) Lookup(data []byte) bool {
	filter.CalcBits(data)
	for i, v := range filter.hashedis[:filter.params.k] {
		if !filter.ps[i].Test(uint(v)) {
			// definitely not in set
			return false
		}
	}
	// may be in set
	return true
}

func (filter *Filter) Reset() {
	filter.params = NewFilterBase(filter.params.n, filter.params.e)
	filter.ps = setParts(filter.params.sp, filter.params.k)
	filter.Counter = uint(0)
}
