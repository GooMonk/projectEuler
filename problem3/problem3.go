package problem3

/*
Problem Statement: The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

// Generic Implementation function which returns the largest prime factor for the provided number n.
func LargestPrimeFactor(n int64) int64 {
	var res int64 = n

	if n <= 3 {
		return res
	}

	var ntemp int64 = n
	var i int64
	for ntemp%2 == 0 {
		ntemp /= 2
		res = 2
	}
	for i = 3; i*i <= n; i += 2 {
		if ntemp%i == 0 {
			res = i
			for ntemp%i == 0 {
				ntemp /= i
			}
		}
	}

	if ntemp > 1 {
		return ntemp
	}
	return res
}
