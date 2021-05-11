package taxes_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"code-cadets-2021/homework_1/02_ProgressiveTax/taxes"
)

func TestNewTaxBrackets(t *testing.T) {
	for idx, tc := range getNewTaxBracketsTestCases() {
		Convey(fmt.Sprintf("Given test case #%v: %+v", idx, tc), t, func() {
			_, actualErr := taxes.NewTaxBrackets(tc.inputBrackets, tc.inputTaxes)

			if tc.expectingError {
				So(actualErr, ShouldNotBeNil)
			} else {
				So(actualErr, ShouldBeNil)
			}
		})
	}
}

func TestCalculateProgressiveTax(t *testing.T) {
	for idx, tc := range getCalculateProgressiveTaxTestCases() {
		Convey(fmt.Sprintf("Given test case #%v: %+v", idx, tc), t, func() {
			taxBrackets, _ := taxes.NewTaxBrackets(tc.inputBrackets, tc.inputTaxes)
			actualOutput, actualErr := taxes.CalculateProgressiveTax(tc.inputValue, taxBrackets)

			if tc.expectingError {
				So(actualErr, ShouldNotBeNil)
			} else {
				So(actualErr, ShouldBeNil)
				So(actualOutput, ShouldResemble, tc.expectedOutput)
			}
		})
	}
}