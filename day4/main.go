package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
Input: word search
Find all: "XMAS"s
Word can be horizontal, vertical, diagonal, backwards, or overlapping otherwords

Goal answer: number of "XMAS"s

(1,0)||(0,0)
============
(1,1)||(0,1)
*/

const (
	xmas = "XMAS"
)

func main() {
	grid := getGrid()
	search(grid)
}

func search(grid [][]rune) {
	var dX int
	length := len(xmas)
	gridlength := len(grid)
	//search diagonal
	for y := 0; y < gridlength; y++ {
		for x := 0; x < gridlength; x++ {
			if grid[y][x] != 'X' {
				continue
			}
			if x-length < 0 || x+length > gridlength {
				continue
			}
			if y-length < 0 || y+length > gridlength {
				continue
			}
			dX++
		}
	}
	fmt.Println(dX)
}

// theGrid is 140 by 140
func getGrid() [][]rune {
	scanner := getScanner()
	var line []rune
	var grid [][]rune
	for {
		EoF := !(scanner.Scan())
		if EoF {
			break
		}
		line = getLine(scanner.Text())
		grid = append(grid, line)

	}
	return grid
}

func getLine(s string) []rune {
	var line []rune
	for _, r := range s {
		line = append(line, r)
	}
	return line
}

func getScanner() *bufio.Scanner {
	file, err := os.OpenFile("input.txt", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return bufio.NewScanner(file)
}
