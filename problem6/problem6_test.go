package problem6

import "testing"

func TestSumSquareDifference(test *testing.T) {
	// Unit Testcase: 1 : Project Euler Problem statement
	testResult := SumSquareDifference(100)
	test.Logf("The result of the method call is: %d", testResult)
	if testResult != 25164150 {
		test.Errorf("The result is not as expected for the input (100). Expected result: 25164150")
	}

	testResult2 := SumSquareDifference(10)
	test.Logf("The result of the method call is: %d", testResult2)
	if testResult2 != 2640 {
		test.Errorf("The result is not as expected for the input (10). Expected result: 2640")
	}

	testResult3 := SumSquareDifference(-11)
	test.Logf("The result of the method call is: %d", testResult3)
	if testResult3 != -1 {
		test.Errorf("The result is not as expected for the input (-11). Expected result: -1")
	}

	testResult4 := SumSquareDifference(5)
	test.Logf("The result of the method call is: %d", testResult4)
	if testResult4 != 170 {
		test.Errorf("The result is not as expected for the input (5). Expected result: 170")
	}
}
