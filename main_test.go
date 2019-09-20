package main

import "testing"

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
