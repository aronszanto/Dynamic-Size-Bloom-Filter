/*

 This file implements the foundation for our CS51 final project.

 Joseph Kahn    josephkahn@college.harvard.edu
 Grace Lin      glin@college.harvard.edu
 Aron Szanto    aszanto@college.harvard.edu

*/

package ScalableFilter

import (
 	"github.com/willf/bitset"
	"hash"
	"hash/fnv"
	"math"
	"StaticFilter"
)



type SBF struct {
	// head points to the current filter that is not yet full
	head *StaticFilter.Filter 
	headcap int64
	// an array of pointers to all of the existing filters
	filter_slice []*StaticFilter.Filter
	// s is the scaling factor for the size of new filters, N is the number of existing filters
	// m_init is the m of the first filter in an SBF
	s, N, m_init int64
	// p is the total final error bound, r is the scaling factor for the error bound of new filters
	p, r float64
}

type SBF interface {
	NewSBF SBF
	SBFlookup bool
	SBFinsert 
	NewBF Filter
}

func NewSBF(p float64) SBF {
	//default values for s, r (hardcoded)
	s := 2
	r := .5
	m_init := 100
	p_init := p * (1-r)
	headcap := int64(float64(m_init) * math.Log(2))
	filter1 := StaticFilter.NewFilter(m_init, p_init)
	return SBF{&filter1, headcap, [&filter1], s, 1, m_init, p, r}
}

func (sbf *SBF) SBFlookup(data []byte) bool {
	ispresent = false 
	for i:= 0; i < sbf.N; i++ {
		if (sbf.filter_slice[i]).Lookup(data) == true {
			ispresent = true
			break
		}
		else continue
	}
}

func AddBF(sbf *SBF) SBF {
	newfilter := StaticFilter.NewFilter(sbf.head.m * sbf.s, sbf.head.p * sbf.r)
	filter_slice = append(sbf.filter_slice, &newfilter)
	head = &newfilter 
	headcap = sbf.headcap * sbf.s
	return SBF{&newfilter, headcap, filter_slice, s, N + 1, m_init, p, r}
}

func (sbf *SBF) SBFinsert(data []byte) {
	if sbf.filter_slice[sbf.N-1].counter < sbf.headcap {
		sbf.filter_slice[sbf.N-1].Insert(data)
	}
	else {
		AddBF(sbf)
		sbf.filter_slice[sbf.N-1].Insert(data)
	}
}



func main {

}