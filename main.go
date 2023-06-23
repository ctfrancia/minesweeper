package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := 8  // columns
	r := 8  // rows
	nb := 3 // # bombs
	start(c, r, nb)
}

func start(c, r, nb int) {
	emptyBoard := createBoard(c, r)
	bombBoard, locs := setBombs(emptyBoard, c, r, nb)
	boardWithNums := setNumbers(bombBoard, locs)
	fmt.Println("w/ nums", boardWithNums)
	for i := 0; i < len(boardWithNums); i++ {
		fmt.Println(boardWithNums[i])
	}

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

	for _, v := range bombLocs {
		row := v[0]
		col := v[1]

		if row-1 >= 0 {
			if numBoard[row-1][col] != -1 {
				numBoard[row-1][col]++
			}
		}

		if row+1 <= len(numBoard[row])-1 {
			if numBoard[row+1][col] != -1 {
				numBoard[row+1][col]++
			}
		}

		if col-1 >= 0 {
			if numBoard[row][col-1] != -1 {
				numBoard[row][col-1]++
			}
		}

		if col+1 <= len(numBoard[col])-1 {
			if numBoard[row][col+1] != -1 {
				numBoard[row][col+1]++
			}
		}

		if row-1 >= 0 && col-1 >= 0 {
			if numBoard[row-1][col-1] != -1 {
				numBoard[row-1][col-1]++
			}
		}
		if row+1 <= len(numBoard[row])-1 && col+1 <= len(numBoard[col])-1 {
			if numBoard[row+1][col+1] != -1 {
				numBoard[row+1][col+1]++
			}
		}
		if row-1 >= 0 && col+1 <= len(numBoard[col])-1 {
			if numBoard[row-1][col+1] != -1 {
				numBoard[row-1][col+1]++
			}
		}

		if row+1 <= len(numBoard[row])-1 && col-1 >= 0 {
			if numBoard[row+1][col-1] != -1 {
				numBoard[row+1][col-1]++
			}
		}

	}

	return numBoard
}
