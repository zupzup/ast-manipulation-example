package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
)

func main() {
	// parse file
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "test.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	comments := []*ast.CommentGroup{}
	ast.Inspect(node, func(n ast.Node) bool {
		// collect comments
		c, ok := n.(*ast.CommentGroup)
		if ok {
			comments = append(comments, c)
		}
		// handle function declarations without documentation
		fn, ok := n.(*ast.FuncDecl)
		if ok {
			if fn.Name.IsExported() && fn.Doc.Text() == "" {
				// print warning
				fmt.Printf("exported function declaration without documentation found on line %d: \n\t%s\n", fset.Position(fn.Pos()).Line, fn.Name.Name)
				// create todo-comment
				comment := &ast.Comment{
					Text:  "// TODO: document exported function",
					Slash: fn.Pos() - 1,
				}
				// create CommentGroup and set it to the function's documentation comment
				cg := &ast.CommentGroup{
					List: []*ast.Comment{comment},
				}
				fn.Doc = cg
				fmt.Println()
			}
		}
		return true
	})
	// set ast's comments to the collected comments
	node.Comments = comments
	// write new ast to file
	f, err := os.Create("new.go")
	defer f.Close()
	if err := printer.Fprint(f, fset, node); err != nil {
		log.Fatal(err)
	}
}
