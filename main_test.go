package main

import (
	"fmt"
	"testing"
)

func TestCellsAreNeighbors(t *testing.T) {
	testShouldPass := [][]int{
		{0, 0, 0, 1},
		{0, 1, 0, 0},
		{5, 5, 4, 4},
		{5, 5, 4, 5},
		{5, 5, 4, 6},
		{5, 5, 5, 4},
		{5, 5, 5, 6},
		{5, 5, 6, 4},
		{5, 5, 6, 5},
		{5, 5, 6, 6},
		{0, 0, -1, -1},
	}

	for _, tc := range testShouldPass {
		got := CellsAreNeighbors(tc[0], tc[1], tc[2], tc[3])
		want := true
		if got != want {
			t.Errorf("CellsAreNeighbors(%d,%d,%d,%d) = >>%v<<; want >>%v<<", tc[0], tc[1], tc[2], tc[3], got, want)
		}
	}

	testShouldFail := [][]int{
		{0, 0, 0, 0},
		{0, 1, 0, 3},
		{5, 5, 4, 41},
		{5, 5, 3, 5},
		{5, 5, 4, 8},
		{1, 1, 5, 4},
		{5, 7, 5, 2},
		{5, 100, 6, 4},
		{5, 5, 65, 5},
		{5, 5, 6, 3},
	}

	for _, tc := range testShouldFail {
		got := CellsAreNeighbors(tc[0], tc[1], tc[2], tc[3])
		want := false
		if got != want {
			t.Errorf("CellsAreNeighbors(%d,%d,%d,%d) = >>%v<<; want >>%v<<", tc[0], tc[1], tc[2], tc[3], got, want)
		}
	}
}

func TestCountNeighbors(t *testing.T) {

	liveCells := [][2]int{{1, 1}, {2, 2}, {5, 7}}

	//
	got := CountNeighbors(liveCells, 0, 0)
	want := 1
	if got != want {
		t.Errorf("CountNeighbors(%v,%d,%d)  got:%v  want %v\n", liveCells, 0, 0, got, want)
	}

	got = CountNeighbors(liveCells, 1, 2)
	want = 2
	if got != want {
		t.Errorf("CountNeighbors(%v,%d,%d)  got:%v  want %v\n", liveCells, 0, 0, got, want)
	}

}

func TestIsCellLive(t *testing.T) {

	liveCells := [][2]int{{1, 1}, {2, 2}, {5, 7}}

	got := IsCellLive(liveCells, 0, 0)
	want := false
	if got != want {
		t.Errorf("CountNeighbors(%v,%d,%d)  got:%v  want %v\n", liveCells, 0, 0, got, want)
	}

	got = IsCellLive(liveCells, 2, 2)
	want = true
	if got != want {
		t.Errorf("CountNeighbors(%v,%d,%d)  got:%v  want %v\n", liveCells, 2, 2, got, want)
	}

}

func TestGetNextBoard(t *testing.T) {
	liveCells := [][2]int{{1, 0}, {1, 1}, {1, 2}}

	got := GetNextBoard(liveCells)
	want := [][2]int{{0, 1}, {1, 1}, {2, 1}}

	//if got != want {
	t.Errorf("GetNextBoard(%v)  got:%v  want %v\n", liveCells, got, want)
	//}

	////// row of 10
	liveCells = [][2]int{{5, 0}, {5, 1}, {5, 2}, {5, 3}, {5, 4}, {5, 5}, {5, 6}, {5, 7}, {5, 8}, {5, 9}, {5, 10}}
	got = GetNextBoard(liveCells)
	want = [][2]int{}

	//if got != want {
	t.Errorf("GetNextBoard(%v)  got:%v  want %v\n", liveCells, got, want)
	PrintBoard(got)
}

func TestPrintBoard(t *testing.T) {
	//rowOf10 := [][2]int{{5, 0}, {5, 1}, {5, 2}, {5, 3}, {5, 4}, {5, 5}, {5, 6}, {5, 7}, {5, 8}, {5, 9}}

	glider := [][2]int{{0, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2}}

	for x := 0; x < 50; x++ {
		PrintBoard(glider)
		fmt.Println(glider)
		glider = GetNextBoard(glider)
	}

	/*
		LiveCells := [][2]int{{4, 1}, {4, 2}, {4, 3}, {4, 4}, {4, 5}, {4, 6}, {4, 7}, {4, 8}, {4, 9}, {5, 1}, {5, 2}, {5, 3}, {5, 4}, {5, 5}, {5, 6}, {5, 7}, {5, 8}, {5, 9}, {6, 1}, {6, 2}, {6, 3}, {6, 4}, {6, 5}, {6, 6}, {6, 7}, {6, 8}, {6, 9}}
		PrintBoard(LiveCells)

		fmt.Println()

		NextBoard1 := GetNextBoard(LiveCells)
		PrintBoard(NextBoard1)

		fmt.Println()

		NextBoard2 := GetNextBoard(NextBoard1)
		PrintBoard(NextBoard2)*/
}

func TestPrintCustomBoard(t *testing.T) {

	/*glider := [][2]int{{0, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2}}

	for x := 0; x < 300; x++ {
		PrintCustomBoard(glider, -5, 15, -5, 100)
		glider = GetNextBoard(glider)
		fmt.Println()
	}*/

	rowOf10 := [][2]int{{5, 0}, {5, 1}, {5, 2}, {5, 3}, {5, 4}, {5, 5}, {5, 6}, {5, 7}, {5, 8}, {5, 9}}
	for x := 0; x < 100; x++ {
		PrintCustomBoard(rowOf10, 0, 10, -5, 15)
		rowOf10 = GetNextBoard(rowOf10)
		fmt.Println()
	}

}
