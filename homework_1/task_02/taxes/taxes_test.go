package taxes_test

import (
	"code-cadets-2021/homework_1/task_02/taxes"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCalculateTax(t *testing.T) {
	for idx, tc := range getTestCases() {
		Convey(fmt.Sprintf("Given test case #%v: %+v", idx, tc), t, func() {

			output, err := taxes.CalculateTax(tc.input, tc.percentages, tc.amounts)

			if tc.error {
				So(err, ShouldNotBeNil)
			} else {
				So(err, ShouldBeNil)
				So(output, ShouldResemble, tc.output)
			}
		})
	}
}
