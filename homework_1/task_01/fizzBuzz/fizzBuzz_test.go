package fizzBuzz_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"code-cadets-2021/homework_1/task_01/fizzBuzz"
)

func TestFizzBuzzPlay(t *testing.T) {
	for i, tc := range getTestCases() {
		Convey(fmt.Sprintf("Given test case #%v: %+v", i, tc), t, func() {

			actualOutput, actualErr := fizzBuzz.FizzBuzzPlay(tc.inputStart, tc.inputEnd)

			if tc.expectingError {
				So(actualErr, ShouldNotBeNil)
			} else {
				So(actualErr, ShouldBeNil)
				So(actualOutput, ShouldResemble, tc.expectedOutput)
			}

		})
	}
}
