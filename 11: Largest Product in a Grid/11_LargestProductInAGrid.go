package main

import (
	"fmt"
)

func main() {
	matrix := make([][]int, 0)

	// Setup.
	matrix = append(matrix, []int{8, 2, 22, 97, 38, 15, 0, 40, 00, 75, 4, 5, 7, 78, 52, 12, 50, 77, 91, 8})
	matrix = append(matrix, []int{49, 49, 99, 40, 17, 81, 18, 57, 60, 87, 17, 40, 98, 43, 69, 48, 04, 56, 62, 00})
	matrix = append(matrix, []int{81, 49, 31, 73, 55, 79, 14, 29, 93, 71, 40, 67, 53, 88, 30, 03, 49, 13, 36, 65})
	matrix = append(matrix, []int{52, 70, 95, 23, 04, 60, 11, 42, 69, 24, 68, 56, 01, 32, 56, 71, 37, 02, 36, 91})
	matrix = append(matrix, []int{22, 31, 16, 71, 51, 67, 63, 89, 41, 92, 36, 54, 22, 40, 40, 28, 66, 33, 13, 80})
	matrix = append(matrix, []int{24, 47, 32, 60, 99, 03, 45, 02, 44, 75, 33, 53, 78, 36, 84, 20, 35, 17, 12, 50})
	matrix = append(matrix, []int{32, 98, 81, 28, 64, 23, 67, 10, 26, 38, 40, 67, 59, 54, 70, 66, 18, 38, 64, 70})
	matrix = append(matrix, []int{67, 26, 20, 68, 02, 62, 12, 20, 95, 63, 94, 39, 63, 8, 40, 91, 66, 49, 94, 21})
	matrix = append(matrix, []int{24, 55, 58, 05, 66, 73, 99, 26, 97, 17, 78, 78, 96, 83, 14, 88, 34, 89, 63, 72})
	matrix = append(matrix, []int{21, 36, 23, 9, 75, 00, 76, 44, 20, 45, 35, 14, 00, 61, 33, 97, 34, 31, 33, 95})
	matrix = append(matrix, []int{78, 17, 53, 28, 22, 75, 31, 67, 15, 94, 03, 80, 04, 62, 16, 14, 9, 53, 56, 92})
	matrix = append(matrix, []int{16, 39, 05, 42, 96, 35, 31, 47, 55, 58, 88, 24, 00, 17, 54, 24, 36, 29, 85, 57})
	matrix = append(matrix, []int{86, 56, 00, 48, 35, 71, 89, 07, 05, 44, 44, 37, 44, 60, 21, 58, 51, 54, 17, 58})
	matrix = append(matrix, []int{19, 80, 81, 68, 05, 94, 47, 69, 28, 73, 92, 13, 86, 52, 17, 77, 04, 89, 55, 40})
	matrix = append(matrix, []int{04, 52, 8, 83, 97, 35, 99, 16, 07, 97, 57, 32, 16, 26, 26, 79, 33, 27, 98, 66})
	matrix = append(matrix, []int{88, 36, 68, 87, 57, 62, 20, 72, 03, 46, 33, 67, 46, 55, 12, 32, 63, 93, 53, 69})
	matrix = append(matrix, []int{04, 42, 16, 73, 38, 25, 39, 11, 24, 94, 72, 18, 8, 46, 29, 32, 40, 62, 76, 36})
	matrix = append(matrix, []int{20, 69, 36, 41, 72, 30, 23, 88, 34, 62, 99, 69, 82, 67, 59, 85, 74, 04, 36, 16})
	matrix = append(matrix, []int{20, 73, 35, 29, 78, 31, 90, 01, 74, 31, 49, 71, 48, 86, 81, 16, 23, 57, 05, 54})
	matrix = append(matrix, []int{01, 70, 54, 71, 83, 51, 54, 69, 16, 92, 33, 48, 61, 43, 52, 01, 89, 19, 67, 48})

	max := 0
	temp := 0

	// Horizontal.
	for j := 0; j < 20; j++ {
		for i := 0; i < 17; i++ {
			if temp = matrix[j][i] * matrix[j][i+1] * matrix[j][i+2] * matrix[j][i+3]; temp > max {
				max = temp
			}
		}
	}

	// Vertical.
	for j := 0; j < 17; j++ {
		for i := 0; i < 20; i++ {
			if temp = matrix[j][i] * matrix[j+1][i] * matrix[j+2][i] * matrix[j+3][i]; temp > max {
				max = temp
			}
		}
	}

	// Left diagonal.
	for j := 0; j < 17; j++ {
		for i := 0; i < 17; i++ {
			if temp = matrix[j][i] * matrix[j+1][i+1] * matrix[j+2][i+2] * matrix[j+3][i+3]; temp > max {
				max = temp
			}
		}
	}

	// Right diagonal.
	for j := 0; j < 17; j++ {
		for i := 0; i < 17; i++ {
			if temp = matrix[j][19-i] * matrix[j+1][18-i] * matrix[j+2][17-i] * matrix[j+3][16-i]; temp > max {
				max = temp
			}
		}
	}

	fmt.Println("Max:", max)
}
