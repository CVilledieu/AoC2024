package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
Parse input of random character to find sequence that matches a set pattern
Pattern is mul(num,num)
num can be any 1 to 3 digit number
only that exact pattern will work

then sum the product of all pairs

byte codes
m = 109
u = 117
l = 108
( = 40
0 = 48
...
9 = 57
, = 44
) = 41

*/

func main() {
	getLines()

}

// mul(X,X)
// mul(XXX,XXX)
// index 4 has to be somewhere between 48 and 57
// index 5 through 10 could also be 48 through 57
// one of 5,6, or 7 have to be a 44
// one of 7 through 11 have to be a 41
func checkPattern(b []byte) bool {
	pattern := []byte("mul(")
	var end int
	var comma int
	for j := 0; j < 4; j++ {
		if j < len(pattern) && b[j] != pattern[j] {
			return false
		}
	}
	for i := 4; i < len(b); i++ {
		if b[i] == []byte(")")[0] && i >= 7 && i <= 11 {
			end = i
		}
		if b[i] == []byte(",")[0] && i >= 5 && i <= 7 {
			comma = i
		}
	}
	if comma == 0 || end == 0 {
		return false
	}
	fmt.Println(b)
	return true

}

func search(b []byte) {
	pattern := []byte("mul(")
	var count int
	for i, r := range b {
		if r != pattern[0] {
			continue
		}
		search := 12 + i
		if search >= len(b) {
			search = len(b)
		}
		ok := checkPattern(b[i:search])
		if ok {
			count++
		}

	}
	fmt.Println(count)
}

func getLines() {
	scanner := getScanner()
	for {
		EoF := !(scanner.Scan())
		if EoF {
			break
		}
		line := scanner.Bytes()
		search(line)

	}
}

func getScanner() *bufio.Scanner {
	file, err := os.OpenFile("input.txt", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return bufio.NewScanner(file)
}

func get_MUL_bytes() []byte {
	return []byte("mul(")
}
