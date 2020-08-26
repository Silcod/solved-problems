// package main

// import (
// 	"fmt"
// 	"math"
// 	"math/bits"
// 	"strconv"
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

// 		fmt.Println(bits.LeadingZeros32(uint32(uintValue)))
// 		fmt.Println(bits.OnesCount(uint(uintValue))) // function to find the number of one's
// 		fmt.Println(bits.Reverse(uint(uintValue)))
// 		fmt.Printf("%032s", binaryPositiveValue)
// 	}

// }
