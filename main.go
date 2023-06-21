package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := 8  // columns
	r := 8  // rows
	nb := 1 // # bombs
	start(c, r, nb)
}

func start(c, r, nb int) {
	emptyBoard := createBoard(c, r)
	// fmt.Printf("len=%d cap=%d %v\n", len(emptyBoard), cap(emptyBoard), emptyBoard)
	bombBoard, locs := setBombs(emptyBoard, c, r, nb)
	boardWithNums := setNumbers(bombBoard, locs)
	fmt.Println("w/ nums", boardWithNums)

}

func createBoard(cols, rows int) [][]int {
	matrix := [][]int{}

	for r := 0; r < rows; r++ {
		rHolder := []int{}
		for c := 0; c < cols; c++ {
			rHolder = append(rHolder, 0)
		}
		matrix = append(matrix, rHolder)
	}

	return matrix
}

func setBombs(eb [][]int, c, r, nb int) ([][]int, [][]int) {
	bwb := eb
	bombLocations := [][]int{}
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for b := 0; b < nb; b++ {
		rowCoord := rd.Intn(r)
		colCoord := rd.Intn(c)
		bwb[rowCoord][colCoord] = -1
		bombLocations = append(bombLocations, []int{rowCoord, colCoord})
	}
	return bwb, bombLocations
}

func setNumbers(bwb, bombLocs [][]int) [][]int {
	numBoard := bwb

	for i, v := range bombLocs {
		fmt.Println(i, v)
		row := v[0]
		col := v[1]
		// rowLength := len(numBoard[row])
		// colLength := len(numBoard[col])
		// fmt.Println("row-1", row-1, row-1 >= 0)
		// mt.Println("col-1", col-1)
		numBoard[row-1][col]++
		numBoard[row+1][col]++
		numBoard[row][col-1]++
		numBoard[row][col+1]++
		numBoard[row+1][col-1]++
		numBoard[row-1][col+1]++
		numBoard[row-1][col-1]++
		numBoard[row+1][col+1]++

		/*
			if row+1 >= rowLength {
				if numBoard[row+1][col] != -1 {
					numBoard[row+1][col]++
				}
			}

			if row-1 >= 0 {
				if numBoard[row-1][col] != -1 {
					numBoard[row-1][col]++
				}
			}

			if col-1 >= colLength {
				if numBoard[row][col+1] != -1 {
					numBoard[row][col+1]++
				}
			}
			if col-1 >= 0 {
				if numBoard[row][col-1] != -1 {
					numBoard[row][col-1]++
				}
			}
		*/
		/*
			if numBoard[row+1][col-1] != -1 {
				numBoard[row+1][col-1]++
			}
			if numBoard[row-1][col+1] != -1 {
				numBoard[row-1][col+1]++
			}
		*/
		/*
			if row-1 >= 0 {
				if col-1 >= 0 {

				}
			}
		*/

		/*
			if row-1 <= 0 && col-1 <= 0 {
				if numBoard[row-1][col] != -1 {
					numBoard[row-1][col]++
				}
			}

			if row+1 > len(numBoard[row]) && col+1 > len(numBoard[col]) {
				if numBoard[row+1][col] != -1 {
					numBoard[row+1][col]++
				}
			}

			if col-1 <= 0 {
				if numBoard[row][col-1] != -1 {
					numBoard[row][col-1]++
				}
			}

			if col+1 <= len(numBoard[col]) {
				if numBoard[row][col+1] != -1 {
					numBoard[row][col+1]++
				}
			}
		*/
	}

	return numBoard
}

func tileInBounds(len, place int) bool {
	return true
}