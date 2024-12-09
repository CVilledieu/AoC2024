package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
--- TO DO ---
- Pair up smallest num in left list with smallest num in right list

- Find difference for each

- Find sum of diffences <-- answer to part 1

*/

func main() {
	left, right := getLists()
	sLeft := sort(left)
	sRight := sort(right)

	var sim int
	for _, n := range sLeft {
		sim += findSim(n, sRight)
	}
	fmt.Println(sim)
}
func getLists() (left, right []int) {
	file, err := os.OpenFile("input.txt", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	for {
		EoF := !(scanner.Scan())
		if EoF {
			break
		}
		l, r := splitPair(scanner.Text())
		left = append(left, l)
		right = append(right, r)
	}
	return
}

func findSim(num int, list []int) int {

	var sim int
	for _, n := range list {
		if n > num {
			return num * sim
		}
		if n == num {
			sim++
		}
	}

	return sim
}

func splitPair(b string) (left, right int) {
	rightIndex := len(b) - 5
	leftIndex := 5
	leftByte := b[:leftIndex]
	rightByte := b[rightIndex:]
	left, err := strconv.Atoi(leftByte)
	right, err = strconv.Atoi(rightByte)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func sort(list []int) []int {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}

	}
	return list
}
