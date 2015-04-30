/*

 This file implements the foundation for our CS51 final project.

 Joseph Kahn    josephkahn@college.harvard.edu
 Grace Lin      glin@college.harvard.edu
 Aron Szanto    aszanto@college.harvard.edu

*/

package ScalableFilterPartition

import (
	"fmt"
	"github.com/aszanto9/Blumo/staticfilterpartition"
	"math"
)

type SBF struct {
	// head points to the current filter that is not yet full
	head *StaticFilterPartition.Filter
	// an array of pointers to all of the existing filters
	filter_slice []*StaticFilterPartition.Filter
	// s is the scaling factor for the size of new filters, N is the number of existing filters
	// m_init is the m of the first filter in an SBF
	s, N, m_init uint // JESUS CHRIST FIX THIS LATER
	// p is the total final error bound, r is the scaling factor for the error bound of new filters
	e, r, p float64
}

/*type SBF interface {
	NewSBF SBF
	SBFlookup bool
	SBFinsert
	NewBF Filter
}*/

func NewFilter(end_e float64) *SBF {
	//default values for s, r (hardcoded)
	m_init_i := uint(10)
	s_i := uint(2)
	p_i := 0.5
	N_i := uint(1)
	r_i := 0.8
	head_i := StaticFilterPartition.NewFilter(uint(m_init_i), end_e*(1-r_i))
	return &SBF{
		m_init:       m_init_i,
		s:            s_i,
		N:            N_i,
		e:            end_e,
		p:            p_i,
		r:            r_i,
		head:         head_i,
		filter_slice: []*StaticFilterPartition.Filter{head_i},
	}
}

func (sbf *SBF) Lookup(data []byte) bool {
	for i := uint(0); i < sbf.N; i++ {
		fmt.Printf(fmt.Sprint("Looking for ", data, " in filter #", i, "\n"))
		if sbf.filter_slice[i].Lookup(data) {
			fmt.Printf(fmt.Sprint("Found in filter ", i, "\n"))
			return true
		}
	}
	return false

}

func (sbf *SBF) addBF() {
	sbf.N++
	newfilter := StaticFilterPartition.NewFilter(sbf.head.N()*sbf.s,
		sbf.e*math.Pow(sbf.r, float64(sbf.N-1)))
	sbf.head = newfilter
	sbf.filter_slice = append(sbf.filter_slice, newfilter)

	fmt.Printf(fmt.Sprint("Bloom filter #", sbf.N, " added\n"))

}

func (sbf *SBF) Insert(data []byte) {
	if sbf.head.ApproxP() > sbf.p {
		sbf.addBF()
	}
	sbf.head.Insert(data)
}
