package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	x, y      int
	direction rune
	next      *Node
	previous  *Node
}

type Chart struct {
	Board [][]int
	Head  *Node
	Guard *Node
}

func main() {
	board := getChart()

}

func (c *Chart) UpdateGame() {

}

func getMovement(r rune) (int, int)

func getChart() *Chart {
	bArray := getByteArray()
	var Board [][]int
	guard := Node{}
	for y, bArr := range bArray {
		rowInt := make([]int, len(bArr))
		for x, b := range bArr {
			switch rune(b) {
			case rune('.'):
				rowInt[x] = 0
			case rune('#'):
				rowInt[x] = 1
			case rune('^'):
				rowInt[x] = 3
				guard.x, guard.y = x, y
				guard.direction = 'N'
			}
		}
		Board = append(Board, rowInt)
	}
	return &Chart{Board: flipBoard(Board), Guard: &guard}
}

func getByteArray() [][]byte {
	scanner := getScanner()
	var byteArray [][]byte
	for {
		EoF := !(scanner.Scan())
		if EoF {
			break
		}
		byteArray = append(byteArray, scanner.Bytes())

	}
	return byteArray
}

func flipBoard(b [][]int) [][]int {
	fmt.Println(b)
	for i := 0; i <= len(b)/2; i++ {
		b[i], b[len(b)-1] = b[len(b)-1], b[i]
	}
	fmt.Println(b)
	return b
}

func getScanner() *bufio.Scanner {
	file, err := os.OpenFile("input.txt", os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	return bufio.NewScanner(file)
}
