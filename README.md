# BloomFilter Data Structure in golnag

In this repo I show how to implements bloom filter using both boolean and byte example

## What is BloomFilter?

A Bloom filter is a space-efficient, probabilistic data structure used to determine if an element is likely a member of a set. It uses a bit array and multiple hash functions to represent the set, offering fast membership testing with a chance of false positives. While it can definitively confirm an element's absence, it may incorrectly indicate an element's presence.

## Understanding the Bit Manipulation In Bit-Level Storage (Byte)

    // Sets the bit at position func  (bf *BloomFilter)  setBit(position uint)  { bytePos := position /  8  // Which byte contains our bit bitPos := position %  8  // Which bit within that byte (0-7) // Use bitwise OR with a single bit set at bitPos bf.bitArray[bytePos]  |=  1  << bitPos } // Gets the bit at position func  (bf *BloomFilter)  getBit(position uint)  bool  { bytePos := position /  8 bitPos := position %  8 // Use bitwise AND to check if the bit is set return  (bf.bitArray[bytePos]  &  (1  << bitPos))  !=  0 }

### Helpful Links

[Arpit Bhayani Youtube Video Link](https://www.youtube.com/live/UVFnabieyzc?si=QyDxfm6AwXCNI9zA)
[Creating a Bloom Filter with Go By Dylan Meeus](https://medium.com/@meeusdylan/creating-a-bloom-filter-with-go-7d4e8d944cfa)
