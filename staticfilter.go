/* 

 This file implements the basis for our CS51 final project. 

 Joseph Kahn    josephkahn@college.harvard.edu
 Grace Lin      glin@college.harvard.edu
 Aron Szanto    aszanto@college.harvard.edu

 */

package main 

// Including library packages referenced in this file 
import (
    //"fmt"
    "math"
)


func main() {

}

// type definition for standard bloom filter
type Bloom struct{
    eps float64 // false positive error bound
    n uint64 // predicted number of elements in filter
    m uint64 // size of bit array
    //hf hash_function // to generate list of bites values h(0)…h(k). type: 
    k uint64 // number of hash functions
    // bset BitSet // will use external library for this
}

// initialize key values using math!!
func new (num uint64, err_bound float64) Bloom {
    eps := err_bound
    n := num
    m := uint64(math.Ceil(-1 * (float64(n) * math.Log(eps)) / ((math.Log(2) * math.Log(2)))))
    //hf := Hash.hash_fun // will use external hash function
    k := uint64(math.Ceil((float64(m) / float64(n)) * math.Log(2)))
    //bset = BitSet.new(m) // will use external bites implementation
    return Bloom{eps, n, m, k}
}

// basic methods. may later include methods for changing epsilon, estimating how many elements are in filter, checking how saturated filter is, etc.

//func insert (key string) Bloom
//func check (key string) bool
//func reset()
