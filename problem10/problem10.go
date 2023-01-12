package problem10

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
	var res, i, j int64

	var notPrime []bool = make([]bool, n)

	notPrime[0] = true
	notPrime[1] = true

	for i = 2; i*i < n; i++ {
		if !notPrime[i] {
			for j = i * i; j < n; j += i {
				notPrime[j] = true
			}
		}
	}

	for number, isNotPrime := range notPrime {
		if !isNotPrime {
			res += int64(number)
		}
	}

	return res
}
