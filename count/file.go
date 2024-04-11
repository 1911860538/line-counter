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

	fileExt := filepath.Ext(f.Name())
	if !f.IsDir() && path == params.Target {
		fi := &File{
			AbsPath:   path,
			Name:      f.Name(),
			Extension: fileExt,
			Size:      f.Size(),
			Mode:      f.Mode(),
		}
		*filePaths = append(*filePaths, fi)
		return nil
	}

	pathSep := string(filepath.Separator)
	if strings.Contains(path, pathSep+".") {
		return nil
	}

	for _, ignoreDf := range params.IgnoreDfs {
		if path == ignoreDf {
			return nil
		}

		if isSubPath(ignoreDf, path, pathSep) {
			return nil
		}
	}

	if !f.IsDir() {
		if fileExt == "" {
			return nil
		}

		if slices.Contains(params.IgnoreTypes, fileExt[1:]) {
			return nil
		}

		if len(params.Types) > 0 {
			if !slices.Contains(params.Types, fileExt[1:]) {
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

func isSubPath(dir string, subPath string, pathSep string) bool {
	ps := strings.SplitN(dir, subPath, 2)
	if len(ps) > 1 && strings.HasPrefix(ps[1], pathSep) {
		return true
	}

	return false
}
