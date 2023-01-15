package problem5

import (
	"fmt"
	"testing"
)

func TestPrimeSum(test *testing.T) {
	// Unit Testcase: 1 : Project Euler Problem statement
	testResult := SmallestMultiple(20)
	test.Logf("The result of the method call is: %d", testResult)
	if testResult != 232792560 {
		test.Errorf("The result is not as expected for the input (20). Expected result: 232792560")
	}

	testResult2 := SmallestMultiple(10)
	test.Logf("The result of the method call is: %d", testResult2)
	if testResult2 != 2520 {
		test.Errorf("The result is not as expected for the input (10). Expected result: 2520")
	}

	testResult3 := SmallestMultiple(-10)
	test.Logf("The result of the method call is: %d", testResult3)
	if testResult3 != -1 {
		test.Errorf("The result is not as expected for the input (-10). Expected result: -1")
	}

	testResult4 := SmallestMultiple(3)
	test.Logf("The result of the method call is: %d", testResult4)
	if testResult4 != 6 {
		test.Errorf("The result is not as expected for the input (3). Expected result: 6")
	}
}

func BenchmarkPrimeSum(ben *testing.B) {
	inputs := []int64{20, 10, -5, 2}
	outputs := []int64{232792560, 2520, -1, 2}
	for j := 0; j < len(inputs); j++ {
		for i := 0; i < ben.N; i++ {
			res := SmallestMultiple(inputs[j])
			if res != outputs[j] {
				ben.Logf("The result is not as expected for input (" + fmt.Sprintf("%d", inputs[j]) + "). Expected result: " + fmt.Sprintf("%d", outputs[j]))
			}
		}
	}

}
