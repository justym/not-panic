package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

const src = "./_examples/case1.go"

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, src, nil, parser.Mode(0))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	decls := f.Decls
	for i, decl := range decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			fmt.Printf("%d: line:%d function %+v\n", i, decl.Name.NamePos, decl.Name.Name)
			lists := decl.Body.List
			for _, item := range lists {
				switch stmt := item.(type) {
				case *ast.IfStmt:
					lbrace := stmt.Body.Lbrace
					rbrace := stmt.Body.Rbrace
					fmt.Printf("Between %d and %d: In Function %v, This is ast type IfStmt\n", lbrace, rbrace, decl.Name.Name)
					BodyList := stmt.Body.List
					for _, item := range BodyList {
						switch isExprStmt := item.(type) {
						case *ast.ExprStmt:
							switch expr := isExprStmt.X.(type) {
							case *ast.CallExpr:
								switch fun := expr.Fun.(type) {
								case *ast.Ident:
									if fun.Name == "panic" {
										position := fun.Pos()
										fmt.Printf("%d: Do not Use panic() !!\n", position)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func stdout(values ...interface{}) {
	if len(values) == 1 {
		fmt.Printf("(1): This is %+v\n", values...)
		return
	}
	for i, v := range values {
		fmt.Printf("(%d): This is %+v\n", i, v)
	}
}
