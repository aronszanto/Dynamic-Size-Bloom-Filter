# Blumo

CS51 Final Project. Optimizing scaled bloom filters for use with sets of strings of arbitrary size. 


Overview:
	A Bloom filter is a probabilistic data structure that stores elements with a constant bit/element ratio. Thus, a Bloom filter has several advantages: it has a constant search time, and can store a large number inputs in a relatively small amount of memory. 

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
bloom filter
 

 Modules/Actual Code
 - exactly what it sounds like


 Timeline:
 Week of 4/11/15
 - determine implementation language: Google Go
 - plan out functions

Progress Report
- what have we done in the last week?

Version Control
- git repository