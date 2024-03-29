// Package instability provides functions to retrieve instability, abstractness and distance from main diagonal of each package in a Golang project
package instability

import (
	"github.com/go-architect/go-architect-lib/internal/utils"
	"github.com/go-architect/go-architect-lib/internal/utils/arrays"
	packageUtils "github.com/go-architect/go-architect-lib/internal/utils/packages"
	"github.com/go-architect/go-architect-lib/packages"
	"github.com/go-architect/go-architect-lib/project"
	"go/ast"
	"go/parser"
	"go/token"
	"math"
	"os"
	"path/filepath"
)

// GetInstability retrieves the PackageInstability data for each package in the provided project
//
// An error is returned when it's not possible to get the packages information
func GetInstability(prj *project.ProjectInfo) ([]*PackageInstability, error) {
	var data []*PackageInstability
	pkgs, err := packages.GetBasicPackagesInfo(prj)
	if err != nil {
		return nil, err
	}

	efferentCoupling, afferentCoupling := calculateCoupling(prj, pkgs)

	for _, pkg := range pkgs {
		ac := afferentCoupling[pkg.Path]
		ec := efferentCoupling[pkg.Path]
		info := &PackageInstability{
			PackageName:      pkg.Path,
			EfferentCoupling: len(arrays.RemoveDuplicatedStrings(ec)),
			AfferentCoupling: len(arrays.RemoveDuplicatedStrings(ac)),
		}
		calculateAbstractionsAndImplementations(info, pkg)
		calculateInstability(info)
		calculateAbstractness(info)
		calculateDistanceFromDiagonal(info)

		data = append(data, info)
	}

	return data, nil
}

func calculateAbstractionsAndImplementations(info *PackageInstability, pkg *packages.PackageInfo) {
	info.AbstractionsCount = 0
	info.ImplementationsCount = 0
	if pkg.PackageData != nil {
		for _, f := range packageUtils.GetCodeFiles(pkg.PackageData) {
			i, s, err := countTypes(pkg.PackageData.Dir, f)
			if err == nil {
				info.AbstractionsCount += i
				info.ImplementationsCount += s
			}
		}
	}
}

func calculateDistanceFromDiagonal(info *PackageInstability) {
	info.Distance = utils.RoundFloat(math.Abs(info.Instability+info.Abstractness-1), 2)
}

func calculateAbstractness(info *PackageInstability) {
	if info.ImplementationsCount+info.AbstractionsCount > 0 {
		info.Abstractness = utils.RoundFloat(float64(info.AbstractionsCount)/float64(info.ImplementationsCount+info.AbstractionsCount), 2)
	}
}

func calculateInstability(info *PackageInstability) {
	if info.AfferentCoupling+info.EfferentCoupling > 0 {
		info.Instability = utils.RoundFloat(float64(info.EfferentCoupling)/float64(info.AfferentCoupling+info.EfferentCoupling), 2)
	}
}

func countTypes(pkgPath string, srcFile string) (int, int, error) {
	data, err := os.ReadFile(filepath.Join(pkgPath, srcFile))
	if err != nil {
		return 0, 0, err
	}
	fileset := token.NewFileSet()
	node, err := parser.ParseFile(fileset, srcFile, data, 0)
	if err != nil {
		return 0, 0, err
	}
	var structs, interfaces, functions, methods int
	for _, td := range node.Decls {
		switch t := td.(type) {
		case *ast.FuncDecl:
			if t.Recv != nil {
				methods++
			} else {
				functions++
			}
		case *ast.GenDecl:
			switch ts := t.Specs[0].(type) {
			case *ast.TypeSpec:
				switch tsx := ts.Type.(type) {
				case *ast.StructType:
					structs++
				case *ast.InterfaceType:
					if tsx.Methods != nil && len(tsx.Methods.List) > 0 {
						interfaces++
					}
				}
			}
		}
	}

	return structs + interfaces, functions + methods, nil
}

func calculateCoupling(prj *project.ProjectInfo, pkgs []*packages.PackageInfo) (map[string][]string, map[string][]string) {
	efferentCoupling := make(map[string][]string)
	afferentCoupling := make(map[string][]string)

	for _, pkg := range pkgs {
		if pkg.PackageData != nil {
			for _, d := range packageUtils.GetImportedPackages(pkg.PackageData) {
				if packageUtils.IsInternalPackage(d, prj.Package) {
					efferentCoupling[pkg.Path] = append(efferentCoupling[pkg.Path], d)
					afferentCoupling[d] = append(afferentCoupling[d], pkg.Path)
				}
			}
		}
	}
	return efferentCoupling, afferentCoupling
}
