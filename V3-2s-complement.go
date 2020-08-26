// package main

// import (
// 	"fmt"
// 	"math"
// 	"math/bits"
// 	"strconv"
// 	"strings"
// 	"unsafe"
// )

// func main() {
// 	var rangeMinimum int32
// 	// var rangeMaximum int64
// 	rangeMinimum = -8
// 	// rangeMaximum = 3

// 	// base2 := strconv.FormatUint(uint64(rangeMinimum), 2)
// 	base2 := strconv.FormatUint(uint64(*(*uint32)(unsafe.Pointer(&rangeMinimum))), 2)
// 	fmt.Println(base2)

// 	// Write a function to find the 2's complement of number if it is negative
// 	isNegative := math.Signbit(float64(rangeMinimum))
// 	if isNegative == true {
// 		positiveValue := math.Abs(float64(rangeMinimum))
// 		binaryPositiveValue := strconv.FormatUint(uint64(positiveValue), 2)
// 		uintValue, err := strconv.ParseUint(binaryPositiveValue, 2, 32)
// 		if err != nil {
// 			fmt.Println(err)
// 		}

// 		fmt.Println(bits.OnesCount(uint(uintValue))) // function to find the number of one's
// 		// fmt.Printf("%032s", binaryPositiveValue)
// 		fmt.Printf("%q\n", strings.Split(binaryPositiveValue, ""))
// 		fmt.Println(bits.Len32(uint32(positiveValue)))
// 		fmt.Println(bits.LeadingZeros32(uint32(positiveValue)))

// 		onesComplement := ((1 << 32) - 1) ^ 22
// 		onesComplementBinaryRepresentation := strconv.FormatUint(uint64(4294967273), 2)
// 		intValue, err := strconv.ParseUint(onesComplementBinaryRepresentation, 2, 32)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		var one uint32
// 		one = 1
// 		twosComplement, _ := bits.Add32(uint32(intValue), one, 0)
// 		fmt.Println(onesComplement)
// 		fmt.Println(twosComplement)
// 		fmt.Println(onesComplementBinaryRepresentation)

// 	}

// }
