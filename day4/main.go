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

(-1,1)  (0,1)  (1,1)
(-1,0) 	(x,y)  (1,0)
(-1,-1) (0,-1) (1,-1)
*/

const (
	xmas  = "MAS"
	UP    = 1
	DOWN  = -1
	LEFT  = -1
	RIGHT = 1
	NONE  = 0
)

type Graph struct {
	Grid  [][]rune
	Count int
}

func main() {
	//Sets graph.grid from input.txt and sets graph.X/graph.Y to 0
	graph := createGraph()
	graph.seek()
	fmt.Println(graph.Count)
}

func (g *Graph) seek() {
	for y, r := range g.Grid {
		for x, n := range r {
			if n != 'A' {
				continue
			}
			g.search(x, y)

		}

	}
}

func (G *Graph) NE_to_SW(x,y int) bool{
	mas := []rune(xmas)
	if G.Grid
}
func (G *Graph) NW_to_SE(x,y int){

}

func (G *Graph) search(x, y int) {
	dX := []int{1, 1, -1, -1}
	dY := []int{1, -1, -1, 1}
	mas := []rune(xmas)
	for i := range dX {
		for j, n := range mas {
			newX := x + (dX[i] * j)
			newY := y + (dY[i] * j)
			if newY < 0 || newX < 0 || newY >= len(G.Grid) || newX >= len(G.Grid[y]) {
				break
			}
			if n != G.Grid[newY][newX] {
				break
			}
			if j == 3 {
				G.Count++
			}
		}
	}
}

/*
func search(grid [][]rune) {
	var dX int
	var dY int
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

		}
	}
	fmt.Println(dX)
}
*/

func createGraph() *Graph {
	return &Graph{Grid: loadGrid(), Count: 0}
}

// theGrid is 140 by 140
func loadGrid() [][]rune {
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
