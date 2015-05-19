package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBranch(t *testing.T) {

	Convey("Enter fixtures testing repo", t, func() {

		repo, _ := OpenRepo("./fixtures/git_read_test")

		Convey("The head branch name should be 'master'", func() {

			branch, _ := repo.Branch()

			So(branch, ShouldEqual, "master")
		})
	})
}

func TestCommitSha(t *testing.T) {

	Convey("Enter fixtures testing repo", t, func() {

		repo, _ := OpenRepo("./fixtures/git_read_test")

		Convey("The head commit sha should be '0bb5b4d2ad6ea3480f4ecfafe83f521ba880c729'", func() {

			sha, _ := repo.CommitSha()

			So(sha, ShouldEqual, "0bb5b4d2ad6ea3480f4ecfafe83f521ba880c729")
		})
	})
}

func TestAuthor(t *testing.T) {

	Convey("Enter fixtures testing repo", t, func() {

		repo, _ := OpenRepo("./fixtures/git_read_test")
		author, _ := repo.Author()

		Convey("The head author name should be 'enzo'", func() {

			So(author.Name, ShouldEqual, "enzo")
		})

		Convey("The head author email should be 'enzo@coverit.io'", func() {

			So(author.Email, ShouldEqual, "enzo@coverit.io")
		})

	})
}

func TestCommitter(t *testing.T) {

	Convey("Enter fixtures testing repo", t, func() {

		repo, _ := OpenRepo("./fixtures/git_read_test")
		committer, _ := repo.Committer()

		Convey("The head committer name should be 'enzo'", func() {

			So(committer.Name, ShouldEqual, "enzo")
		})

		Convey("The head committer email should be 'enzo@coverit.io'", func() {

			So(committer.Email, ShouldEqual, "enzo@coverit.io")
		})

	})
}

func TestCommitMessage(t *testing.T) {

	Convey("Enter fixtures testing repo", t, func() {

		repo, _ := OpenRepo("./fixtures/git_read_test")

		Convey("The head commit message should be 'Initial commit'", func() {

			commitMsg, _ := repo.CommitMessage()

			So(commitMsg, ShouldEqual, "Initial commit\n")
		})
	})

}
