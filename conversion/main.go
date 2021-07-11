package main

import (
	"fmt"
	"math"
)

func main() {
	var intdeger int = math.MaxInt8
	var int32deger int32 = math.MaxInt32
	var int64deger int64 = math.MaxInt64
	var intdegerfor64 int = math.MaxInt64
	fmt.Printf("intdeger deger tipi : %T\n", intdeger)
	fmt.Printf("int32deger deger tipi : %T\n", int32deger)
	fmt.Printf("int64deger deger tipi : %T\n", int64deger)
	fmt.Printf("intdegerfor64 deger tipi : %T\n", intdegerfor64)
}
