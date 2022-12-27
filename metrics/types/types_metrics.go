// Package types provides functions to count how many elements of each type are part of a Golang project
package types

import (
	packageUtils "github.com/fdaines/go-architect-lib/internal/utils/packages"
	"github.com/fdaines/go-architect-lib/packages"
	"github.com/fdaines/go-architect-lib/project"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

// ResolveProjectTypes returns a structure of type ProjectTypes which represents the number of each element type in
// a Golang project.
//
// An error is returned when it's not possible to get the packages information
func ResolveProjectTypes(prj *project.ProjectInfo) (*ProjectTypes, error) {
	var pkgCount, srcFileCount int
	var structCount, interfaceCount, functionCount, methodCount int
	var varCount, constCount int

	pkgs, err := packages.GetBasicPackagesInfo(prj)
	if err != nil {
		return nil, err
	}

	for _, pkg := range pkgs {
		pkgCount++
		if pkg.PackageData != nil {
			for _, f := range packageUtils.GetCodeFiles(pkg.PackageData) {
				srcFileCount++
				s, i, fn, m, v, c, err := countTypes(pkg.PackageData.Dir, f)
				if err == nil {
					structCount += s
					interfaceCount += i
					functionCount += fn
					methodCount += m
					varCount += v
					constCount += c
				}
			}
		}
	}

	return &ProjectTypes{
		Packages: pkgCount,
		counter: counter{
			SourceFiles: srcFileCount,
			Structs:     structCount,
			Interfaces:  interfaceCount,
			Functions:   functionCount,
			Methods:     methodCount,
			Variables:   varCount,
			Constants:   constCount,
		},
	}, nil
}

func countTypes(pkgPath string, srcFile string) (int, int, int, int, int, int, error) {
	data, err := os.ReadFile(filepath.Join(pkgPath, srcFile))
	if err != nil {
		return 0, 0, 0, 0, 0, 0, err
	}
	fileset := token.NewFileSet()
	node, err := parser.ParseFile(fileset, srcFile, data, 0)
	if err != nil {
		return 0, 0, 0, 0, 0, 0, err
	}
	var s, i, fn, m, v, c int
	ast.Inspect(node, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.FuncDecl:
			if t.Recv != nil {
				m++
			} else {
				fn++
			}
		case *ast.InterfaceType:
			if t.Methods != nil && len(t.Methods.List) > 0 {
				i++
			}
		case *ast.StructType:
			s++
		case *ast.GenDecl:
			if t.Tok == token.VAR {
				v++
			}
			if t.Tok == token.CONST {
				v++
			}
		}
		return true
	})

	return s, i, fn, m, v, c, nil
}
