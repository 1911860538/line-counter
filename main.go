package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/1911860538/line-counter/lang"
)

type File struct {
	AbsPath   string
	Name      string
	Extension string
	Size      int64
	Mode      fs.FileMode
	Line      Line
}

type Line struct {
	Total   int64
	Blank   int64
	Comment int64
}

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

func statistic() error {
	files, err := getFiles()
	if err != nil {
		return err
	}

	// todo use goroutine to getLineCount
	for i := range files {
		file := &(files[i])
		total, blank, comment, err := getLineCount(file)
		if err != nil {
			return err
		}
		file.Line = Line{
			Total:   total,
			Blank:   blank,
			Comment: comment,
		}
	}

	for _, filename := range files {
		fmt.Printf("%+v\n", filename)
	}

	//extCounter := make(map[string]int)
	//for _, fi := range files {
	//	if _, ok := extCounter[fi.Extension]; !ok {
	//		extCounter[fi.Extension] = 0
	//	}
	//
	//	extCounter[fi.Extension] += 1
	//}
	//
	//fmt.Println(extCounter)

	/*
		multiComment
	*/

	/**/

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

		staticMap[file.Extension] = row
	}

	for _, rowPtr := range staticMap {
		fmt.Printf("%+v\n", *rowPtr)
	}

	return nil
}

func getFiles() ([]File, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	var fileInfos []File
	if err := filepath.Walk(pwd, func(path string, f fs.FileInfo, err error) error {
		return visit(path, f, err, &fileInfos)
	}); err != nil {
		return nil, err
	}

	return fileInfos, err
}

func getLineCount(file *File) (total int64, blank int64, comment int64, err error) {
	fh, err := os.Open(file.AbsPath)
	if fh != nil {
		defer fh.Close()
	}
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(fh)
	inMultiCommentLine := false
	latestPrefix := ""
	latestPrefixLine := int64(0)
	for scanner.Scan() {
		total += 1
		trimLine := strings.TrimSpace(scanner.Text())

		if !inMultiCommentLine {
			isSingleLineComment, supported := lang.IsSingleLineComment(file.Extension, trimLine)
			if supported && isSingleLineComment {
				comment += 1
				continue
			}
		}

		if !inMultiCommentLine && trimLine == "" {
			blank += 1
			continue
		}

		if !inMultiCommentLine {
			isMultiLineCommentStart, prefix, supported := lang.IsMultiLineCommentStart(file.Extension, trimLine)
			if supported && isMultiLineCommentStart {
				inMultiCommentLine = true
				latestPrefix = prefix
				latestPrefixLine = total
			}
		}

		if inMultiCommentLine {
			comment += 1
			if latestPrefixLine != total {
				isMultiLineCommentEnd, supported := lang.IsMultiLineCommentEnd(file.Extension, trimLine, latestPrefix)
				if supported && isMultiLineCommentEnd {
					inMultiCommentLine = false
					latestPrefix = ""
					latestPrefixLine = int64(0)
				}
			}
		}
	}
	return
}

func visit(path string, f os.FileInfo, err error, filePaths *[]File) error {
	if err != nil {
		return err
	}

	if strings.Contains(path, "/.") {
		return nil
	}

	if !f.IsDir() {
		fileExt := filepath.Ext(f.Name())
		if fileExt == "" {
			return nil
		}

		fi := File{
			AbsPath:   path,
			Name:      f.Name(),
			Extension: fileExt,
			Size:      f.Size(),
			Mode:      f.Mode(),
		}
		*filePaths = append(*filePaths, fi)
	}

	return nil
}

func main() {
	start := time.Now()

	if err := statistic(); err != nil {
		panic(err)
	}

	cost := time.Now().Sub(start).Milliseconds()
	fmt.Println(cost)
}
