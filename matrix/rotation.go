package matrix

import "fmt"

func Rotation() {
	rotateClockwise90()
	rotateAntiClockwise90()
}

/**
Transpose and the reverse each row
*/
func rotateClockwise90() {
	matrix := [][]int{
		{1, 2},
		{3, 4},
	}
	fmt.Printf("Before Rotation:\n%v\n", matrix)
	transpose(matrix)

	//Reverse each row
	for row := 0; row < len(matrix); row++ {
		end := len(matrix[row]) - 1
		for col := 0; col < end; col++ {
			temp := matrix[row][col]
			matrix[row][col] = matrix[row][end]
			matrix[row][end] = temp
			end--
		}
	}

	fmt.Printf("After Rotation +90:\n%v", matrix)
}

/**
Transpose and the reverse each column
*/
func rotateAntiClockwise90() {
	matrix := [][]int{
		{1, 2},
		{3, 4},
	}
	fmt.Printf("\nBefore Rotation:\n%v\n", matrix)
	transpose(matrix)

	//Reverse each column
	for row := 0; row < len(matrix); row++ {
		end := len(matrix[row]) - 1
		for col := 0; col < end; col++ {
			temp := matrix[col][row]
			matrix[col][row] = matrix[end][row]
			matrix[end][row] = temp
			end--
		}
	}

	fmt.Printf("After Rotation -90:\n%v", matrix)
}
