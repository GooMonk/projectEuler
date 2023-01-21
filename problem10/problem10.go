package problem10

import util "projectEuler/util"

/*
Problem Statement: The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

// Generic Implementation which takes a positive natural number n as input and returns the sum of all primes numbers less than the number n.
func PrimeSum(n int64) int64 {
	if n < 0 {
		return -1
	}

	if n < 2 {
		return 0
	}

	// Implementation using Sieve of Eratosthenes
	var res int64

	var notPrime []bool = util.Sieve(n - 1)

	for number, isNotPrime := range notPrime {
		if !isNotPrime {
			res += int64(number)
		}
	}

	return res
}
