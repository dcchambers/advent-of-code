package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Open input file
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//read input file into a Slice (type:string)
	var diagnostics []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		diagnostics = append(diagnostics, scanner.Text())
	}

	//find the most and least common bit in each
	var mostCommonBits []string
	var leastCommonBits []string
	numBits := len(diagnostics[0])
	for i := 0; i < numBits; i++ {
		mostCommonBits = append(mostCommonBits, findMostCommonBit(diagnostics, i))
		leastCommonBits = append(leastCommonBits, findLeastCommonBit(diagnostics, i))
	}

	//calculate Gamma Rate and Epsilon Rate using leastCommonBits and mostCommonBits
	//(converts binary to decimal)
	gr := calculateGammaRate(strings.Join(mostCommonBits, ""))
	er := calculateEpsilonRate(strings.Join(leastCommonBits, ""))

	fmt.Println(gr * er)

}

/*
parameters
s []string - A string slice/array to analyze( eg. an array containing: "0000", "0001", "1010", "1101", "0010", "0100")
x int - which location in the string to perform the analysis
*/
func findMostCommonBit(s []string, x int) string {
	numZeroBits := 0
	numOneBits := 0
	for i := 0; i < len(s); i++ {
		str := s[i]
		//DEBUGfmt.Println(str)
		if string(str[x]) == "0" {
			numZeroBits++
		} else if string(str[x]) == "1" {
			numOneBits++
		}
	}

	if numZeroBits > numOneBits {
		return "0"
	} else {
		return "1"
	}

}

//super inefficient to do this again instead of just bit flipping the most common one but oh well this is more fun
func findLeastCommonBit(s []string, x int) string {
	numZeroBits := 0
	numOneBits := 0
	for i := 0; i < len(s); i++ {
		str := s[i]
		if string(str[x]) == "0" {
			numZeroBits++
		} else if string(str[x]) == "1" {
			numOneBits++
		}
	}

	if numZeroBits < numOneBits {
		return "0"
	} else {
		return "1"
	}
}

func calculateGammaRate(s string) int64 {
	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func calculateEpsilonRate(s string) int64 {
	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return i
}
