package fizzbuzz_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"code-cadets-2021/homework_1/Task-01/fizzbuzz"
)

// TestGetFizzBuzz tests fizzbuzz.GetFizzBuzz method
func TestGetFizzBuzz(t *testing.T) {
	for index, unitTestCase := range getTestCases() {
		Convey(fmt.Sprintf("Given test case #%v: %+v", index, unitTestCase), t, func() {

			actualOutput, actualErr := fizzbuzz.GetFizzBuzz(unitTestCase.inputStart, unitTestCase.inputEnd)

			if unitTestCase.expectingError {
				So(actualErr, ShouldNotBeNil)
			} else {
				So(actualErr, ShouldBeNil)
				So(actualOutput, ShouldResemble, unitTestCase.expectedOutput)
			}
		})
	}
}
