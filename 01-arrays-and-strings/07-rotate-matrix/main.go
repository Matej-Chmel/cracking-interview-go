package main

import (
	goi "github.com/Matej-Chmel/go-interview"
)

func solution(matrix [][]int32) [][]int32 {
	layers := len(matrix) / 2

	for layer := 0; layer < layers; layer++ {
		last := len(matrix) - 1 - layer

		for i := layer; i < last; i++ {
			offset := i - layer
			top := matrix[layer][i]

			// top = left
			matrix[layer][i] = matrix[last-offset][layer]

			// left = bottom
			matrix[last-offset][layer] = matrix[last][last-offset]

			// bottom = right
			matrix[last][last-offset] = matrix[i][last]

			// right = top
			matrix[i][last] = top
		}
	}

	return matrix
}

func main() {
	i := goi.NewInterview[[][]int32, [][]int32]()
	i.ReadCases("data/01-07-in.txt", "data/01-07-out.txt")
	i.AddSolution(solution)
	i.Print()
}
