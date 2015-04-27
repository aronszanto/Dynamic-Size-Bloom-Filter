/*

 This file implements the extension for our CS51 final project.

 Joseph Kahn    josephkahn@college.harvard.edu
 Grace Lin      glin@college.harvard.edu
 Aron Szanto    aszanto@college.harvard.edu

*/

package ScalableFilter

import (
	"github.com/aszanto9/Blumo/StaticFilter"
	"math"
)

type SBF struct {
	// head points to the current filter that is not yet full
	head    *StaticFilter.Filter
	headcap int64
	// an array of pointers to all of the existing filters
	filter_slice []*StaticFilter.Filter
	// s is the scaling factor for the size of new filters, N is the number of existing filters
	// m_init is the m of the first filter in an SBF
	s, N, m_init int64
	// p is the total final error bound, r is the scaling factor for the error bound of new filters
	p, r  float64
	count int64
}

/*type SBF interface {
	NewSBF SBF
	SBFlookup bool
	SBFinsert
	NewBF Filter
}*/

func NewSBF(end_p float64) *SBF {
	//default values for s, r (hardcoded)
	m_init_i := 100
	s_i := 2
	N_i := 1
	r_i := 0.5
	head_i := StaticFilter.NewFilter(m_init_i, p*(1-r_i))
	return &SBF{
		m_init:       m_init_i,
		s:            s_i,
		N:            N_i,
		p:            end_p,
		r:            r_i,
		head:         head_i,
		headcap:      int64(float64(m_init_i) * math.Log(2)),
		filter_slice: []*SBF{head_i},
		counter:      0,
	}
}

func (sbf *SBF) SBFlookup(data []byte) bool {
	for i := range sbf.filter_slice {
		if sbf.filter_slice[i].Lookup(data) {
			return true
		}
	}
	return false

}

// maybe insert should simply mutate the existing SBF, not return a completely new one...?

func (sbf *SBF) AddBF() SBF {
	newfilter := StaticFilter.NewFilter(sbf.head.m*sbf.s, sbf.head.p*sbf.r)

	return &SBF{
		head:         newfilter,
		headcap:      sbf.headcap * sbf.s,
		filter_slice: append(sbf.filter_slice, newfilter),
		s:            sbf.s,
		N:            sbf.N + 1,
		m_init:       sbf.m_init,
		p:            sbf.p,
		r:            sbf.r,
		counter:      0,
	}
}

func (sbf *SBF) SBFinsert(data []byte) {
	if sbf.filter_slice[sbf.N-1].counter < sbf.headcap {
		sbf.AddBF()
	}

	sbf.filter_slice[sbf.N-1].Insert(data)
	(sbf.filter_slice[sbf.N-1].counter)++
}
