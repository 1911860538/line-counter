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
			sizeDisplay(r.SizeSum),
			sizeDisplay(r.SizeMin),
			sizeDisplay(r.SizeMax),
			sizeDisplay(r.SizeAvg),
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

func sizeDisplay(size int64) string {
	if size == 0 {
		return "0"
	}

	if size < 1024 {
		return fmt.Sprintf("%d B", size)
	}

	if size < 1024*1024 {
		return fmt.Sprintf("%.2f KB", float64(size)/1024)
	}

	return fmt.Sprintf("%.2f MB", float64(size)/(1024*1024))
}
