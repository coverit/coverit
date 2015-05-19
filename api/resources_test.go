package main

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func TestCreateBuild(t *testing.T) {

	session, _ := mgo.Dial(os.Getenv("MONGO_PORT_27017_TCP_ADDR"))
	db := session.DB(os.Getenv("DB_NAME"))

	Convey("Given a simple build", t, func() {

		buildResource := &BuildResource{db: db}
		defer buildResource.db.DropDatabase()

		project := "foo project"
		git := GitInfo{Commit: "1x"}
		service := "foo service"

		Convey("When create with resource", func() {

			build, err := buildResource.CreateBuild(project, git, service)

			Convey("Should succeed", func() {
				So(err, ShouldBeNil)
				So(build.Id, ShouldNotBeNil)
				So(build.Git, ShouldResemble, git)
				So(build.ProjectName, ShouldEqual, project)
				So(build.Service, ShouldEqual, service)
			})

			Convey("Should create a build in db", func() {
				var builds []Build

				err := buildResource.db.C("builds").Find(nil).All(&builds)
				So(err, ShouldBeNil)
				So(len(builds), ShouldEqual, 1)
				So(builds[0].Id, ShouldEqual, build.Id)
				So(builds[0].Git, ShouldResemble, build.Git)
				So(builds[0].ProjectName, ShouldEqual, build.ProjectName)
				So(builds[0].Service, ShouldEqual, build.Service)
			})
		})
	})
}

func TestListBuilds(t *testing.T) {

	session, _ := mgo.Dial(os.Getenv("MONGO_PORT_27017_TCP_ADDR"))
	db := session.DB(os.Getenv("DB_NAME"))

	Convey("Given 3 builds in db", t, func() {

		buildResource := &BuildResource{db: db}
		defer buildResource.db.DropDatabase()

		build1 := Build{
			Id:          bson.NewObjectId(),
			ProjectName: "foo project",
		}
		build2 := Build{
			Id:          bson.NewObjectId(),
			ProjectName: "bar project",
		}
		build3 := Build{
			Id:          bson.NewObjectId(),
			ProjectName: "bar project",
		}

		buildResource.db.C("builds").Insert(build1)
		buildResource.db.C("builds").Insert(build2)
		buildResource.db.C("builds").Insert(build3)

		Convey("When get projects by name", func() {

			builds, err := buildResource.ListBuilds("bar project")

			Convey("Should return 2 builds", func() {
				So(err, ShouldBeNil)
				So(len(builds), ShouldEqual, 2)
			})

			Convey("Should return in order", func() {
				So(builds[0].Id, ShouldEqual, build2.Id)
				So(builds[1].Id, ShouldEqual, build3.Id)
			})
		})
	})
}
