package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//Open input file
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//Read input file into a Slice
	var commands []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		commands = append(commands, scanner.Text())
	}

	//Determine horizontal and vertical movement of the ship.
	horizontalDistance, depth, aim := 0, 0, 0
	for i := 0; i < len(commands); i++ {
		directionAndDistance := strings.Fields(commands[i])

		if directionAndDistance[0] == "up" {
			distance, err := strconv.Atoi(directionAndDistance[1])
			if err != nil {
				panic(err)
			}
			aim -= distance
		} else if directionAndDistance[0] == "down" {
			distance, err := strconv.Atoi(directionAndDistance[1])
			if err != nil {
				panic(err)
			}
			aim += distance
		} else if directionAndDistance[0] == "forward" {
			distance, err := strconv.Atoi(directionAndDistance[1])
			if err != nil {
				panic(err)
			}
			horizontalDistance += distance
			depth += (aim * distance)
		} else {
			panic("Unexpected directional input.")
		}
	}

	//Return desired result.
	fmt.Println(horizontalDistance * depth)

}
