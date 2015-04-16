# Blumo

CS51 Final Project. Optimizing scaled bloom filters for use with sets of strings of arbitrary size. 


Overview:
	A Bloom filter is a probabilistic data structure that stores elements with a constant bit/element ratio. Thus, a Bloom filter has several advantages: it has a constant search time, and can store a large number inputs in a relatively small amount of memory. One of the downsides is that a Bloom filter can return a false positive, but we can easily decrease the false positive percentage to a desired degree. 

	We want to implement a basic Bloom filter, then explore how different hash functions effect performance. Then, we will generate hash functions based on a model hash function by modifying constants and use these hash functions to implement a Bloom filter of variable size. Finally, we will implement a scalable Bloom filter. 

	The advantage of a scalable Bloom filter is that it can take in a stream of inputs, as opposed to requiring a a definite input size a priori. 

I) Create Bloom filters of static size. 
	a) Implement a Bloom filter with static size (m), static false positive percentage (p), static number of hash functions (k), and a static number of inputs (n). 
	b) Implement search and insert functions.
II) Implement a function that creates a Bloom filter of an optimal size given n and p. 
	a) Investigate common hash functions - looking for speed, distribution. Select one that performs well and is easily modified, such that modifying the constants does not significantly alter performance.
	b) Implement a function that generates hash functions similar to our chosen model. 
	c) Implement a function that optimizes m and k given a particular n and p (this can be done through some probability calculations) and creates a Bloom filter with size m, k. 
III) Wrap bloom filters to make them scaleable. 
	a) Take in a stream of inputs, create additional Bloom filters as preceding Bloom filters are full (full-ness is determined by whether or not the preceding Bloom filter has hit the desired false-positive percentage).
	b) Implement search and insert functions. 


Detailed Description
 - i.e. a specific outline of the functions we're writing 

 Signatures/Interfaces
 - i.e. abstractions
 - given a certain function we are writing, what sort of abstraction barrier do we want it to have, what are it's input outputs etc. 
 

 Modules/Actual Code
 - exactly what it sounds like


 Timeline:
 Week of 4/11/15
 - determine implementation language: Google Go
 - plan out final specification
 - plan out functions/interfaces/higher level design
 

 Week of 4/18/15
 - implement static Bloom filter
 - implement lookup and insert for a static Bloom filter
 - implement hash function generating function
 - implement a Bloom filter that is created to be optimal for a given n and a given k
 - determine at what n and p threshold we establish that a Bloom filter is "full"
 - create, essentially a "set" of Bloom filters that takes in a stream of inputs and self-expands to accomodate an unknown number of inputs
 - implement a search for this Scalable Bloom filter


Progress Report
- carefully understood how Bloom filters work and what factors effect the effectiveness of a Bloom filter
- outlined our game plan
- decided to use Google Go after researching implementations of Bloom filters and discussing with our TF
- discuss potential extensions of our project (i.e. optimization, Scalable Bloom filters, working on an application)
- started learning Go


Version Control
- git repository