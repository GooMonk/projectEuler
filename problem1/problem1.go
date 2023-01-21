package problem1

/*
Problem Statement: If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000.
*/

func sumN(n int64) int64 {
	var res int64
	if n%2 == 0 {
		res = (n / 2) * (n + 1)
	} else {
		res = ((n + 1) / 2) * n
	}
	return res
}

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int64) int64 {
	return ((a * b) / gcd(a, b))
}

// Generic implementation of the problem statement where the function takes 3 arguments a,b and n and returns the sum of all the multiples of a and b less than n.
func MultipleSum(a, b, n int64) int64 {
	var res int64

	if a <= 0 || b <= 0 || n <= 0 {
		return -1
	}
	if a >= n || b >= n {
		return -1
	}
	n--
	if a == b {
		x := n / a
		res = sumN(x)
		res *= a
	} else {
		x := n / a
		y := n / b
		lcmxy := lcm(a, b)
		z := n / lcmxy
		res = (a * sumN(x)) + (b * sumN(y)) - (lcmxy * sumN(z))
	}
	return res
}
