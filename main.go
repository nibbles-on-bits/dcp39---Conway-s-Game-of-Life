package main

import "fmt"

func main() {

	fmt.Println("dcp39 - Conway's Game of Life")
	//liveCells := [][2]int{{1, 6}, {3, 2}, {3, 3}, {3, 4}, {5, 2}, {5, 6}, {5, 7}, {6, 1}, {6, 2}, {6, 6}}

	liveCells := [][2]int{{1, 1}, {2, 2}, {5, 7}} // live cells in row,col where upper left is origin

	rowMin := liveCells[0][0]
	rowMax := liveCells[0][0]
	colMin := liveCells[0][1]
	colMax := liveCells[0][1]

	// First determine the board dimensions
	for i := 0; i < len(liveCells); i++ {
		r := liveCells[i][0]
		c := liveCells[i][1]

		if r < rowMin {
			rowMin = r
		}

		if r > rowMax {
			rowMax = r
		}

		if c < colMin {
			colMin = c
		}

		if c > colMax {
			colMax = c
		}
	}

	fmt.Printf("rowMin=%d  rowMax=%d  colMin=%d  colMax=%d\n", rowMin, rowMax, colMin, colMax)
	fmt.Printf("Board is %d rows and %d columns\n", rowMax-rowMin+1, colMax-colMin+1)

	printBoard(liveCells)
}

/*func countNeighbors(liveCells [][2]int, row int, col int) {
	for i := 0; x < len(liveCells); x ++ {
		if liveCells[i]
	}
}*/

// CellsAreNeighbors will test to see if two cell coordinates are neighboring each other
func CellsAreNeighbors(r1 int, c1 int, r2 int, c2 int) bool {
	ret := false
	if r1 == r2 && c1 == c2 {
		return false
	}

	if ((r1 - r2) >= -1) && ((r1 - r2) <= 1) {
		if ((c1 - c2) >= -1) && ((c1 - c2) <= 1) {
			return true
		}
	}

	return ret
}

// given a board represented by boolean values

func printBoard(liveCells [][2]int) {

	rowMin := liveCells[0][0]
	rowMax := liveCells[0][0]
	colMin := liveCells[0][1]
	colMax := liveCells[0][1]

	for i := 0; i < len(liveCells); i++ {
		r := liveCells[i][0]
		c := liveCells[i][1]

		if r < rowMin {
			rowMin = r
		}

		if r > rowMax {
			rowMax = r
		}

		if c < colMin {
			colMin = c
		}

		if c > colMax {
			colMax = c
		}
	}

	width := colMax - colMin + 1
	height := rowMax - rowMin + 1

	board := make([][]bool, height)
	for y := 0; y < height; y++ {
		board[y] = make([]bool, width)
	}

	for i := 0; i < len(liveCells); i++ {
		r := liveCells[i][0]
		c := liveCells[i][1]

		board[r-rowMin][c-colMin] = true
	}

	fmt.Printf("height=%d  width=%d\n", height, width)

	fmt.Println(board)

	// now we have a boolean representation of the board
	for row := rowMin; row <= rowMax; row++ {
		for col := colMin; col <= colMax; col++ {
			if board[row-rowMin][col-colMin] {
				fmt.Print("*")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

}
