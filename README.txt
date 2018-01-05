# Dynamic Size Bloom Filter

## Research Abstract

A Bloom filter is a probabilistic data structure that stores elements with a constant bit/element ratio. Thus, a Bloom filter has several advantages: it has a nearly constant search time, and can store a large number inputs in a relatively small amount of memory. One of the disadvantages of traditional Bloom filters, however, is that the input size is required ahead of insertion.

In this work, we develop a scalable Bloom filter, inspired by a paper by Almeida et al., that can take in a stream of inputs, as opposed to requiring a a definite input size a priori. Testing on commodity hardware gives performance of 400K inserts/sec and 1M lookups/sec under 0.001% false positive tolerance.
