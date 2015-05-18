package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"path"
	"testing"
)

func TestLcovParser(t *testing.T) {

	pwd, _ := os.Getwd()

	Convey("Given a lcov file with coverage data", t, func() {

		lcovFile := path.Join(pwd, "fixtures", "coverage.info")

		Convey("When call lcov parser", func() {

			report, err := ParseLcov(lcovFile)

			Convey("Should returns expected coverage report and no error occurs", func() {
				So(err, ShouldBeNil)
				So(len(report), ShouldBeGreaterThan, 0)
				So(report[0].File, ShouldStartWith, "/coverit/file.js")
			})
		})
	})
}
