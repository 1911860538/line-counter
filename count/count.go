package count

import (
	"bufio"
	"context"
	"os"
	"strings"

	"golang.org/x/sync/errgroup"

	"github.com/1911860538/line-counter/lang"
)

func multiSetLineCount(files []*File) error {
	eg, ctx := errgroup.WithContext(context.Background())
	eg.SetLimit(4)

	for i := range files {
		file := files[i]

		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return nil
			default:
				if err := setLineCount(file); err != nil {
					return err
				}
				return nil
			}
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}

func setLineCount(file *File) error {
	fh, err := os.Open(file.AbsPath)
	if fh != nil {
		defer fh.Close()
	}
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(fh)
	inMultiCommentLine := false
	latestPrefix := ""
	latestPrefixLine := int64(0)

	var (
		total   int64
		comment int64
		blank   int64
	)

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

	file.Line = Line{
		Total:   total,
		Blank:   blank,
		Comment: comment,
	}

	return nil
}
