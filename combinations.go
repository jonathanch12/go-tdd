package main

func combine(n int, k int) [][]int {
	var result [][]int

	if n <= 0 || k <= 0 || k > n {
		return result
	}

	var current []int
	backtrack(1, n, k, current, &result)
	return result
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
