package util

import (
	"fmt"
	"testing"
)

// Comparing the Iterative and Recursive versions of the functions
func BenchmarkGcd(b *testing.B) {
	input := [][]int64{{10031123, 9023281}, {241242, 43242144}, {43, 31}}
	for i := 0; i < len(input); i++ {
		b.Run(fmt.Sprintf("Iterative: For Input: %d , %d : Out: %d", input[i][0], input[i][1], IterativeGcd(input[i][0], input[i][1])), func(b *testing.B) { subBenchmarkIterativeGCD(input[i][0], input[i][1], b) })
		b.Run(fmt.Sprintf("Recursive: For Input: %d , %d: Out: %d", input[i][0], input[i][1], Gcd(input[i][0], input[i][1])), func(b *testing.B) { subBenchmarkGCD(input[i][0], input[i][1], b) })
	}
}

func subBenchmarkIterativeGCD(a, b int64, bench *testing.B) {
	for i := 0; i <= bench.N; i++ {
		IterativeGcd(a, b)
	}
}

func subBenchmarkGCD(a, b int64, bench *testing.B) {
	for i := 0; i <= bench.N; i++ {
		Gcd(a, b)
	}
}

func BenchmarkPow(b *testing.B) {
	input := [][]int64{{2, 11}, {47, 3}, {631, 8}, {9, 2}}
	for i := 0; i < len(input); i++ {
		b.Run(fmt.Sprintf("Pow: For Input: %d , %d : Out: %d", input[i][0], input[i][1], Pow(input[i][0], input[i][1])), func(b *testing.B) { subBenchmarkPow(input[i][0], input[i][1], b) })
		b.Run(fmt.Sprintf("RecPow: For Input: %d , %d : Out: %d", input[i][0], input[i][1], RecPow(input[i][0], input[i][1])), func(b *testing.B) { subBenchmarkRecPow(input[i][0], input[i][1], b) })
		b.Run(fmt.Sprintf("LinPow: For Input: %d , %d: Out: %d", input[i][0], input[i][1], LinPow(input[i][0], input[i][1])), func(b *testing.B) { subBenchmarkLinPow(input[i][0], input[i][1], b) })
	}
}

func subBenchmarkPow(a, b int64, bench *testing.B) {
	for i := 0; i <= bench.N; i++ {
		Pow(a, b)
	}
}

func subBenchmarkLinPow(a, b int64, bench *testing.B) {
	for i := 0; i <= bench.N; i++ {
		LinPow(a, b)
	}
}

func subBenchmarkRecPow(a, b int64, bench *testing.B) {
	for i := 0; i <= bench.N; i++ {
		RecPow(a, b)
	}
}
