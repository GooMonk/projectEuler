package problem1

import (
	"testing"
)

func TestMultipleSum(test *testing.T) {
	// Unit Testcase: 1 : Project Euler
	testResult := MultipleSum(3, 5, 1000)
	test.Logf("The result of the method call is: %d", testResult)
	if testResult != 233168 {
		test.Errorf("The result is not as expected for the input (3, 5, 1000). Expected result: 233168")
	}

	testResult2 := MultipleSum(5, 5, 10)
	test.Logf("The result of the method call is: %d", testResult2)
	if testResult2 != 5 {
		test.Errorf("The result is not as expected for the input (5, 5, 10). Expected result: 5")
	}

	testResult3 := MultipleSum(11, 5, 10)
	test.Logf("The result of the method call is: %d", testResult3)
	if testResult3 != -1 {
		test.Errorf("The result is not as expected for the input (11, 5, 10). Expected result: -1")
	}

	testResult4 := MultipleSum(-3, 5, 10)
	test.Logf("The result of the method call is: %d", testResult4)
	if testResult4 != -1 {
		test.Errorf("The result is not as expected for the input (-5, 5, 10). Expected result: -1")
	}
}

func TestMultipleSumUpdated(test2 *testing.T) {
	//Cases where one of the input number is a factor of the other number
	testResult := MultipleSum(2, 8, 1000)
	test2.Logf("The result of the method call is: %d", testResult)
	if testResult != 249500 {
		test2.Errorf("The result is not as expected for the input (2, 8, 1000). Expected result: 249500")
	}

	testResult2 := MultipleSum(30, 15, 31)
	test2.Logf("The result of the method call is: %d", testResult2)
	if testResult2 != 45 {
		test2.Errorf("The result is not as expected for the input (30, 15, 31). Expected result: 45")
	}

	testResult3 := MultipleSum(1000, 100, 1000)
	test2.Logf("The result of the method call is: %d", testResult3)
	if testResult3 != -1 {
		test2.Errorf("The result is not as expected for the input (1000, 100, 1000). Expected result: -1")
	}
}
