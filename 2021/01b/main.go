package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	data, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	defer data.Close()

	// Read the input file into a Slice
	var dataArr []int
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		value, err2 := strconv.Atoi(scanner.Text())
		if err2 != nil {
			panic(err2)
		}
		dataArr = append(dataArr, value)
	}

	// Add up Three Measure Sums
	var sumArr []int
	threeMeasureSum := 0
	for i := 0; i < (len(dataArr) - 2); i++ {
		threeMeasureSum = dataArr[i] + dataArr[i+1] + dataArr[i+2]
		sumArr = append(sumArr, threeMeasureSum)
	}

	// Count the number of times each subsquent element in the slice increases
	count := 0
	for j := 1; j < len(sumArr); j++ {
		if sumArr[j] > sumArr[j-1] {
			count++
		}
	}

	fmt.Println(count)
}
