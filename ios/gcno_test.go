package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGcno(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given root dir contains 3 gcno files", t, func() {

		gcnos, err := CollectGcno("/fixtures/project_root")

		Convey("The value should be 3 gcno files", func() {
			So(err, ShouldBeNil)
			So(len(gcnos), ShouldEqual, 3)
			So(gcnos, ShouldContain, "/fixtures/project_root/sub/1.gcno")
			So(gcnos, ShouldContain, "/fixtures/project_root/sub/2.gcno")
			So(gcnos, ShouldContain, "/fixtures/project_root/1.gcno")
		})
	})
}
