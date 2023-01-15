package problem4

import (
	"fmt"
	"testing"
)

func TestLargestPalindromeProduct(test *testing.T) {
	// Unit Testcase : Project Euler Problem statement
	testResult := LargestPalindromeProduct(3)
	test.Logf("The result of the method call is: %d", testResult)
	if testResult != 906609 {
		test.Errorf("The result is not as expected for the input (3). Expected result: 906609")
	}

	testResult2 := LargestPalindromeProduct(2)
	test.Logf("The result of the method call is: %d", testResult2)
	if testResult2 != 9009 {
		test.Errorf("The result is not as expected for the input (2). Expected result: 9009")
	}

	testResult3 := LargestPalindromeProduct(-8)
	test.Logf("The result of the method call is: %d", testResult3)
	if testResult3 != -1 {
		test.Errorf("The result is not as expected for the input (-8). Expected result: -1")
	}

	testResult4 := LargestPalindromeProduct(15)
	test.Logf("The result of the method call is: %d", testResult4)
	if testResult4 != -1 {
		test.Errorf("The result is not as expected for the input (15). Expected result: -1")
	}
}

func BenchmarkLargestPalindromeProduct(ben *testing.B) {
	inputs := []int8{3, 2, -5, 1}
	outputs := []int64{906609, 9009, -1, 9}
	for j := 0; j < len(inputs); j++ {
		for i := 0; i < ben.N; i++ {
			res := LargestPalindromeProduct(inputs[j])
			if res != outputs[j] {
				ben.Logf("The result is not as expected for input (" + fmt.Sprintf("%d", inputs[j]) + "). Expected result: " + fmt.Sprintf("%d", outputs[j]))
			}
		}
	}
}
