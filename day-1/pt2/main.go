package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	dataSlice, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var strValues []string = strings.Fields(string(dataSlice))

	var col1Str []string
	var col2Str []string
	for i := 0; i < len(strValues); i++ {
		if i%2 == 0 {
			col1Str = append(col1Str, strValues[i])
		} else {
			col2Str = append(col2Str, strValues[i])
		}
	}
	var col1 []uint32
	for i := 0; i < len(col1Str); i++ {
		j, err := strconv.Atoi(col1Str[i])
		if err != nil {
			panic(err)
		}
		col1 = append(col1, uint32(j))
	}

	var col2 []uint32
	for i := 0; i < len(col2Str); i++ {
		j, err := strconv.Atoi(col2Str[i])
		if err != nil {
			panic(err)
		}
		col2 = append(col2, uint32(j))
	}

	slices.Sort(col1)
	slices.Sort(col2)
	fmt.Printf("index: %v ", len(col1))
	fmt.Printf("index: %v\n", len(col2))

	var scoreArr []uint32
	for i := 0; i < len(col1); i++ {
		var score uint32 = col1[i]
		var scoreMultiplier int
		for j := 0; j < len(col2); j++ {
			if col1[i] == col2[j] {
				scoreMultiplier++
			}
		}

		scoreArr = append(scoreArr, score*uint32(scoreMultiplier))
	}

	var sum uint32
	for i := 0; i < len(scoreArr); i++ {
		sum += scoreArr[i]
	}

	fmt.Printf("%v\n", sum)
}
