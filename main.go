package main

import (
	"fmt"

	bloombytes "github.com/imPrathamDev/golang-bloom-filter/internal/bloom-bytes"
)



func main() {
	// BloomFilter example with bool
	// f := bloombool.NewBloomFilter(100)
	// f.Set("hello")
	// f.Set("world")
	// fmt.Println("is pratham exist", f.Check("pratham"))
	// fmt.Println("is hello exist", f.Check("hello"))

	f := bloombytes.NewBloom(1000, 3)
	f.Set("pratham")
	f.Set("hello")
	f.Set("test")
	f.Set("hiii")


	fmt.Println("is pratham", f.Check("pratham"))
	fmt.Println("is mahesh", f.Check("mahesh"))
	fmt.Println("is mohan", f.Check("mohan"))
	fmt.Println("is hi", f.Check("hi"))
	fmt.Println("is hello", f.Check("hello"))
	fmt.Println("is test", f.Check("test"))
	fmt.Println("is hiii", f.Check("hiii"))

}