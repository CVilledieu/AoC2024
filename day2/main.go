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
		if checkList(list) {
			answer = append(answer, list)
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

func checkList(list []int) bool {
	if list[0] < list[1] {
		for i := range list {
			if i+1 > len(list)-1 {
				return true
			}
			left := list[i]
			right := list[i+1]
			if left >= right {
				return false
			}
			if right-left > 3 {
				return false
			}
		}
	} else if list[0] > list[1] {
		for i := range list {
			if i+1 > len(list)-1 {
				return true
			}
			left := list[i]
			right := list[i+1]
			if left <= right {
				return false
			}
			if left-right > 3 {
				return false
			}
		}
	}
	return false
}
