package main

import "fmt"

func countWaysRecursive(n int, memo []int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if memo[n] != -1 {
		return memo[n]
	}

	memo[n] = countWaysRecursive(n-2, memo) + countWaysRecursive(n-3, memo)
	return memo[n]
}

func countWaysIterative(n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		dp[i] = 0
		if i-2 >= 0 {
			dp[i] += dp[i-2]
		}
		if i-3 >= 0 {
			dp[i] += dp[i-3]
		}
	}

	return dp[n]
}

func main() {
	var n int
	fmt.Print("Masukkan jarak (km): ")
	fmt.Scan(&n)

	if n < 100 {
		fmt.Println("Jarak harus minimal 100 km.")
		return
	}

	// Rekursif dengan Memoization
	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}
	resultRecursive := countWaysRecursive(n, memo)

	// Iteratif
	resultIterative := countWaysIterative(n)

	// Menampilkan hasil
	fmt.Printf("Jumlah cara mencapai jarak %d km:\n", n)
	fmt.Printf("- Rekursif: %d\n", resultRecursive)
	fmt.Printf("- Iteratif: %d\n", resultIterative)
}