package main

import (
	"fmt"
	"log"
	"math"
	"math/bits"
	"strconv"
)

var (
	rangeMinimum      int64
	rangeMaximum      int64
	numberOfOnes      int32
	oneCount          int32
	twosComplement    int64
	onesComplement    int64
	isNegativeInteger bool
	validity          bool
)

func checkDataValidity(rangeMinimum, rangeMaximum int64) bool {
	negativeBasePower := int64(math.Pow(-2, 31))
	basePowerMinusOne := int64(math.Pow(2, 31) - 1)
	if rangeMinimum >= negativeBasePower && rangeMinimum <= rangeMaximum && rangeMaximum <= basePowerMinusOne {
		validity = true
	} else {
		log.Fatal("Invalid Data for range")
	}

	return validity
}

func findTwosComplement(number int64) int64 {
	positiveValueOfNumber := math.Abs(float64(number))
	onesComplement = int64(((1 << 32) - 1) ^ int(positiveValueOfNumber))
	twosComplement = onesComplement + 1

	return twosComplement
}

func negativeIntegerNumberOfOnes(number int64) {
	twosComplementRepresentation := findTwosComplement(number)
	oneCount = int32(bits.OnesCount(uint(twosComplementRepresentation)))
	numberOfOnes += oneCount
}

func positiveIntegerNumberOfOnes(number int64) {
	binaryRepresentation := strconv.FormatUint(uint64(number), 2)
	binaryRepresentationInUint, err := strconv.ParseUint(string(binaryRepresentation), 2, 32)
	if err != nil {
		fmt.Println(err)
	}
	oneCount = int32(bits.OnesCount(uint(binaryRepresentationInUint)))
	numberOfOnes += oneCount
}

func findTotalNumberOfOnes(number int64) int32 {
	isNegativeInteger = math.Signbit(float64(number))
	if isNegativeInteger == true {
		negativeIntegerNumberOfOnes(number)
	} else {
		positiveIntegerNumberOfOnes(number)
	}
	return numberOfOnes
}

func calculateNumberOfOnesIn2sComplement(rangeMinimum, rangeMaximum int64) int32 {
	numberOfOnes = 0
	validityResult := checkDataValidity(rangeMinimum, rangeMaximum)
	if validityResult == true {
		for i := rangeMinimum; i <= rangeMaximum; i++ {
			numberOfOnes = findTotalNumberOfOnes(i)
		}
	} else {
		log.Fatal("Invalid Data for range")
	}

	return numberOfOnes
}

func main() {
	result := calculateNumberOfOnesIn2sComplement(-2147483648, 2147483647)
	fmt.Println(result)
}
