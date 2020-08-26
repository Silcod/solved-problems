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
// 	isNegativeNumber     bool
// 	validity bool
// )

// func findTwosComplement(number int32) int64 {
// 	positiveValueOfNumber := math.Abs(float64(number))
// 	onesComplement = int64(((1 << 32) - 1) ^ int(positiveValueOfNumber))
// 	twosComplement = onesComplement + 1

// 	return twosComplement
// }

// func checkDataValidity(rangeMinimum, rangeMaximum int32) bool {
// 	negativeBasePower := int32(math.Pow(-2, 31))
// 	basePowerMinusOne := int32(math.Pow(2, 31) - 1)
// 	if rangeMinimum >= negativeBasePower && rangeMinimum <= rangeMaximum && rangeMaximum <= basePowerMinusOne {
// 		validity = true
// 	} else {
// 		log.Fatal("Invalid Data for range")
// 	}

// 	return validity
// }

// func calculateTheTwosComplement(rangeMinimum, rangeMaximum int32) int32 {
// 	numberOfOnes = 0
// 	validityResult := checkDataValidity(rangeMinimum, rangeMaximum)
// 	if validityResult == true {
// 		for i := rangeMinimum; i <= rangeMaximum; i++ {
// 			isNegativeNumber = math.Signbit(float64(i))
// 			if isNegativeNumber == true {
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

// 	return numberOfOnes
// }

// func main() {
// 	rangeMinimum = -1
// 	rangeMaximum = 4
	
// 	result := calculateTheTwosComplement(rangeMinimum, rangeMaximum)
// 	fmt.Println(result)
// }