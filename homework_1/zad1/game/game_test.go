
package game_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/MislavPeric/code-cadets-2021/homework_1/zad1/game"
)

func TestGameLogic(t *testing.T) {
	for idx, tc := range getTestCases(){
		Convey(fmt.Sprintf("Test case for #%v: %+v", idx, tc), t, func() {
			output, err := game.GameLogic(tc.start, tc.end)

			if tc.expectingError {
				So(err, ShouldNotBeNil)
			} else {
				So(err, ShouldBeNil)
				So(output, ShouldResemble, tc.expectedOutput)
			}
		})
	}
}