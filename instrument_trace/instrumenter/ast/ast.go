package ast

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"

	"golang.org/x/tools/go/ast/astutil"
)

type instrumenter struct {
	traceImport string
	tracePkg    string
	traceFunc   string
}

func New(traceImport string, tracePkg string, traceFunc string) *instrumenter {
	return &instrumenter{
		traceImport: traceImport,
		tracePkg:    tracePkg,
		traceFunc:   traceFunc,
	}
}

func hasFunDec1(f *ast.File) bool {
	if len(f.Decls) == 0 {
		return false
	}
	for _, d := range f.Decls {
		_, ok := d.(*ast.FuncDecl)
		if ok {
			return true
		}
	}
	return false
}

func (a instrumenter) Instrument(filename string) ([]byte, error) {
	fset := token.NewFileSet()
	curAST, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("error parsing %s: %w", filename, err)
	}
	if !hasFunDec1(curAST) {
		return nil, nil
	}
	// add import declaration
	astutil.AddImport(fset, curAST, a.traceImport)

	// inject code into each function declaration
	a.addDeferTraceIntoFunDecls(curAST)

	buf := &bytes.Buffer{}
	err = format.Node(buf, fset, curAST)
	if err != nil {
		return nil, fmt.Errorf("error formatting new code:%w", err)
	}
	return buf.Bytes(), nil
}

func (a instrumenter) addDeferTraceIntoFunDecls(f *ast.File) {
	for _, d := range f.Decls {
		fd, ok := d.(*ast.FuncDecl)
		if ok {
			a.addDeferStmt(fd)
		}
	}
}

func (a instrumenter) addDeferStmt(fd *ast.FuncDecl) (added bool) {
	stmts := fd.Body.List

	// check whether "defer trace.Trace()()" has already exists
	for _, stmt := range stmts {
		ds, ok := stmt.(*ast.DeferStmt)
		if !ok {
			continue
		}
		// it is a defer stmt
		ce, ok := ds.Call.Fun.(*ast.CallExpr)
		if !ok {
			continue
		}
		se, ok := ce.Fun.(*ast.SelectorExpr)
		if !ok {
			continue
		}
		x, ok := se.X.(*ast.Ident)
		if !ok {
			continue
		}
		if (x.Name == a.tracePkg) && (se.Sel.Name == a.traceFunc) {
			return false
		}
	}
	// not found "defer trace.Trace()()"
	// add one
	ds := &ast.DeferStmt{
		Call: &ast.CallExpr{
			Fun: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X: &ast.Ident{
						Name: a.tracePkg,
					},
					Sel: &ast.Ident{
						Name: a.traceFunc,
					},
				},
			},
		},
	}
	newList := make([]ast.Stmt, len(stmts)+1)
	copy(newList[1:], stmts)
	newList[0] = ds
	fd.Body.List = newList
	return true
}
