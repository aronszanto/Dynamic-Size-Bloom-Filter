/*

 This file implements the extension for our CS51 final project.

 Joseph Kahn    josephkahn@college.harvard.edu
 Grace Lin      glin@college.harvard.edu
 Aron Szanto    aszanto@college.harvard.edu

*/

package ScalableFilterPartition //MOVE M K ETC SP ETC TO INTERFACE

import (
	//"fmt"
	"github.com/aszanto9/Blumo/staticfilterpartition"
	"math"
)

type SBF struct {
	// head points to the current filter that is not yet full
	head *StaticFilterPartition.Filter
	// an array of pointers to all of the existing filters
	filter_slice                                  []*StaticFilterPartition.Filter
	scaling_factor, N, n_init                     uint
	total_err_bound, tightening_ratio, fill_ratio float64
}

func NewFilter(end_e float64) *SBF {
	//default values for s, r (hardcoded)
	n_init_i := uint(10000)
	scaling_factor_i := uint(2) // factor by which each additional static filter increases in size
	fill_ratio_i := 0.05        // approximation for ratio of 1s to size of bitset
	N_i := uint(1)              // number of static filters in SBF
	tightening_ratio_i := 0.8
	head_i := StaticFilterPartition.NewFilter(uint(n_init_i), end_e*(1-tightening_ratio_i))
	return &SBF{
		n_init:           n_init_i,
		scaling_factor:   scaling_factor_i,
		N:                N_i,
		total_err_bound:  end_e,
		fill_ratio:       fill_ratio_i,
		tightening_ratio: tightening_ratio_i,
		head:             head_i,
		filter_slice:     []*StaticFilterPartition.Filter{head_i},
	}
}

func (sbf *SBF) Lookup(data []byte) bool {

	for i := range sbf.filter_slice {
		//fmt.Printf(fmt.Sprint("Looking for ", data, " in filter #", i, "\n"))
		if sbf.filter_slice[i].Lookup(data) {
			//fmt.Printf(fmt.Sprint("Found in filter ", i, "\n"))
			return true
		}
	}
	return false

}

func (sbf *SBF) addBF() {
	sbf.N++
	newfilter := StaticFilterPartition.NewFilter(sbf.head.N()*sbf.scaling_factor,
		sbf.total_err_bound*math.Pow(sbf.tightening_ratio, float64(sbf.N-1)))
	sbf.head = newfilter
	sbf.filter_slice = append(sbf.filter_slice, newfilter)

}

func (sbf *SBF) Insert(data []byte) {
	if sbf.head.ApproxP() > sbf.fill_ratio {
		//fmt.Printf(fmt.Sprint("Approx fill of filter with ", sbf.head.M(), " bits, of which ", sbf.head.BitsFlipped(), " are flipped, is ", sbf.head.ApproxP(), " so adding new filter\n"))

		sbf.addBF()
	}
	sbf.head.Insert(data)
}
