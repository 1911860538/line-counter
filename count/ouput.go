package count

import (
	"fmt"

	"github.com/scylladb/termtables"
)

func output(staticSlice []*StatisticRow) {
	t := termtables.CreateTable()
	t.AddTitle("line-counter")
	t.AddHeaders(
		"Extension",
		"Count",
		"SizeSum",
		"Count",
		"SizeSum",
		"SizeMin",
		"SizeMax",
		"SizeAvg",
		"Lines",
		"LinesMin",
		"LinesMax",
		"LinesAvg",
		"LinesCode",
		"LinesComment",
		"LinesBlank")

	for _, r := range staticSlice {
		t.AddRow(
			r.Extension,
			r.Count,
			r.SizeSum,
			r.Count,
			r.SizeSum,
			r.SizeMin,
			r.SizeMax,
			r.SizeAvg,
			r.Lines,
			r.LinesMin,
			r.LinesMax,
			r.LinesAvg,
			r.LinesCode,
			r.LinesComment,
			r.LinesBlank,
		)
	}

	fmt.Println(t.Render())
}
