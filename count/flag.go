package count

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type Params struct {
	Target      string
	IgnoreDfs   []string
	Types       []string
	IgnoreTypes []string
}

func parseFlags() (Params, error) {
	var (
		pwd string

		target      string
		ignoreDfs   string
		types       string
		ignoreTypes string
	)

	flag.StringVar(&target, "n", "", "target directory or file for statistics. Default is the current directory.")
	flag.StringVar(&ignoreDfs, "in", "", "ignore directories or files from statistics.\nFor example: a/b/,a/c/x.txt")
	flag.StringVar(&types, "t", "", "file types for statistics.\nFor example: txt,go,py")
	flag.StringVar(&ignoreTypes, "it", "", "ignore types from statistics.\nFor example: json,xml")

	flag.Parse()

	params := Params{
		Target:      strings.TrimSpace(target),
		IgnoreDfs:   make([]string, 0),
		Types:       make([]string, 0),
		IgnoreTypes: make([]string, 0),
	}

	if params.Target == "" {
		var err error
		pwd, err = os.Getwd()
		if err != nil {
			return Params{}, err
		}
		params.Target = pwd
	}

	_, err := os.Stat(params.Target)
	if err != nil {
		return Params{}, err
	}

	targetExt := filepath.Ext(params.Target)
	if targetExt != "" {
		targetExt = targetExt[1:]
	}

	for _, idf := range strings.Split(strings.TrimSpace(ignoreDfs), ",") {
		if idf == "" {
			continue
		}

		if idf == params.Target {
			return Params{}, fmt.Errorf("%s: 'in'(ignore directory or file from statistics) equals to directory or file for statistics", idf)
		}

		absPath, err := getAbsolutePath(idf)
		if err != nil {
			return Params{}, err
		}

		if !slices.Contains(params.IgnoreDfs, absPath) {
			params.IgnoreDfs = append(params.IgnoreDfs, absPath)
		}
	}

	for _, t := range strings.Split(strings.TrimSpace(types), ",") {
		if t != "" && !slices.Contains(params.Types, t) {
			params.Types = append(params.Types, t)
		}
	}

	for _, it := range strings.Split(strings.TrimSpace(ignoreTypes), ",") {
		if it == "" {
			continue
		}

		if it == targetExt {
			return Params{}, fmt.Errorf("'it'(ignore type from statistics) is conflicted with the target file(%s) type", params.Target)
		}

		if !slices.Contains(params.IgnoreTypes, it) {
			params.IgnoreTypes = append(params.IgnoreTypes, it)
		}
	}

	if len(params.Types) > 0 && len(params.IgnoreTypes) > 0 {
		return Params{}, errors.New("the parameters 't' and 'it' cannot be used simultaneously")
	}

	return params, nil
}

func getAbsolutePath(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return "", err
	}

	return absPath, nil
}
