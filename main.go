package main

import "fmt"

type gameBoard struct {
	xMin      int
	xMax      int
	yMin      int
	yMax      int
	liveCells [][2]int // stores the x,y coordinates of each live cell
}

func GetBoardFromLiveCells(liveCells [][2]int) gameBoard {
	board := gameBoard{}
	board.xMin = liveCells[0][0]
	board.xMax = liveCells[0][0]
	board.yMin = liveCells[0][1]
	board.yMax = liveCells[0][1]
	board.liveCells = liveCells

	// First determine the board dimensions
	for i := 0; i < len(liveCells); i++ {
		x := liveCells[i][0]
		y := liveCells[i][1]

		if x < board.xMin {
			board.xMin = x
		}

		if x > board.xMax {
			board.xMax = x
		}

		if y < board.yMin {
			board.yMin = y
		}

		if y > board.yMax {
			board.yMax = y
		}
	}

	width := board.xMax - board.xMin + 1
	height := board.yMax - board.yMin + 1

	// liveMatrix is a 2 dimensional slice of bool's representing the board in it's currentState
	liveMatrix := make([][]bool, height)
	for y := 0; y < height; y++ {
		liveMatrix[y] = make([]bool, width)
	}

	for i := 0; i < len(liveCells); i++ {
		x := liveCells[i][0]
		y := liveCells[i][1]

		liveMatrix[x-board.xMin][y-board.yMin] = true
	}

	fmt.Printf("height=%d  width=%d\n", height, width)
	fmt.Printf("xmin=%d xmax=%d ymin=%d ymax=%d\n", board.xMin, board.xMax, board.yMin, board.yMax)
	fmt.Println(board)

	// now we have a boolean representation of the board
	for row := yMin; row <= yMax; row++ {
		for col := xMin; col <= xMax; col++ {
			if board[col-xMin][row-yMin] {
				fmt.Print("*")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	return ret
}

func (gameBoard) print() {
	fmt.Println("Welcome to gameBoard.print()")
}

func main() {

	fmt.Println("dcp39 - Conway's Game of Life")
	//liveCells := [][2]int{{1, 6}, {3, 2}, {3, 3}, {3, 4}, {5, 2}, {5, 6}, {5, 7}, {6, 1}, {6, 2}, {6, 6}}
	liveCells := [][2]int{{1, 1}, {2, 2}, {5, 5}}

	GetBoardFromLiveCells(liveCells)

	fmt.Printf("len(liveCells)=%d\n", len(liveCells))

	fmt.Println()
	printBoard(liveCells)
}

func getNextBoard(thisBoard [][]bool) {
	// First thing we want to do is determine if the next board will be the same dimensions
	// or will increase dimensions.

}

// given a board represented by boolean values

func printBoard(liveCells [][2]int) {
	xMin := liveCells[0][0]
	xMax := liveCells[0][0]
	yMin := liveCells[0][1]
	yMax := liveCells[0][1]

	for i := 0; i < len(liveCells); i++ {
		x := liveCells[i][0]
		y := liveCells[i][1]

		if x < xMin {
			xMin = x
		}

		if x > xMax {
			xMax = x
		}

		if y < yMin {
			yMin = y
		}

		if y > yMax {
			yMax = y
		}
	}

	width := xMax - xMin + 1
	height := yMax - yMin + 1

	board := make([][]bool, height)
	for y := 0; y < height; y++ {
		board[y] = make([]bool, width)
	}

	for i := 0; i < len(liveCells); i++ {
		x := liveCells[i][0]
		y := liveCells[i][1]

		board[x-xMin][y-yMin] = true
	}

	fmt.Printf("height=%d  width=%d\n", height, width)
	fmt.Printf("xmin=%d xmax=%d ymin=%d ymax=%d\n", xMin, xMax, yMin, yMax)
	fmt.Println(board)

	// now we have a boolean representation of the board
	for row := yMin; row <= yMax; row++ {
		for col := xMin; col <= xMax; col++ {
			if board[col-xMin][row-yMin] {
				fmt.Print("*")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

}
