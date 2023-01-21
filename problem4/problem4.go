package problem4

/*
Problem Statement: A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

func isPalindrome(n int64) bool {
	var nrev, ntemp int64 = 0, n

	for ntemp > 0 {
		nrev = 10*nrev + (ntemp % 10)
		ntemp /= 10
	}
	return (nrev == n)

}

// Generic Implementation which finds the largest palindrome product made by two n-digit numbers.
func LargestPalindromeProduct(n int8) int64 {
	if n <= 0 || n > 9 {
		return -1
	}
	var low, high, i, j, maxpro int64 = 1, 9, 0, 0, -1
	for k := int8(0); k < n-1; k++ {
		low *= 10
		high = 10*high + 9
	}

	for i = high; i >= low; i-- {
		for j = i; j >= low; j-- {
			product := i * j
			if product > maxpro {
				if isPalindrome(product) {
					maxpro = product
				}
			} else {
				break
			}
		}
	}
	return maxpro
}
