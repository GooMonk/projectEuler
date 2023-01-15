package problem7

import (
	"fmt"
	"testing"
)

func TestNthPrime(test *testing.T) {
	// Unit Testcase: 1 : Project Euler Problem statement
	testResult := NthPrime(10001)
	test.Logf("The result of the method call is: %d", testResult)
	if testResult != 104743 {
		test.Errorf("The result is not as expected for the input (10001). Expected result: 104743")
	}

	testResult2 := NthPrime(10)
	test.Logf("The result of the method call is: %d", testResult2)
	if testResult2 != 29 {
		test.Errorf("The result is not as expected for the input (10). Expected result: 29")
	}

	testResult3 := NthPrime(-11)
	test.Logf("The result of the method call is: %d", testResult3)
	if testResult3 != -1 {
		test.Errorf("The result is not as expected for the input (-11). Expected result: -1")
	}

	testResult4 := NthPrime(3)
	test.Logf("The result of the method call is: %d", testResult4)
	if testResult4 != 5 {
		test.Errorf("The result is not as expected for the input (3). Expected result: 5")
	}
}

func BenchmarkNthPrime(ben *testing.B) {
	inputs := []int64{10001, 10, -11, 3}
	outputs := []int64{104743, 29, -1, 5}
	for j := 0; j < len(inputs); j++ {
		for i := 0; i < ben.N; i++ {
			res := NthPrime(inputs[j])
			if res != outputs[j] {
				ben.Logf("The result is not as expected for input (" + fmt.Sprintf("%d", inputs[j]) + "). Expected result: " + fmt.Sprintf("%d", outputs[j]))
			}
		}
	}

}
