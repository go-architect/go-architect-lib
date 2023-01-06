// Package interfaces provides functions to retrieve information about interfaces in a Golang project
package interfaces

import (
	packageUtils "github.com/fdaines/go-architect-lib/internal/utils/packages"
	"github.com/fdaines/go-architect-lib/packages"
	"github.com/fdaines/go-architect-lib/project"
	"go/ast"
	"go/parser"
	"go/token"
	"math"
	"os"
	"path/filepath"
)

// ResolveInterfaceMetrics retrieves metrics about interfaces in the provided Golang project.
//
// An error is returned when it's not possible to get the packages information
func ResolveInterfaceMetrics(prj *project.ProjectInfo) (*InterfaceMetrics, error) {
	var interfaceInfos []InterfaceInfo
	pkgs, err := packages.GetBasicPackagesInfo(prj)
	if err != nil {
		return nil, err
	}

	for _, pkg := range pkgs {
		if pkg.PackageData != nil {
			for _, f := range packageUtils.GetCodeFiles(pkg.PackageData) {
				currentInterfaces, _ := countInterfaceMethods(pkg.PackageData.Dir, f)
				interfaceInfos = append(interfaceInfos, currentInterfaces...)
			}
		}
	}

	return &InterfaceMetrics{
		AverageMethods:      resolveAverageMethods(interfaceInfos),
		InterfaceMaxMethods: resolveInterfacesWithMaxMethods(interfaceInfos),
		InterfaceMinMethods: resolveInterfacesWithMinMethods(interfaceInfos),
	}, nil
}

func resolveInterfacesWithMaxMethods(interfaces []InterfaceInfo) []InterfaceInfo {
	max := 0
	var collection []InterfaceInfo
	if len(interfaces) == 0 {
		return collection
	}

	for _, i := range interfaces {
		if i.Methods > max {
			max = i.Methods
		}
	}

	for _, i := range interfaces {
		if i.Methods == max {
			collection = append(collection, i)
		}
	}

	return collection
}

func resolveInterfacesWithMinMethods(interfaces []InterfaceInfo) []InterfaceInfo {
	min := math.MaxInt
	var collection []InterfaceInfo
	if len(interfaces) == 0 {
		return collection
	}

	for _, i := range interfaces {
		if i.Methods < min {
			min = i.Methods
		}
	}

	for _, i := range interfaces {
		if i.Methods == min {
			collection = append(collection, i)
		}
	}

	return collection
}

func resolveAverageMethods(interfaces []InterfaceInfo) float64 {
	if len(interfaces) == 0 {
		return 0
	}

	var sum, total int
	for _, i := range interfaces {
		total++
		sum += i.Methods
	}

	return float64(sum) / float64(total)
}

func countInterfaceMethods(pkgPath string, srcFile string) ([]InterfaceInfo, error) {
	data, err := os.ReadFile(filepath.Join(pkgPath, srcFile))
	if err != nil {
		return nil, err
	}
	fileset := token.NewFileSet()
	astFile, err := parser.ParseFile(fileset, srcFile, data, 0)
	if err != nil {
		return nil, err
	}

	var interfaceInfos []InterfaceInfo
	for _, x := range astFile.Decls {
		if x, ok := x.(*ast.GenDecl); ok {
			if x.Tok != token.TYPE {
				continue
			}
			for _, x := range x.Specs {
				if x, ok := x.(*ast.TypeSpec); ok {
					interfaceName := x.Name.String()
					if x, ok := x.Type.(*ast.InterfaceType); ok {
						currentInterface := InterfaceInfo{
							Package: pkgPath,
							File:    srcFile,
							Name:    interfaceName,
							Methods: 0,
						}

						if x.Methods != nil {
							currentInterface.Methods = len(x.Methods.List)
						}
						interfaceInfos = append(interfaceInfos, currentInterface)
					}
				}
			}
		}
	}

	return interfaceInfos, nil
}
