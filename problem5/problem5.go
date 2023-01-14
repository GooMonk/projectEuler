package problem5

/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

func pow(a, b int64) int64 {
	var res int64 = 1
	var i int64
	for i = 0; i < b; i++ {
		res *= a
	}
	return res
}

// Generic Implementation which returns smallest number that can be divided by each of the numbers from 1 to n.
func SmallestMultiple(n int64) int64 {
	//Handling side and simple cases
	if n < 0 {
		return -1
	}

	if n < 4 {
		return n
	}

	//Intialization
	var res, i, j int64
	res = 1
	var factor []int64 = make([]int64, n+1)

	factor[0] = 0
	factor[1] = 1

	// Sieve algorithm to find the populate a factor
	for i = 2; i*i <= n; i++ {
		if factor[i] == 0 {
			for j = i * i; j <= n; j += i {
				factor[j] = i
			}
		}
	}
	// Looping through to find all the factors of the given number and find the maximum times an individual number is needed
	var numberOfFact []int64 = make([]int64, n+1)
	for i = 2; i <= n; i++ {
		if factor[i] == 0 {
			numberOfFact[i]++
		} else {
			factorsOfI := make(map[int64]int64)
			for j = i; j > 1; {
				if factor[j] == 0 {
					factorsOfI[j]++
					j = 1
				} else {
					factorsOfI[factor[j]]++
					j /= factor[j]
				}
			}
			for ind, ele := range factorsOfI {
				if numberOfFact[ind] < ele {
					numberOfFact[ind] = ele
				}
			}
		}
	}
	for number, exponent := range numberOfFact {
		if number > 0 && exponent > 0 {
			res *= (pow(int64(number), exponent))
		}
	}
	return res
}
