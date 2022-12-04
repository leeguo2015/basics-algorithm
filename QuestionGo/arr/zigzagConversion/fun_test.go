package zigzagConversion

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCreateSLParams(t *testing.T) {

	Convey("ZigzagConversion", t, func() {

		So(ZigzagConversion("PAYPALISHIRING", 3), ShouldEqual, "PAHNAPLSIIGYIR")
		So(ZigzagConversion("PAYPALISHIRING", 4), ShouldEqual, "PINALSIGYAHRPI")
		So(ZigzagConversion("PAYPALISHIRING", 1), ShouldEqual, "PAYPALISHIRING")
	})

}
func TestZigzagConversionPolling(t *testing.T) {

	Convey("ZigzagConversion", t, func() {

		So(ZigzagConversionPolling("PAYPALISHIRING", 3), ShouldEqual, "PAHNAPLSIIGYIR")
		//So(ZigzagConversionPolling("PAYPALISHIRING", 4), ShouldEqual, "PINALSIGYAHRPI")
		//So(ZigzagConversionPolling("PAYPALISHIRING", 1), ShouldEqual, "PAYPALISHIRING")
	})

}
