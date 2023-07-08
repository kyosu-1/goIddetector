package goIddetector

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "goIddetector is a tool to detect identifiers containing 'Id'"

// Analyzer is an analysis analyzer that checks for identifiers containing 'Id'.
var Analyzer = &analysis.Analyzer{
	Name:     "goIddetector",
	Doc:      doc,
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.Ident)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.Ident:
			if n.Name == "Id" {
				pass.Reportf(n.Pos(), "Id should be an ID")
			}
		}
	})

	return nil, nil
}
