package main

func main() {

}

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here

	var left int32 = 0
	var right int32 = 0
	var length = len(arr)

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			left += arr[i][j]
			right += arr[i][length-j-1]
		}
	}

	return left - right

}
