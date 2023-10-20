// Package types provides functions to count how many elements of each type are part of a Golang project
package types

import (
	packageUtils "github.com/go-architect/go-architect-lib/internal/utils/packages"
	"github.com/go-architect/go-architect-lib/packages"
	"github.com/go-architect/go-architect-lib/project"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"unicode"
)

// ResolveProjectTypes returns a structure of type ProjectTypes which represents the number of each element type in
// a Golang project.
//
// An error is returned when it's not possible to get the packages information
func ResolveProjectTypes(prj *project.ProjectInfo) (*ProjectTypes, error) {
	var pkgCount, srcFileCount, pkgSrcFileCount int
	var projectCounter counter
	var packageDetails []PackageTypes

	pkgs, err := packages.GetBasicPackagesInfo(prj)
	if err != nil {
		return nil, err
	}

	for _, pkg := range pkgs {
		pkgSrcFileCount = 0
		pkgCount++
		if pkg.PackageData != nil {
			currentPackageType := PackageTypes{
				Package: pkg.Path,
			}
			for _, f := range packageUtils.GetCodeFiles(pkg.PackageData) {
				srcFileCount++
				pkgSrcFileCount++
				ct, err := countTypes(pkg.PackageData.Dir, f)
				if err == nil {
					projectCounter.Add(ct)
					currentPackageType.Add(ct)
				}
			}
			currentPackageType.SourceFiles = pkgSrcFileCount
			packageDetails = append(packageDetails, currentPackageType)
		}
	}
	projectCounter.SourceFiles = srcFileCount

	return &ProjectTypes{
		ProjectPackage: prj.Package,
		Packages:       pkgCount,
		counter:        projectCounter,
		PackageDetails: packageDetails,
	}, nil
}

func countTypes(pkgPath string, srcFile string) (counter, error) {
	data, err := os.ReadFile(filepath.Join(pkgPath, srcFile))
	if err != nil {
		return counter{}, err
	}
	fileset := token.NewFileSet()
	node, err := parser.ParseFile(fileset, srcFile, data, 0)
	if err != nil {
		return counter{}, err
	}
	var s, i, fn, m, v, c int
	var pubS, pubI, pubFn, pubM, pubV, pubC int
	for _, td := range node.Decls {
		switch t := td.(type) {
		case *ast.FuncDecl:
			if t.Recv != nil {
				m++
				if isPublic(t.Name) {
					pubM++
				}
			} else {
				if isPublic(t.Name) {
					pubFn++
				}
				fn++
			}
		case *ast.GenDecl:
			switch ts := t.Specs[0].(type) {
			case *ast.TypeSpec:
				switch tsx := ts.Type.(type) {
				case *ast.StructType:
					s++
					if isPublic(ts.Name) {
						pubS++
					}
				case *ast.InterfaceType:
					if tsx.Methods != nil && len(tsx.Methods.List) > 0 {
						i++
						if isPublic(ts.Name) {
							pubI++
						}
					}
				}
			case *ast.ValueSpec:
				if t.Tok == token.VAR {
					v++
					if isPublic(ts.Names[0]) {
						pubV++
					}
				}
				if t.Tok == token.CONST {
					c++
					if isPublic(ts.Names[0]) {
						pubC++
					}
				}
			}
		}
	}

	return counter{
		Structs:          s,
		Interfaces:       i,
		Functions:        fn,
		Methods:          m,
		Variables:        v,
		Constants:        c,
		PublicStructs:    pubS,
		PublicInterfaces: pubI,
		PublicFunctions:  pubFn,
		PublicMethods:    pubM,
		PublicVariables:  pubV,
		PublicConstants:  pubC,
	}, nil
}

func isPublic(i *ast.Ident) bool {
	if i == nil {
		return false
	}
	runes := []rune(i.Name)
	return unicode.IsUpper(runes[0])
}
