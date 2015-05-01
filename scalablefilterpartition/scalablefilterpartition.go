/*

 This file implements the extension for our CS51 final project, building upon what we learned
 from the non-partionined filters, and upon the foundation laid out in staticfilterpartition.go.

 As such, all of the function calls like insert and lookup are the same ones used in static filters, except
 that they must be called on the correct static filters within the scalable filter.

 Joseph Kahn    josephkahn@college.harvard.edu
 Grace Lin      glin@college.harvard.edu
 Aron Szanto    aszanto@college.harvard.edu

*/

package ScalableFilterPartition

import (
	"github.com/aszanto9/Blumo/staticfilterpartition"
	"math"
)

type SBF struct {
	// head points to the current filter that is not yet full
	head *StaticFilterPartition.Filter
	// an array of pointers to all of the existing filters
	filter_slice     []*StaticFilterPartition.Filter
	scaling_factor   uint    // Factor which determines increase in size of each subsequent static filter
	N                uint    // Number of bloom filters in scalable filter
	n_init           uint    // initial
	total_err_bound  float64 // cumulative false-positive rate of static filters in scalable
	tightening_ratio float64 // Factor by which the subsequent error bound decreases
	fill_ratio       float64 // approximation for ratio of 1s to size of bitset
}

// Scaling factor and tightening ratio are hard-coded
func NewFilter(end_e float64) *SBF {
	n_init_i := uint(10000)
	scaling_factor_i := uint(2)
	fill_ratio_i := 0.05
	N_i := uint(1)
	tightening_ratio_i := 0.8
	head_i := StaticFilterPartition.NewFilter(uint(n_init_i), end_e*(1-tightening_ratio_i))
	return &SBF{
		n_init:           n_init_i,
		scaling_factor:   scaling_factor_i,
		N:                N_i,
		total_err_bound:  end_e,
		fill_ratio:       fill_ratio_i,
		tightening_ratio: tightening_ratio_i,
		// Keeps track of current static filter in scalable
		head:         head_i,
		filter_slice: []*StaticFilterPartition.Filter{head_i},
	}
}

// Lookup must check through the slice of constitutent static filters
func (sbf *SBF) Lookup(data []byte) bool {
	for i := range sbf.filter_slice {
		fmt.Printf(fmt.Sprint("Looking for ", data, " in filter #", i, "\n"))
		if sbf.filter_slice[i].Lookup(data) {
			fmt.Printf(fmt.Sprint("Found in filter ", i, "\n"))
			return true
		}
	}
	return false
}

/*
 Helper function called once the fill-ratio of the current static filter
 is saturated (i.e. exceeds the optimized ratio )
*/
func (sbf *SBF) addBF() {
	sbf.N++
	newfilter := StaticFilterPartition.NewFilter(sbf.head.N()*sbf.scaling_factor,
		sbf.total_err_bound*math.Pow(sbf.tightening_ratio, float64(sbf.N-1)))
	sbf.head = newfilter
	sbf.filter_slice = append(sbf.filter_slice, newfilter)
}

/*
 Insert operates purely on the curent static filter provided its fill ratio has not been
 saturated.
*/
func (sbf *SBF) Insert(data []byte) {
	if sbf.head.ApproxP() > sbf.fill_ratio {
		// fmt.Printf(fmt.Sprint("Approx fill of filter with ", sbf.head.M(), " bits, of which ", sbf.head.BitsFlipped(), " are flipped, is ", sbf.head.ApproxP(), " so adding new filter\n"))
		sbf.addBF()
	}
	sbf.head.Insert(data)
}
