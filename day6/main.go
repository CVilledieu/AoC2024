package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
input: a "map"
ex:
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...

^ = a guard and the way she is facing
# = an obstruction

guards movement:
if next tile is blocked turn right
otherwise move forward until she hits an obstruction(#)

Task:
find how many spots the guard will visit

extra notes: area is not a rectangle not a square, (.) dots are unvisited locations
*/

type Game struct {
	board       [][]int
	objCount    int
	vertexs     int
	lastVisited *Vertex
	guard       *Guard
}

type Vertex struct {
	previous *Vertex
	x, y     int
}

type Guard struct {
	x, y, dic int
}

func main() {
	game := createGame()

	for {
		end := game.moveGuard()
		game.CheckSquare()
		if end {
			break
		}
	}
	fmt.Println(game.vertexs)
}

func (G *Game) CheckSquare() int {
	youngest := G.lastVisited
	middleChild := youngest.previous
	oldest := middleChild.previous

	if oldest.x == youngest.x {

	} else {

	}

}

func (G *Game) moveGuard() bool {
	guard := G.guard
	dX, dY := guard.getMovement()
	for {
		if guard.y+dY < 0 || guard.x+dX < 0 || guard.y+dY >= len(G.board) || guard.x+dX >= len(G.board[0]) {
			return true
		}
		tile := G.board[guard.y+dY][guard.x+dX]
		if tile == 1 {
			break
		}
		G.board[guard.y+dY][guard.x+dX] = 3

		guard.y += dY
		guard.x += dX

	}
	v := Vertex{x: guard.x, y: guard.y, previous: G.lastVisited}
	G.lastVisited = &v
	guard.ChangeDic()
	return false
}

func (Gu *Guard) ChangeDic() {
	switch Gu.dic {
	case 'N':
		Gu.dic = 'E'
	case 'E':
		Gu.dic = 'S'
	case 'S':
		Gu.dic = 'W'
	case 'W':
		Gu.dic = 'N'
	}
}

func (Gu *Guard) getMovement() (int, int) {
	var dX, dY int
	switch Gu.dic {
	case 'N':
		dX = 0
		dY = -1
	case 'S':
		dX = 0
		dY = 1
	case 'E':
		dX = 1
		dY = 0
	case 'W':
		dX = -1
		dY = 0
	}
	return dX, dY
}

// 0 = free
// 1 = object
// 2 = guard
// 3 = visited
func createGame() *Game {
	nG := Game{board: [][]int{}}
	sc := getScanner()
	var y int
	for {
		EoF := !(sc.Scan())
		if EoF {
			break
		}
		inputLine := sc.Bytes()
		var newRow []int
		for x, b := range inputLine {
			switch rune(b) {
			case rune('.'):
				newRow = append(newRow, 0)
			case rune('#'):
				newRow = append(newRow, 1)
				nG.objCount++
			case rune('^'):
				newRow = append(newRow, 3)
				nG.lastVisited = &Vertex{x: x, y: y}
				nG.guard = &Guard{x: x, y: y, dic: 'N'}
			}
		}
		nG.board = append(nG.board, newRow)
		y++
	}
	return &nG
}

func getScanner() *bufio.Scanner {
	file, err := os.OpenFile("input.txt", os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	return bufio.NewScanner(file)
}
