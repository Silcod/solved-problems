// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"unsafe"
// )

// func main() {
// 	var rangeMinimum int32
// 	// var rangeMaximum int64
// 	rangeMinimum = -2
// 	// rangeMaximum = 3

// 	// base2 := strconv.FormatUint(uint64(rangeMinimum), 2)
// 	base2 := strconv.FormatUint(uint64(*(*uint32)(unsafe.Pointer(&rangeMinimum))), 2)
// 	fmt.Println(base2)
// }
