package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"strings"
)

func (c *Config) isSafe(code string) bool {

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", code, parser.ImportsOnly)
	if err != nil {
		fmt.Println(err)
		return false
	}

	for _, imp := range node.Imports {
		importPath := strings.Trim(imp.Path.Value, "\"")
		for _, banned := range c.Module {
			if strings.HasPrefix(importPath, banned) {
				return false
			}
		}
	}

	return true
}

func (c *Config) InfiniteLoop(code string) (bool, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", code, parser.AllErrors)
	if err != nil {
		return true, err
	}

	var hasInfiniteLoop bool = false
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.ForStmt:
			if x.Cond == nil {
				hasInfiniteLoop = true
				return false
			}
			if binExpr, ok := x.Cond.(*ast.BinaryExpr); ok && binExpr.Op == token.LEQ {
				if ident, ok := binExpr.Y.(*ast.BasicLit); ok && ident.Kind == token.INT {
					if val, err := strconv.Atoi(ident.Value); err == nil && val > c.MaxInt {
						hasInfiniteLoop = true
						return false
					}
				}
			}
		case *ast.RangeStmt:
			if x.Key != nil || x.Value != nil {
				if lit, ok := x.X.(*ast.BasicLit); ok && lit.Kind == token.INT {
					if val, err := strconv.Atoi(lit.Value); err == nil && val > c.MaxInt {
						hasInfiniteLoop = true
						return false
					}
				}
			}
		case *ast.SelectStmt:
			hasInfiniteLoop = true
			return false
		case *ast.SwitchStmt:
			hasInfiniteLoop = true
			return false
		case *ast.BranchStmt:
			if x.Tok == token.GOTO {
				hasInfiniteLoop = true
				return false
			}
		}
		return true
	})

	return hasInfiniteLoop, nil
}
