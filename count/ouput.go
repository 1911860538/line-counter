package count

import (
	"fmt"
	"log"

	"github.com/scylladb/termtables"
)

func output(params Params, staticSlice []*StatisticRow) {
	t := termtables.CreateTable()
	t.AddTitle("line-counter")
	t.AddHeaders(
		"Extension",
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
		"LinesBlank",
	)

	for _, r := range staticSlice {
		t.AddRow(
			r.Extension,
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

	log.Printf("The target is `%s`.\n\n", params.Target)
	fmt.Println(t.Render())
}
