package count

import (
	"cmp"
	"slices"
)

type StatisticRow struct {
	Extension    string
	Count        int64
	SizeSum      int64
	SizeMin      int64
	SizeMax      int64
	SizeAvg      int64
	Lines        int64
	LinesMin     int64
	LinesMax     int64
	LinesAvg     int64
	LinesCode    int64
	LinesComment int64
	LinesBlank   int64
}

func static(files []*File) []*StatisticRow {
	staticMap := make(map[string]*StatisticRow)
	for _, file := range files {
		row, ok := staticMap[file.Extension]
		if !ok {
			row = &StatisticRow{
				Extension: file.Extension,
			}
			staticMap[file.Extension] = row
		}

		row.Count += 1

		if row.Count == 1 {
			row.SizeMin = file.Size
			row.LinesMin = file.Line.Total
		}

		row.SizeSum += file.Size
		if file.Size < row.SizeMin {
			row.SizeMin = file.Size
		}
		if file.Size > row.SizeMax {
			row.SizeMax = file.Size
		}
		row.SizeAvg = row.SizeSum / row.Count
		row.Lines += file.Line.Total
		if file.Line.Total < row.LinesMin {
			row.LinesMin = file.Line.Total
		}
		if file.Line.Total > row.LinesMax {
			row.LinesMax = file.Line.Total
		}
		row.LinesAvg = row.Lines / row.Count
		row.LinesCode += file.Line.Total - file.Line.Blank - file.Line.Comment
		row.LinesComment += file.Line.Comment
		row.LinesBlank += file.Line.Blank
	}

	var staticSlice []*StatisticRow
	for _, rowPtr := range staticMap {
		staticSlice = append(staticSlice, rowPtr)
	}
	slices.SortFunc(staticSlice, func(a, b *StatisticRow) int {
		return cmp.Compare(a.Extension, b.Extension)
	})

	return staticSlice
}
