package problem3

import "testing"

func TestLargestPrimeFactor(test *testing.T) {
	// Unit Testcase: 1 : Project EulerProblem statement
	testResult := LargestPrimeFactor(600851475143)
	test.Logf("The result of the method call is: %d", testResult)
	if testResult != 6857 {
		test.Errorf("The result is not as expected for the input (600851475143). Expected result: 6857")
	}

	testResult2 := LargestPrimeFactor(13195)
	test.Logf("The result of the method call is: %d", testResult2)
	if testResult2 != 29 {
		test.Errorf("The result is not as expected for the input (13195). Expected result: 29")
	}

	testResult3 := LargestPrimeFactor(-1)
	test.Logf("The result of the method call is: %d", testResult3)
	if testResult3 != -1 {
		test.Errorf("The result is not as expected for the input (-1). Expected result: -1")
	}

	testResult4 := LargestPrimeFactor(138)
	test.Logf("The result of the method call is: %d", testResult4)
	if testResult4 != 23 {
		test.Errorf("The result is not as expected for the input (138). Expected result: 23")
	}
}
