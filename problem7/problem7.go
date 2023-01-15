package problem7

import "math"

/*
Problem Statement: By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
*/

func sieve(n int64) []bool {
	var i, j int64
	var notPrime []bool = make([]bool, n+1)

	notPrime[0] = true
	notPrime[1] = true

	for i = 2; i*i <= n; i++ {
		if !notPrime[i] {
			for j = i * i; j <= n; j += i {
				notPrime[j] = true
			}
		}
	}
	return notPrime
}

// Generic Implementation where for a the given number n, the function returns the nth prime. If not found returns -1
func NthPrime(n int64) int64 {
	if n <= 0 {
		return -1
	}
	//handling simple cases
	switch n {
	case 1:
		return 2
	case 2:
		return 3
	case 3:
		return 5
	case 4:
		return 7
	case 5:
		return 11
	}
	var i int64
	var upperLimit int64 = 15

	if n > 5 {
		approxLimit := int64(math.Ceil(float64(n) * (math.Log(float64(n)) + math.Log(math.Log(float64(n)))))) // An Approximation limit till where the Sieve needs to be run
		if approxLimit > upperLimit {
			upperLimit = approxLimit
		}
	}
	result := sieve(upperLimit)

	i = 1
	for number, isNotPrime := range result {
		if !isNotPrime {
			if i < n {
				i++
			} else {
				return int64(number)
			}
		}
	}
	return -1
}
