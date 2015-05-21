package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateReport(t *testing.T) {

	Convey("Given the `data` field of create_task notification", t, func() {
		// report := Report{
		// 	RunAt: Runtime{
		// 		Platform: "p"
		// 		Device: "d"
		// 		SDK: "sdk"
		// 	}
		//
		// }
		//
		// stats := []CoverageStat{
		// 	CoverageStat{
		// 		Lines: 1,
		// 	},
		// }
		// functions := []CoverageFunction{
		// 	CoverageFunction{
		// 		LineNo: 100,
		// 		Hit:    12,
		// 	},
		// }
		// lines := []CoverageLine{
		// 	CoverageLine{
		// 		Source: "file1",
		// 		LineNo: 100,
		// 		Hit:    12,
		// 	},
		// }
		// tree := CoverageTree{
		// 	Stats:     stats,
		// 	Functions: functions,
		// 	Lines:     lines,
		// }
		//

		client := Client{Host: "http://localhost:5803"}

		Convey("When decode as struct", func() {

			report, err := client.CreateReport()

			Convey("Should parsed succeed", func() {
				So(err, ShouldBeNil)
				So(report.Id, ShouldNotBeNil)
			})
		})
	})
}
