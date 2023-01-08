package problem1

import (
	"testing"
)

func TestMultipleSum(test *testing.T) {
	// Unit Testcase: 1 : Project Euler
	testResult := MultipleSum(3, 5, 1000)
	test.Logf("The result of the method call is: %d", testResult)
	if testResult != 233168 {
		test.Errorf("The result is not as expeccted for the input (3, 5, 1000) with 233168")
	}

	testResult2 := MultipleSum(5, 5, 10)
	test.Logf("The result of the method call is: %d", testResult2)
	if testResult2 != 5 {
		test.Errorf("The result is not as expeccted for the input (5, 5, 10) with 5")
	}

	testResult3 := MultipleSum(11, 5, 10)
	test.Logf("The result of the method call is: %d", testResult3)
	if testResult3 != -1 {
		test.Errorf("The result is not as expected for the input (11, 5, 10) with -1")
	}

	testResult4 := MultipleSum(-3, 5, 10)
	test.Logf("The result of the method call is: %d", testResult4)
	if testResult4 != -1 {
		test.Errorf("The result is not as expeccted for the input (-5, 5, 10) with -1")
	}
}
