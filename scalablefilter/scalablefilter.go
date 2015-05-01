/*

 We started out with this, but were never able to get it to work perfectly.
 Without the partitioning we have since added, we got runtime errors for trying to
 flip bits outside the bounds of the slices we created.

 As we broadened the scope of our academic research we found that scaleable filters
 are seldom implemented without partitioning due to the low flexibility in the number
 of keys that can be inserted into the filter. As such, we chose to prioritise
 implementing the partioned filter over debugging this one.

 This code is purely included to show part of our development process, and no tests
 have been written for it.

 Joseph Kahn    josephkahn@college.harvard.edu
 Grace Lin      glin@college.harvard.edu
 Aron Szanto    aszanto@college.harvard.edu

*/

package ScalableFilter

import (
	"fmt"
	"github.com/aszanto9/Blumo/staticfilter"
	"math"
)

type SBF struct {
	// head points to the current filter that is not yet full
	head    *StaticFilter.Filter
	headcap uint
	// an array of pointers to all of the existing filters
	filter_slice []*StaticFilter.Filter
	// s is the scaling factor for the size of new filters, N is the number of existing filters
	// n_init is the n of the first filter in an SBF
	s, N, n_init uint // JESUS CHRIST FIX THIS LATER
	// p is the total final error bound, r is the scaling factor for the error bound of new filters
	p, r float64
}

func NewFilter(end_p float64) *SBF {
	//default values for s, r (hardcoded)
	n_init_i := uint(1000)
	s_i := uint(4)
	N_i := uint(1)
	r_i := 0.8
	head_i := StaticFilter.NewFilter(uint(n_init_i), end_p*(1-r_i))
	return &SBF{
		n_init:       n_init_i,
		s:            s_i,
		N:            N_i,
		p:            end_p,
		r:            r_i,
		head:         head_i,
		headcap:      uint(math.Ceil((float64(head_i.m) * math.Log(2)))),
		filter_slice: []*StaticFilter.Filter{head_i},
	}
}

func (sbf *SBF) Lookup(data []byte) bool {
	for i := range sbf.filter_slice {
		fmt.Printf(fmt.Sprint("Looking for ", data, " in filter #", i, "\n"))
		if sbf.filter_slice[i].Lookup(data) {
			return true
		}
	}
	return false

}

func (sbf *SBF) addBF() {
	newfilter := StaticFilter.NewFilter((sbf.head.M())*sbf.s, (sbf.head.E())*sbf.r)
	sbf.head = newfilter
	sbf.headcap *= sbf.s
	sbf.filter_slice = append(sbf.filter_slice, newfilter)
	sbf.N++
	fmt.Printf(fmt.Sprint("Bloom filter #", sbf.N, " added\n"))

}

func (sbf *SBF) Insert(data []byte) {
	if sbf.filter_slice[sbf.N-1].Counter < sbf.headcap {
		sbf.addBF()
	}
	sbf.filter_slice[sbf.N-1].Insert(data)
	(sbf.filter_slice[sbf.N-1].Counter)++
}
