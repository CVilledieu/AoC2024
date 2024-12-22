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
type scramble struct {
	pairs [][]int
	sum   int
	do    bool
}

func main() {
	scanner := getScanner()
	scram := scramble{sum: 0, pairs: [][]int{}, do: true}
	var total int
	for {
		EoF := !(scanner.Scan())
		if EoF {
			break
		}
		line := scanner.Bytes()
		total += scram.parse(line)

	}
	fmt.Println(scram.pairs)
	fmt.Println(total)

}

// Parsing to find do() / don't()
// Only after a byte that matches "d" will this func get called
func (s *scramble) parseDont(b []byte) {
	doByte := []byte("do()")
	dontByte := []byte("don't()")
	var index int
	if b[index] != doByte[index] || b[index+1] != doByte[index+1] {
		return
	}
	index++
	if b[index] == doByte[index] && b[index+1] == doByte[index+1] {
		s.do = true
		return
	} else if b[index] == dontByte[index] && b[index+1] == dontByte[index+1] {
		index++
		if b[index] != dontByte[index] {
			return
		}
		index++
		if b[index] != dontByte[index] {
			return
		}
		index++
		if b[index] == dontByte[index] {
			s.do = false
		}
	}

}

// mul(X,X)
// mul(XXX,XXX)
// index 4 has to be somewhere between 48 and 57
// index 5 through 10 could also be 48 through 57
// one of 5,6, or 7 have to be a 44
// one of 7 through 11 have to be a 41
func (s *scramble) checkPattern(b []byte) int {
	pattern := []byte("mul(")
	var end int
	var comma int
	for j := 0; j < 4; j++ {
		if j < len(pattern) && b[j] != pattern[j] {
			return 0
		}
	}
	for i := 4; i < len(b); i++ {
		if b[i] == []byte(")")[0] && i >= 7 && i <= 11 {
			end = i
		}
		if b[i] == []byte(",")[0] && i >= 5 && i <= 7 {
			comma = i
		}
		if comma > 0 && end > 0 {
			break
		}
	}
	if comma == 0 || end == 0 || comma > end {
		return 0
	}

	left := b[4:comma]
	right := b[comma+1 : end]
	l := bToI(left)
	r := bToI(right)
	s.pairs = append(s.pairs, []int{l, r})
	sum := l * r
	s.sum += sum
	return sum

}

// Called by main
// Returns total sum of products found in checkPattern
func (s *scramble) parse(b []byte) int {
	mulByte := []byte("mul(")
	dontByte := []byte("don't()")
	var sum int
	for i, r := range b {
		if r != dontByte[0] && s.do == false {
			continue
		}
		search := len(dontByte) + i
		s.parseDont(b[i:search])
		if r != mulByte[0] || s.do == false {
			continue
		}
		//12 = len(mulByte) + up to 3 digits for left and right + , + )
		search = 12 + i
		if search >= len(b) {
			search = len(b)
		}
		prod := s.checkPattern(b[i:search])
		sum += prod

	}
	return sum
}

func getScanner() *bufio.Scanner {
	file, err := os.OpenFile("input.txt", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return bufio.NewScanner(file)
}

func bToI(b []byte) int {
	init := make([]int, len(b))
	var number int
	for i, n := range b {
		if n < 48 || n > 57 {
			return 0
		}
		init[i] = int(n) - 48
		if number > 0 {
			number *= 10
		}
		number += int(n) - 48
	}
	return number
}
