package pkg

import (
	"fmt"
	"go/ast"
	"go/token"
)

func InspectPanic(node ast.Node, fset *token.FileSet) {
	ast.Inspect(
		node,
		func(n ast.Node) bool {
			funcIdent, ok := n.(*ast.Ident)
			if !ok {
				return true
			}

			if funcIdent.Name == "panic" {
				identPos := funcIdent.Pos()
				f := node.(*ast.File)
				fileName := f.Name.Name
				fmt.Printf("%s:%d Found %s function\n", fileName, fset.Position(identPos).Line, funcIdent.Name)
			}

			return true
		},
	)
}
