package problem10

import (
	"testing"
)

func TestPrimeSum(test *testing.T) {
	// Unit Testcase: 1 : Project EulerProblem statement
	testResult := PrimeSum(2000000)
	test.Logf("The result of the method call is: %d", testResult)
	if testResult != 142913828922 {
		test.Errorf("The result is not as expected for the input (2000000). Expected result: 142913828922")
	}

	testResult2 := PrimeSum(10)
	test.Logf("The result of the method call is: %d", testResult2)
	if testResult2 != 17 {
		test.Errorf("The result is not as expected for the input (10). Expected result: 17")
	}

	testResult3 := PrimeSum(-1)
	test.Logf("The result of the method call is: %d", testResult3)
	if testResult3 != -1 {
		test.Errorf("The result is not as expected for the input (-1). Expected result: -1")
	}

	testResult4 := PrimeSum(1)
	test.Logf("The result of the method call is: %d", testResult4)
	if testResult4 != 0 {
		test.Errorf("The result is not as expected for the input (1). Expected result: 0")
	}
}
