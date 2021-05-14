package taxes_test

import (
	"code-cadets-2021/homework_1/task_02/taxes"
	"fmt"
	"math"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCalculateTax(t *testing.T) {
	percentages := []float64{0.0, 0.10, 0.20, 0.30}
	amounts := []float64{1000, 5000, 10000, math.Inf(1)}
	for idx, tc := range getTestCases() {
		Convey(fmt.Sprintf("Given test case #%v: %+v", idx, tc), t, func() {

			output, err := taxes.CalculateTax(tc.input, percentages, amounts)

			if tc.error {
				So(err, ShouldNotBeNil)
			} else {
				So(err, ShouldBeNil)
				So(output, ShouldResemble, tc.output)
			}
		})
	}
}