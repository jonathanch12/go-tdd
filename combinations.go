package main

import "fmt"

func combine(n int, k int) ([][]int, error) {
	var result [][]int

	if n <= 0 {
		return nil, fmt.Errorf("invalid input: n must be greater than 0, got %d", n)
	}
	if k < 0 {
		return nil, fmt.Errorf("invalid input: k must not be negative, got %d", k)
	}
	if k > n {
		return nil, fmt.Errorf("invalid input: k (%d) must not be greater than n (%d)", k, n)
	}

	if k == 0 {
		return result, nil
	}

	var current []int
	backtrack(1, n, k, current, &result)
	return result, nil
}

func backtrack(start int, n int, k int, current []int, result *[][]int) {
	if len(current) == k {
		combination := make([]int, k)
		copy(combination, current)
		*result = append(*result, combination)
		return
	}

	for i := start; i <= n-(k-len(current))+1; i++ {
		current = append(current, i)
		backtrack(i+1, n, k, current, result)
		current = current[:len(current)-1]
	}
}
