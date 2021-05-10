// Package contains unit tests for FizzBuzz.
package fizzbuzz_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"code-cadets-2021/homework_1/Task-01/fizzbuzz"
)

// TestGetFizzBuzz tests fizzbuzz.GetFizzBuzz method
func TestGetFizzBuzz(t *testing.T) {
	for idx, tc := range getTestCases() {
		Convey(fmt.Sprintf("Given test case #%v: %+v", idx, tc), t, func() {

			actualOutput, actualErr := fizzbuzz.GetFizzBuzz(tc.inputStart, tc.inputEnd)

			if tc.expectingError {
				So(actualErr, ShouldNotBeNil)
			} else {
				So(actualErr, ShouldBeNil)
				So(actualOutput, ShouldResemble, tc.expectedOutput)
			}
		})
	}
}
