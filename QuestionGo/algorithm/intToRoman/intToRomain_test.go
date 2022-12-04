package intToRoman

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCreateSLParams(t *testing.T) {
	Convey("ZigzagConversion", t, func() {
		So(IntToRoman(1994), ShouldEqual, "MCMXCIV")
	})
}
