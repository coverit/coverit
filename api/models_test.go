package main

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/mgo.v2/bson"
)

func TestReportBsonMarshal(t *testing.T) {

	Convey("Given the `data` field of create_task notification", t, func() {

		// report := Report{
		// 	RunAt: Runtime{
		// 		Platform: "p"
		// 		Device: "d"
		// 		SDK: "sdk"
		// 	}
		//
		// }

		stats := []CoverageStat{
			CoverageStat{
				Lines: 1,
			},
		}
		functions := []CoverageFunction{
			CoverageFunction{
				LineNo: 100,
				Hit:    12,
			},
		}
		lines := []CoverageLine{
			CoverageLine{
				Source: "file1",
				LineNo: 100,
				Hit:    12,
			},
		}
		tree := CoverageTree{
			Stats:     stats,
			Functions: functions,
			Lines:     lines,
		}

		Convey("When decode as struct", func() {

			bytes, err := bson.Marshal(tree)
			fmt.Printf("%q", bytes)
			var a CoverageTree
			err = bson.Unmarshal(bytes, &a)

			Convey("Should parsed succeed", func() {
				So(err, ShouldBeNil)
				So(a, ShouldResemble, tree)
			})
		})
	})
}
