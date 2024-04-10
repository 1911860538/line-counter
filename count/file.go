package count

import (
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
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

func getFiles(params Params) ([]*File, error) {
	var fileInfos []*File
	if err := filepath.Walk(params.Target, func(path string, f fs.FileInfo, err error) error {
		return visit(path, f, err, &fileInfos, params)
	}); err != nil {
		return nil, err
	}

	return fileInfos, nil
}

func visit(path string, f os.FileInfo, err error, filePaths *[]*File, params Params) error {
	if err != nil {
		return err
	}

	if strings.Contains(path, "/.") {
		return nil
	}

	for _, ignoreDf := range params.IgnoreDfs {
		if path == ignoreDf {
			return nil
		}

		ps := strings.SplitN(path, ignoreDf, 2)
		if len(ps) > 1 && strings.HasPrefix(ps[1], string(filepath.Separator)) {
			return nil
		}
	}

	if !f.IsDir() {
		fileExt := filepath.Ext(f.Name())
		if fileExt == "" {
			return nil
		}
		fileExt = fileExt[1:]

		if slices.Contains(params.IgnoreTypes, fileExt) {
			return nil
		}

		if len(params.Types) > 0 {
			if !slices.Contains(params.Types, fileExt) {
				return nil
			}
		}

		fi := &File{
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
