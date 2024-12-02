package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	dataSlice, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var strSlice []string = strings.Fields(string(dataSlice))
	var intSlice = []uint8{}

	for _, i := range strSlice {
		s, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		intSlice = append(intSlice, uint8(s))
	}

	unsafeReports := controlSafety(intSlice)
	fmt.Println(unsafeReports)
}

func controlSafety(slc []uint8) int {
	a := make([][]uint8, len(slc)/5)
	for i := range a {
		a[i] = make([]uint8, 5)
		for j := range a[i] {
			a[i][j] = slc[i*5+j]
		}
	}

	// for i := range a {
	// 	fmt.Printf("%v \n", a[i])
	// }

	unsafeReports := 0
	for i := range a {
		ascending := math.Signbit(float64(a[i][0]) - float64(a[i][1]))
		for j := range a[i] {
			unsafe := false
			if ascending != math.Signbit(float64(a[i][j])-float64(a[i][j])) {
				unsafe = true
			}
			if j == 4 {
				continue
			}
			difference := int(a[i][j]) - int(a[i][j+1])
			if difference > 3 || difference < -3 || difference == 0 {
				unsafe = true
			}

			if unsafe {
				unsafeReports++
				break
			}
		}

	}
	return len(a) - unsafeReports
}
