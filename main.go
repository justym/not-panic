package main

import (
	"go/parser"
	"go/token"
	"log"
	"os"

	"github.com/justym/not-panic/pkg"
)

const src = "./_examples/case1.go"

func main() {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, src, nil, parser.Mode(0))
	check(err)
	pkg.InspectPanic(node, fset)
}

func check(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
