package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
	rows_ := make([][]int, rows)
	for i, _ := range rows_ {
		cols_ := make([]int, cols)
		rows_[i] = cols_
	}
	for i, _ := range rows_ {
		for j, _ := range rows_[i] {
			rows_[i][j] = i * j
		}
	}
	return rows_

}

// dont edit below this line

func test(rows, cols int) {
	fmt.Printf("Creating %v x %v matrix...\n", rows, cols)
	matrix := createMatrix(rows, cols)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test(3, 3)
	test(5, 5)
	test(10, 10)
	test(15, 15)
}
