package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
create a map and then an array of numbers that you should not come across

Task:
Find which arrays fit
For all the arrays that are correct find the middle
sum all the middles
*/
func main() {
	var sum int
	rules, arr := getInput()
	for _, li := range arr {
		sum += checkArray(rules, li)
	}
	fmt.Println(sum)
}

func checkArray(m map[int][]int, arr []int) int {
	var middle int
	noFlyList := make(map[int]bool)
	fmt.Println(arr)
	for _, n := range arr {
		if noFlyList[n] {
			return 0
		}
		for _, i := range m[n] {
			noFlyList[i] = true
		}
	}
	middle = arr[(len(arr) / 2)]
	fmt.Println(middle)
	return middle
}

func getInput() (map[int][]int, [][]int) {
	newMap := createMap()
	scan := getScanner()
	var arr [][]int
	for {
		EoF := !(scan.Scan())
		if EoF {
			break
		}
		curLine := scan.Text()
		//There is a blank line between the 2 different types of info
		//Once that blank line is hit we process the second
		if curLine == "" {
			break
		}
		value, key := getRules(curLine)
		newMap[key] = append(newMap[key], value)

	}
	for {
		EoF := !(scan.Scan())
		if EoF {
			break
		}
		curLine := getArr(scan.Text())
		arr = append(arr, curLine)
	}

	return newMap, arr
}

func getArr(s string) []int {
	var arr []int
	remaining := s
	for {
		l, r, b := strings.Cut(remaining, ",")
		num, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		arr = append(arr, num)
		if !b {
			break
		}
		remaining = r
	}
	return arr
}

func getRules(s string) (int, int) {
	var left, right int

	l, r, b := strings.Cut(s, "|")
	if b == false {
		panic("Cut rules err")
	}

	left, err := strconv.Atoi(l)
	if err != nil {
		panic(err)
	}
	right, err = strconv.Atoi(r)
	if err != nil {
		panic(err)
	}

	return left, right
}

func createMap() map[int][]int {
	m := make(map[int][]int)
	return m
}

func getScanner() *bufio.Scanner {
	file, err := os.OpenFile("input.txt", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return bufio.NewScanner(file)
}
