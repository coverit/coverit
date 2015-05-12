package main

import (
	"archive/zip"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"path"
	"testing"
)

func TestZip(t *testing.T) {

	pwd, _ := os.Getwd()

	Convey("Given root dir contains 3 gcno files", t, func() {

		zipfile := path.Join(pwd, "/fixtures/src.gcno.zip")
		files := []string{
			path.Join(pwd, "/fixtures/project_root/sub/1.gcno"),
			path.Join(pwd, "/fixtures/project_root/sub/2.gcno"),
			path.Join(pwd, "/fixtures/project_root/1.gcno"),
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

			So(gcnos, ShouldContain, path.Join(pwd, "/fixtures/project_root/sub/1.gcno"))
			So(gcnos, ShouldContain, path.Join(pwd, "/fixtures/project_root/sub/2.gcno"))
			So(gcnos, ShouldContain, path.Join(pwd, "/fixtures/project_root/1.gcno"))

		})
	})
}
