package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"path"
	"testing"
)

func TestRglob(t *testing.T) {

	pwd, _ := os.Getwd()

	// Only pass t into top-level Convey calls
	Convey("Given root dir contains 3 gcno files", t, func() {

		files, err := Rglob(path.Join(pwd, "fixtures/project_root"), ".gcno")

		Convey("The value should be 3 gcno files", func() {
			So(err, ShouldBeNil)
			So(len(files), ShouldEqual, 3)
			So(files, ShouldContain, path.Join(pwd, "fixtures/project_root/sub/1.gcno"))
			So(files, ShouldContain, path.Join(pwd, "fixtures/project_root/sub/2.gcno"))
			So(files, ShouldContain, path.Join(pwd, "fixtures/project_root/1.gcno"))
		})
	})

	Convey("Given root dir contains 3 gcda files", t, func() {

		files, err := Rglob(path.Join(pwd, "/fixtures/project_root"), ".gcda")

		Convey("The value should be 3 gcda files", func() {
			So(err, ShouldBeNil)
			So(len(files), ShouldEqual, 3)
			So(files, ShouldContain, path.Join(pwd, "/fixtures/project_root/sub/1.gcda"))
			So(files, ShouldContain, path.Join(pwd, "/fixtures/project_root/sub/2.gcda"))
			So(files, ShouldContain, path.Join(pwd, "/fixtures/project_root/1.gcda"))
		})
	})
}
