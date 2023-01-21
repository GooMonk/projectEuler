package util

func SumN(n int64) int64 {
	var res int64
	if n%2 == 0 {
		res = (n / 2) * (n + 1)
	} else {
		res = ((n + 1) / 2) * n
	}
	return res
}

func Gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

func IterativeGcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Lcm(a, b int64) int64 {
	return ((a * b) / Gcd(a, b))
}

func Sieve(n int64) []bool {
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

func Pow(a, b int64) int64 {
	var res int64 = 1
	var i int64
	for i = 0; i < b; i++ {
		res *= a
	}
	return res
}

func IsPalindrome(n int64) bool {
	var nrev, ntemp int64 = 0, n

	for ntemp > 0 {
		nrev = 10*nrev + (ntemp % 10)
		ntemp /= 10
	}
	return (nrev == n)

}
