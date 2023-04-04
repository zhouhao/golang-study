package main

import "fmt"

func main() {
	fmt.Println("Hello, playground")
	fmt.Println(findMissingInteger([]int{3, 4, -1, 1}))
	fmt.Println(findMissingInteger([]int{3, 4, -1, 1, 2}))
	fmt.Println(findMissingInteger([]int{3, 1, 2}))
}

func findMissingInteger(arr []int) int {
	if arr == nil && len(arr) == 0 {
		return 1
	}

	for i := 0; i < len(arr); {
		if arr[i]-1 >= 0 && arr[i]-1 < len(arr) && arr[i] != arr[arr[i]-1] {
			swap(arr, i, arr[i]-1)
		} else {
			i++
		}
	}
	for i, v := range arr {
		if i != v-1 {
			return i + 1
		}
	}

	return len(arr) + 1
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func Solution(R string, V []int) []int {
	// Implement your solution here

	// invalid state if R length is not equal to V length
	// in this case, I will return {-1, -1} to indicate this is in a invalid state
	if len(R) < 1 || len(R) != len(V) {
		return []int{-1, -1}
	}

	sum := 0
	for _, v := range V {
		sum += v
	}

	return []int{getMinForTarget(R, V, 0, sum, 'A'), getMinForTarget(R, V, 0, sum, 'B')}
}

func getMinForTarget(R string, V []int, min int, max int, target byte) int {
	// binary search
	for min <= max {
		mid := min + (max-min)/2
		if check(R, V, mid, target) {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}

	return min
}

func check(R string, V []int, balance int, target byte) bool {
	// check if the balance is valid
	for i := 0; i < len(R); i++ {
		if R[i] == target {
			balance += V[i]
		} else {
			balance -= V[i]
		}
		if balance < 0 {
			return false
		}
	}
	return true
}
