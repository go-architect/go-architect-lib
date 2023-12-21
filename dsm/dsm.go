// Package dsm provides functions to create the Dependency Structure Matrix for a Golang project
package dsm

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-architect/go-architect-lib/internal/utils/arrays"
	packageUtils "github.com/go-architect/go-architect-lib/internal/utils/packages"
	"github.com/go-architect/go-architect-lib/packages"
	"github.com/go-architect/go-architect-lib/project"
)

// GetDependencyStructureMatrix calculates the Dependency Structure Matrix for a given Golang project.
//
// It returns an error if it's not possible to load de DSM.
func GetDependencyStructureMatrix(prj *project.ProjectInfo) (*DependencyStructureMatrix, error) {
	dependencyMatrix := &DependencyStructureMatrix{
		Module: prj.Package,
	}
	pkgs, err := packages.GetBasicPackagesInfo(prj)
	if err != nil {
		return nil, err
	}

	for _, pkg := range pkgs {
		dependencyMatrix.Packages = append(dependencyMatrix.Packages, pkg.Path)
		if pkg.PackageData != nil {
			for _, d := range pkg.PackageData.Imports {
				dependencyMatrix.Packages = append(dependencyMatrix.Packages, d)
			}
		}
	}
	dependencyMatrix.Packages = arrays.RemoveDuplicatedStrings(dependencyMatrix.Packages)
	fillDSM(dependencyMatrix, pkgs)
	dependencyMatrix.Packages = reArrangeDSM(*dependencyMatrix, prj)
	fillDSM(dependencyMatrix, pkgs)

	return dependencyMatrix, nil
}

func fillDSM(dependencyMatrix *DependencyStructureMatrix, pkgs []*packages.PackageInfo) {
	dependencyMatrix.Dependencies = make([][]int64, len(dependencyMatrix.Packages))
	for i := 0; i < len(dependencyMatrix.Packages); i++ {
		dependencyMatrix.Dependencies[i] = make([]int64, len(dependencyMatrix.Packages))
	}
	for _, pkg := range pkgs {
		currentPackageIndex := arrays.IndexOf(dependencyMatrix.Packages, pkg.Path)
		if pkg.PackageData != nil {
			for _, d := range pkg.PackageData.Imports {
				dependencyIndex := arrays.IndexOf(dependencyMatrix.Packages, d)
				for _, f := range packageUtils.GetCodeFiles(pkg.PackageData) {
					srcPath := filepath.Join(pkg.PackageData.Dir, f)
					deps, _ := calculateFileDependencies(srcPath, d)
					dependencyMatrix.Dependencies[dependencyIndex][currentPackageIndex] += deps
				}
			}
		}
	}
}

func calculateFileDependencies(srcFile, dependency string) (int64, error) {
	data, err := os.ReadFile(srcFile)
	if err != nil {
		return 0, err
	}
	fileset := token.NewFileSet()
	astFile, err := parser.ParseFile(fileset, srcFile, data, 0)
	if err != nil {
		return 0, err
	}

	total := countDependencies(fileset, astFile, dependency)
	return total, nil
}

func countDependencies(fileset *token.FileSet, astFile *ast.File, dep string) int64 {
	var dependencyPrefix string
	var counter int64
	ast.Inspect(astFile, func(n ast.Node) bool {
		if n == nil {
			return true
		}
		switch t := n.(type) {
		case *ast.ImportSpec:
			if sameDependency(t.Path.Value, dep) {
				if t.Name != nil && t.Name.Name != "." {
					dependencyPrefix = t.Name.Name
				} else {
					dependencyPrefix = retrievePrefix(dep)
				}
				counter++
			}
		case *ast.SelectorExpr:
			var buf bytes.Buffer
			printer.Fprint(&buf, fileset, t)
			if dependencyPrefix != "" && strings.HasPrefix(buf.String(), dependencyPrefix) {
				counter++
			}
		}
		return true
	})
	return counter
}

func sameDependency(d1, d2 string) bool {
	dx1 := strings.Replace(d1, "\"", "", 2)
	dx2 := strings.Replace(d2, "\"", "", 2)

	return dx1 == dx2
}

func retrievePrefix(dep string) string {
	split := strings.Split(dep, "/")
	return split[len(split)-1]
}
