package main

func transpose(matrix [][]int) [][]int {
	var out [][]int

	for i := 0; i < len(matrix[0]); i++ {
		var arr []int
		for j := 0; j < len(matrix); j++ {
			arr = append(arr, matrix[j][i])
		}
		out = append(out, arr)
	}
	return out
}
