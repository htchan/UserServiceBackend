package utils

import (
	"testing"
)

func TestContainString(t *testing.T) {
	testcases := []struct{
		inputListStr string
		inputDelimiter string
		inputTarget string
		expectResult bool
	} {
		{ "1,2,3", ",", "1", true },
		{ "", ",", "1", false },
		{ "1-2-3", "-", "4", false },
	}
	for _, testcase := range testcases {
		actualResult := ContainString(
			testcase.inputListStr, testcase.inputDelimiter, testcase.inputTarget)
		if actualResult != testcase.expectResult {
			t.Fatalf("utils.Cotnains(%v, %v) returns %v, but not %v",
				testcase.inputListStr, testcase.inputTarget, actualResult, testcase.expectResult)
		}
	}
}