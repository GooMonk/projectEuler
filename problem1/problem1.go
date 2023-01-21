package problem1

import util "projectEuler/util"

/*
Problem Statement: If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000.
*/

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
		res = util.SumN(x)
		res *= a
	} else {
		x := n / a
		y := n / b
		lcmxy := util.Lcm(a, b)
		z := n / lcmxy
		res = (a * util.SumN(x)) + (b * util.SumN(y)) - (lcmxy * util.SumN(z))
	}
	return res
}
