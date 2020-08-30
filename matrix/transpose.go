package matrix

import "fmt"

func TransposeInPlace() {
	matrix := [][]int{
		{1, 2},
		{3, 4},
	}
	fmt.Printf("Before Transpose:\n%v", matrix)
	transpose(matrix)
	fmt.Printf("\nAfter Transpose:\n%v", matrix)
}

func transpose(matrix [][]int) {
	for i, row := range matrix {
		for j, _ := range row {
			if j >= i+1 {
				temp := matrix[i][j]
				matrix[i][j] = matrix[j][i]
				matrix[j][i] = temp
			}
		}
	}
}
