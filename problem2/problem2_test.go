package problem2

import (
	"testing"
)

func TestEvenFibSum(test *testing.T) {
	testResult1 := EvenFibSum(1, 2, 10)
	test.Logf("The result of the method call is: %d", testResult1)
	if testResult1 != 10 {
		test.Errorf("The result is not as expected for the input (1, 2, 10). Expected result: 10")
	}

	testResult2 := EvenFibSum(10, 2, 10)
	test.Logf("The result of the method call is: %d", testResult2)
	if testResult2 != -1 {
		test.Errorf("The result is not as expected for the input (10, 2, 10). Expected result: -1")
	}

	testResult3 := EvenFibSum(-1, 2, 10)
	test.Logf("The result of the method call is: %d", testResult3)
	if testResult3 != -1 {
		test.Errorf("The result is not as expected for the input (-1, 2, 10). Expected result: -1")
	}

	testResult4 := EvenFibSum(2, 1, 35)
	test.Logf("The result of the method call is: %d", testResult4)
	if testResult4 != 44 {
		test.Errorf("The result is not as expected for the input (2, 1, 10). Expected result: 10")
	}

	testResult5 := EvenFibSum(1, 2, 4000000)
	test.Logf("The result of the method call is: %d", testResult5)
	if testResult5 != 4613732 {
		test.Errorf("The result is not as expected for the input (1, 2, 4000000). Expected result: 4613732")
	}
}
