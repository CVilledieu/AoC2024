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
	board        [][]int
	objCount     int
	guardVisited int
	guard        *Guard
}

type Guard struct {
	x   int
	y   int
	dic rune
}

func main() {
	game := createGame()

	for {
		end := game.moveGuard()
		if end {
			break
		}
	}
	fmt.Println(game.guardVisited)
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
		if tile == 0 {
			G.guardVisited++
		}
		G.board[guard.y+dY][guard.x+dX] = 3

		guard.y += dY
		guard.x += dX

	}
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
				nG.guardVisited = 1
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
