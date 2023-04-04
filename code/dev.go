package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello, playground")

	name := []string{"a", "b", "c", "d"}
	prices := [][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	fmt.Println(getTopPrice(name, prices))
}

func getTopPrice(name []string, prices [][]float64) []string {
	if len(name) == 0 || len(prices) == 0 {
		return []string{}
	}

	result := make([]string, 3)
	sum := make([]float64, len(prices[0]))

	for i := 0; i < len(prices); i++ {
		for j := 0; j < len(prices[0]); j++ {
			sum[j] += prices[i][j]
		}
	}
	fmt.Println(sum)

	sumCopy := clone(sum)
	sort.Float64s(sumCopy)

	for i := 0; i < 3; i++ {
		v := sumCopy[i]
		for j := 0; j < len(sum); j++ {
			if sum[j] == v {
				result[i] = name[j]
				sum[j] = -1
				break
			}
		}
	}

	return result
}

func clone(prices []float64) []float64 {
	result := make([]float64, len(prices))
	for i := 0; i < len(prices); i++ {
		result[i] = prices[i]
	}
	return result
}
