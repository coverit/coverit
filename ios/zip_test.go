package main

import (
	"archive/zip"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestZip(t *testing.T) {

	Convey("Given root dir contains 3 gcno files", t, func() {

		zipfile := "/fixtures/src.gcno.zip"
		files := []string{
			"/fixtures/project_root/sub/1.gcno",
			"/fixtures/project_root/sub/2.gcno",
			"/fixtures/project_root/1.gcno",
		}

		err := Zip(zipfile, files)

		Convey("The value should be 3 gcno files", func() {
			So(err, ShouldBeNil)

			// Open a zip archive for reading.
			r, err := zip.OpenReader(zipfile)
			defer r.Close()
			defer os.Remove(zipfile)

			So(err, ShouldBeNil)

			gcnos := make([]string, 0, 10)
			// Iterate through the files in the archive,
			// printing some of their contents.
			for _, f := range r.File {
				gcnos = append(gcnos, f.Name)
			}

			fmt.Println("zip contains", gcnos)
			So(gcnos, ShouldContain, "/fixtures/project_root/sub/1.gcno")
			So(gcnos, ShouldContain, "/fixtures/project_root/sub/2.gcno")
			So(gcnos, ShouldContain, "/fixtures/project_root/1.gcno")

		})
	})
}
