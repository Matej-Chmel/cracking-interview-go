package main

import (
	goi "github.com/Matej-Chmel/go-interview"
)

func solution(matrix [][]int32) [][]int32 {
	cols := []int{}
	rows := []int{}

	for rowIndex := 0; rowIndex < len(matrix); rowIndex++ {
		for colIndex := 0; colIndex < len(matrix[rowIndex]); colIndex++ {
			val := matrix[rowIndex][colIndex]

			if val == 0 {
				cols = append(cols, colIndex)
				rows = append(rows, rowIndex)
			}
		}
	}

	for _, c := range cols {
		for rowIndex := 0; rowIndex < len(matrix); rowIndex++ {
			matrix[rowIndex][c] = 0
		}
	}

	for _, r := range rows {
		for colIndex := 0; colIndex < len(matrix[r]); colIndex++ {
			matrix[r][colIndex] = 0
		}
	}

	return matrix
}

func main() {
	i := goi.NewInterview[[][]int32, [][]int32]()
	i.ReadCases("data/01-08-in.txt", "data/01-08-out.txt")
	i.AddSolution(solution)
	i.Print()
}
