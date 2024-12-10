// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package toSql

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "toSql",
	Doc:  "check for usage of deprecated ToSql method calls",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		var currentFunction string
		ast.Inspect(file, func(node ast.Node) bool {
			switch x := node.(type) {
			case *ast.FuncDecl:
				currentFunction = x.Name.Name
			case *ast.CallExpr:
				if fun, ok := x.Fun.(*ast.SelectorExpr); ok {
					if fun.Sel.Name == "ToSql" {
						pass.Reportf(node.Pos(), "Function %q uses ToSql method which is deprecated and should be migrated", currentFunction)
					}
				}
			}
			return true
		})
	}
	return nil, nil
}
