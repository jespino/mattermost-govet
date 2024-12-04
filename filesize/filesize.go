// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package filesize

import (
	"strings"

	"github.com/mattermost/mattermost-govet/helpers"
	"golang.org/x/tools/go/analysis"
)

var (
	Analyzer = &analysis.Analyzer{
		Name: "filesize",
		Doc:  "check for files that exceed 3000 lines",
		Run:  run,
	}
	ignoreFilesPattern string
)

func init() {
	Analyzer.Flags.StringVar(&ignoreFilesPattern, "ignore", "", "Comma separated list of files to ignore")
}

func run(pass *analysis.Pass) (interface{}, error) {
	var ignoreFiles []string
	if ignoreFilesPattern != "" {
		ignoreFiles = strings.Split(ignoreFilesPattern, ",")
	}

	for _, file := range pass.Files {
		if node := file; node != nil {
			f := pass.Fset.File(node.Pos())
			if helpers.IsFileIgnored(f.Name(), ignoreFiles) {
				continue
			}

			// Get the number of lines in the file
			lineCount := f.LineCount()
			if lineCount > 3000 {
				pass.Reportf(node.Pos(), "file has %d lines, which exceeds the maximum of 3000 lines", lineCount)
			}
		}
	}
	return nil, nil
}
