package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
  "os"
)

func TestGetBranchName(t *testing.T) {

  Convey("Only set branch name into environment variable", t, func() {

    os.Setenv("coverit_branch_name", "master")

    defer os.Unsetenv("coverit_branch_name")

    branch := GetBranchName()

    Convey("The branch name got should be master", func() {

      So(branch, ShouldEqual, "master")
    })
  })

  Convey("Set branch name both into environment variable and yml file", t, func(){

    os.Setenv("coverit_branch_name", "branch_1")

    var data =
    `coverit_branch_name: master`

    fout, _ := os.Create(".coverit.yml")

    defer os.Remove(".coverit.yml")
    defer fout.Close()
    defer os.Unsetenv("coverit_branch_name")

    fout.WriteString(data)

    branch := GetBranchName()

    Convey("The branch name got should be branch_1", func() {

      So(branch, ShouldEqual, "branch_1")
    })
  })

  Convey("Only set branch name into yml file", t, func(){

    var data =
    `coverit_branch_name: master`

    fout, _ := os.Create(".coverit.yml")

    defer os.Remove(".coverit.yml")
    defer fout.Close()

    fout.WriteString(data)

    branch := GetBranchName()

    Convey("The branch name got should be master", func() {

      So(branch, ShouldEqual, "master")
    })
  })
}
