package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	reports := getReports()
	var answer [][]int
	for _, list := range reports {
		//Check for early return
		if list[0] == list[1] {
			continue
		}
		firstCheck := verifyOrder(list)
		if firstCheck == -1 {
			answer = append(answer, list)
		} else {
			left := list[:firstCheck]
			right := list[firstCheck+1:]

			new := make([]int, len(left)+len(right))
			for i := range left {
				new[i] = left[i]
			}
			for j := range right {
				new[j+len(left)] = right[j]
			}

			secondCheck := verifyOrder(new)

			if secondCheck == -1 {
				answer = append(answer, new)

			} else {

				finalList := list[1:]
				finalCheck := verifyOrder(finalList)
				if finalCheck == -1 {
					answer = append(answer, list)
				}
			}

		}

	}

	fmt.Println(len(answer))
}

func getReports() [][]int {
	var b [][]int
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
		original := scanner.Bytes()
		list := convertToNumber(original)
		b = append(b, list)
	}
	return b
}

// subtract "Zero" or 48 from the byte to get the number.
// If next byte is 32 return the number
// If next byte is not then multiply current value by 10 to get ready for next number
func convertToNumber(bytes []byte) []int {

	var list []int
	var current int
	offset := byte(48)
	for i := 0; i < len(bytes); i++ {
		num := bytes[i]
		if num != 32 {
			current *= 10
			current += int(num - offset)
		} else {
			list = append(list, current)
			current = 0
		}

	}
	list = append(list, current)
	return list
}

func checkList(list []int) int {
	var orderErr int
	for i := range list {
		if i == len(list) {
			break
		}
		left := list[i]
		right := list[i+1]

	}
}

//-----------------------------------------
//------------ PREVIOUS ATTEMPTS-----------
//-----------------------------------------
/*
// Treating the values like a graph. Position in the array is the X value and value in that spot is the Y value
func verifyOrder(list []int) int {
	// Direction of the slope. Positive slope means the graph is increasing
	var originDirection, direction, delta int
	maxDelta := 3
	originDirection, _ = checkDelta(list[0], list[1])
	listLen := len(list)
	//Loop will check last element against the element before it
	for i := 0; i < listLen-1; i++ {
		j := i + 1
		direction, delta = checkDelta(list[i], list[j])
		if delta == 0 {
			return j
		}
		if originDirection != direction || delta > maxDelta {
			return j
		}
	}

	return -1
}

func checkDelta(left, right int) (int, int) {
	if left > right {
		delta := left - right
		return 1, delta
	} else {
		delta := right - left
		return -1, delta
	}
}
*/
