package utils

import (
	"go/ast"
	"go/token"
	"strings"
)

func TrimIdent(errMsg string) *string {
	msgs := strings.Split(errMsg, ":")
	if len(msgs) == 2 {
		ident := strings.Trim(msgs[1], " ")
		return &ident
	}
	return nil
}

func FuncDecl(fnName string, rcvType string) *ast.FuncDecl {
	return &ast.FuncDecl{
		Name: ast.NewIdent(fnName),
		Recv: &ast.FieldList{
			List: []*ast.Field{
				&ast.Field{
					Names: []*ast.Ident{ast.NewIdent("_")},
					Type:  ast.NewIdent(rcvType),
				},
			},
		},
		Type: &ast.FuncType{},
		Body: &ast.BlockStmt{},
	}
}

func EmptyInterface(name string) *ast.GenDecl {
	return &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: ast.NewIdent(name),
				Type: &ast.InterfaceType{
					Methods: &ast.FieldList{},
				},
			},
		},
	}
}
