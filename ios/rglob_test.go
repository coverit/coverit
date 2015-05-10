package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRglob(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given root dir contains 3 gcno files", t, func() {

		files, err := Rglob("/fixtures/project_root", ".gcno")

		Convey("The value should be 3 gcno files", func() {
			So(err, ShouldBeNil)
			So(len(files), ShouldEqual, 3)
			So(files, ShouldContain, "/fixtures/project_root/sub/1.gcno")
			So(files, ShouldContain, "/fixtures/project_root/sub/2.gcno")
			So(files, ShouldContain, "/fixtures/project_root/1.gcno")
		})
	})

	Convey("Given root dir contains 3 gcda files", t, func() {

		files, err := Rglob("/fixtures/project_root", ".gcda")

		Convey("The value should be 3 gcda files", func() {
			So(err, ShouldBeNil)
			So(len(files), ShouldEqual, 3)
			So(files, ShouldContain, "/fixtures/project_root/sub/1.gcda")
			So(files, ShouldContain, "/fixtures/project_root/sub/2.gcda")
			So(files, ShouldContain, "/fixtures/project_root/1.gcda")
		})
	})
}
