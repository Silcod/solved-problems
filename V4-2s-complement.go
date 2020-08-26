// package main

// import (
// 	"fmt"
// 	"math"
// 	"math/bits"
// 	"strconv"
// 	"log"
// )

// var (
// 	rangeMinimum   int32
// 	rangeMaximum   int32
// 	numberOfOnes int32
// 	oneCount int32
// 	twosComplement int64
// 	onesComplement int64
// 	isNegative     bool
// 	validity bool
// )

// func findTwosComplement(number int32) int64 {
// 	positiveValueOfNumber := math.Abs(float64(number))
// 	onesComplement = int64(((1 << 32) - 1) ^ int(positiveValueOfNumber))
// 	twosComplement = onesComplement + 1
// 	// onesComplementBinaryRepresentation := strconv.FormatInt(int64(twosComplement), 2)

// 	return twosComplement
// }

// func checkDataValidity(rangeMinimum, rangeMaximum int32) bool {
// 	negativeBasePower := int32(math.Pow(-2, 31))
// 	basePowerMinusOne := int32(math.Pow(2, 31) - 1)
// 	if rangeMinimum >= negativeBasePower && rangeMaximum >= rangeMinimum && rangeMaximum <= basePowerMinusOne {
// 		fmt.Println("Valid data for range")
// 		validity = true
// 	} else {
// 		log.Fatal("Invalid Data for range")
// 	}

// 	return validity
// }

// func main() {
// 	rangeMinimum = -1
// 	rangeMaximum = 4
// 	numberOfOnes = 0

// 	validityResult := checkDataValidity(rangeMinimum, rangeMaximum)

// 	if validityResult == true {
// 		for i := rangeMinimum; i <= rangeMaximum; i++ {
// 			isNegative = math.Signbit(float64(i))
// 			if isNegative == true {
// 				twosComplementRepresentation := findTwosComplement(i)
// 				oneCount = int32(bits.OnesCount(uint(twosComplementRepresentation)))
// 				numberOfOnes += oneCount
// 			} else {
// 				binaryRepresentation := strconv.FormatUint(uint64(i), 2)
// 				binaryRepresentationInUint, err := strconv.ParseUint(string(binaryRepresentation), 2, 32)
// 				if err != nil {
// 					fmt.Println(err)
// 				}
// 				oneCount = int32(bits.OnesCount(uint(binaryRepresentationInUint)))
// 				numberOfOnes += oneCount
// 			}
// 		}
// 	} else {
// 		log.Fatal("Invalid Data for range")
// 	}

// 	fmt.Println(numberOfOnes)
// }
