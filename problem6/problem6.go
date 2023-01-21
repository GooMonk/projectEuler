package problem6

import util "projectEuler/util"

/*
	Problem StatementThe sum of the squares of the first ten natural numbers is 385,
	The square of the sum of the first ten natural numbers is 3025,
	Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 2640.
	Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

// Generic Implementation which tries to find the difference between the sum of squares of the first n natural numbers and the squares of the sum.
func SumSquareDifference(n int64) int64 {
	if n < 0 {
		return -1
	}
	var res int64 = util.Pow(util.SumN(n), 2)

	for i := int64(1); i <= n; i++ {
		res -= util.Pow(i, 2)
	}
	return res
}
