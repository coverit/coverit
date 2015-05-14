package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
)

func TestLcov(t *testing.T) {

	pwd, _ := os.Getwd()

	// Only pass t into top-level Convey calls
	// $ lcov --compat split_crc=off --capture --directory fixtures --output-file coverage.info
	// genhtml -o coverage.html coverage.info
	Convey("Given dir contains 3 gcno files", t, func() {

		dir := path.Join(pwd, "fixtures")
		info := "coverage.info"
		defer os.Remove(info)

		Convey("When call Lcov()", func() {

			err := Lcov(dir, info)

			Convey("Should returns nil and info is created correctlly", func() {
				So(err, ShouldBeNil)

				content, err := ioutil.ReadFile(info)
				So(err, ShouldBeNil)
				So(strings.HasPrefix(string(content), "TN:"), ShouldBeTrue)
			})
		})
	})
}
