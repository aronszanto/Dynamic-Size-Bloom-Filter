Project Description:

A Bloom filter is a probabilistic data structure that stores elements with a constant bit/element ratio. Thus, a Bloom filter has several advantages: it has a nearly constant search time, and can store a large number inputs in a relatively small amount of memory. One of the disadvantages of traditional Bloom filters, however, is that the input size is required ahead of insertion.

We want to implement a scalable Bloom filter, inspired by a paper by Almeida et al., that can take in a stream of inputs, as opposed to requiring a a definite input size a priori. We will then modularize the structure, creating a wrapper class bloom_i that represents all types of filters, allowing an abstraction barrier that streamlines the user's experience and tailors it to their needs.
