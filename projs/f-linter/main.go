package main

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/analysis/singlechecker"
	"golang.org/x/tools/go/ast/inspector"
)

func main() {
	singlechecker.Main(Analyzer)
}

var Analyzer = &analysis.Analyzer{
	Name:     "goprintffuncname",
	Doc:      "Checks that printf-like functions are named with `f` at the end.",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (any, error) {
	pass.ResultOf[inspect.Analyzer].(*inspector.Inspector).
		Preorder([]ast.Node{(*ast.FuncDecl)(nil)}, func(n ast.Node) {
			funcD, ok := n.(*ast.FuncDecl)
			if !ok {
				return
			}

			params := funcD.Type.Params.List
			if len(params) != 2 {
				return
			}

			if param1, ok := params[0].Type.(*ast.Ident); !ok || param1.Name != "string" {
				return
			}

			param2, ok := params[1].Type.(*ast.Ellipsis)
			if !ok { // must be ...interface{}
				return
			}

			switch t := param2.Elt.(type) {
			case *ast.InterfaceType:
				if t.Methods != nil && len(t.Methods.List) != 0 {
					return
				}
			case *ast.Ident:
				if t.Name != "any" {
					return
				}
			default:
				return
			}

			if strings.HasSuffix(funcD.Name.Name, "f") {
				return
			}

			pass.Reportf(n.Pos(),
				`printf-like formatting function %[1]q should be named "%[1]sf"`,
				funcD.Name.Name)
			return
		})
	return nil, nil
}
