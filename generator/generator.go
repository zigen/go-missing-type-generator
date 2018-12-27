package generator

import (
	"fmt"
	"github.com/zigen/go-missing-type-generator/utils"
	"go/ast"
	"go/format"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io"
	"io/ioutil"
	"path"
	"strings"
)

type Generator struct {
	BasePath    string
	Config      *types.Config
	FSet        *token.FileSet
	Files       []*ast.File
	Errors      []*error
	NeededTypes []*OneOfType
	Pkg         *types.Package
}

type OneOfType struct {
	Name  string
	Types []string
}

func (g *Generator) HandleTypeCheckError(err error) {
	if e, ok := err.(types.Error); ok {
		if strings.HasPrefix(e.Msg, "undeclared name") {
			ident := utils.TrimIdent(e.Msg)
			names := g.DecomposeIdent(ident)
			if !g.findNeededTypeByName(*ident) {
				g.NeededTypes = append(g.NeededTypes, &OneOfType{Name: *ident, Types: names})
			}
		} else {
			g.Errors = append(g.Errors, &err)
		}
	}
}

func (g *Generator) findNeededTypeByName(name string) bool {
	for _, t := range g.NeededTypes {
		if t.Name == name {
			return true
		}
	}
	return false
}

// decompose identifier from composed type name which starts from OneOf-, AnyOf-, AllOf- into array of string.
// OneOfObjAObjB -> ["ObjA", "ObjB"]
func (g *Generator) DecomposeIdent(composedName *string) []string {
	var (
		names = []string{}
		name  = strings.TrimPrefix(*composedName, "OneOf")
	)
	if name != *composedName {
		var (
			i = 0
			s = 0
		)
		for i <= len(name) {
			var typeName = name[s:i]
			if g.findDeclaredType(typeName) != nil {
				s = i
				names = append(names, typeName)
			} else {
				if i == len(name) {
					fmt.Printf("Error: cannot decompose composed oneof type: %s\n", composedName)
				}
			}

			i++
		}
	}
	return names
}

func (g *Generator) findDeclaredType(name string) *ast.TypeSpec {
	var foundType *ast.TypeSpec = nil
	for _, f := range g.Files {
		ast.Inspect(f, func(node ast.Node) bool {
			if n, ok := node.(*ast.GenDecl); ok {
				for _, s := range n.Specs {

					if t, ok := s.(*ast.TypeSpec); ok && t.Name.Name == name {
						//fmt.Printf("decl: %#v\n", t.Name)
						foundType = t
					}
				}
			}
			return true
		})
	}
	return foundType
}
func NewGenerator(basePath string) *Generator {
	g := &Generator{
		BasePath: basePath,
		FSet:     token.NewFileSet(),
		Files:    []*ast.File{},
		Errors:   []*error{},
	}

	g.Config = &types.Config{
		Importer: importer.Default(),
		Error:    g.HandleTypeCheckError,
	}
	return g
}

func collectSources(basePath string) ([]string, error) {
	files, err := ioutil.ReadDir(basePath)

	if err != nil {
		return nil, err
	}
	var ret = []string{}
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".go") {
			ret = append(ret, path.Join(basePath, f.Name()))
		}
	}
	return ret, nil
}

func (g *Generator) Parse() error {
	sources, err := collectSources(g.BasePath)
	if err != nil {
		return err
	}

	for _, src := range sources {
		f, err := parser.ParseFile(g.FSet, src, nil, parser.Mode(0))
		if err != nil {
			return fmt.Errorf("parse error: %s\n", err)
		}
		g.Files = append(g.Files, f)
	}
	return nil
}

func (g *Generator) Check() {
	pkg, err := g.Config.Check(g.BasePath, g.FSet, g.Files, nil)
	if err != nil {
		fmt.Printf("TypeCheck Error: %#v\n", err)
	}
	g.Pkg = pkg
	for _, t := range g.NeededTypes {
		fmt.Printf("types to generate %#v\n", t)
	}
}

func (g *Generator) GenerateNeededTypes(dst io.Writer) {
	f := &ast.File{
		Name:  ast.NewIdent(g.Pkg.Name()),
		Decls: []ast.Decl{},
	}
	for _, oneOfTypes := range g.NeededTypes {

		checkerFuncName := "Is" + oneOfTypes.Name
		decls := []*ast.FuncDecl{}

		f.Decls = append(f.Decls, utils.EmptyInterface(oneOfTypes.Name))
		for _, typeName := range oneOfTypes.Types {
			if foundDecls := g.findFuncDecls(typeName); foundDecls != nil {
				decls = append(decls, foundDecls...)
			}
			f.Decls = append(f.Decls, utils.FuncDecl(checkerFuncName, "*"+typeName))
		}

	}

	format.Node(dst, token.NewFileSet(), f)
}

func (g *Generator) findFuncDecls(rcvName string) []*ast.FuncDecl {
	funcDecls := []*ast.FuncDecl{}
	for _, f := range g.Files {
		ast.Inspect(f, func(node ast.Node) bool {
			if n, ok := node.(*ast.FuncDecl); ok {
				if n.Recv != nil {
					if t, ok := n.Recv.List[0].Type.(*ast.Ident); ok {
						if t.Name == rcvName {
							funcDecls = append(funcDecls, n)
						}
					}
				}
			}
			return true
		})
	}
	return funcDecls
}
