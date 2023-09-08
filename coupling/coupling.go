// Package coupling provides functions to evaluate the coupling level respect to a dependency in a Golang project
package coupling

import (
	packageUtils "github.com/fdaines/go-architect-lib/internal/utils/packages"
	"github.com/fdaines/go-architect-lib/packages"
	"github.com/fdaines/go-architect-lib/project"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
)

// CalculateCoupling retrieves information about the coupling level of a provided Golang project respect to a specific dependency
//
// An error is returned when it's not possible to get the packages information
func CalculateCoupling(prj *project.ProjectInfo, dep string) (*DependencyCoupling, error) {
	pkgs, err := packages.GetBasicPackagesInfo(prj)
	if err != nil {
		return nil, err
	}

	dependants := resolveDependantPackages(pkgs, dep)
	dc := calculateDependencyCouplingForPackages(dependants, dep)
	sort.Sort(SortPackagesByDependencyLevel(dc.PackageDetails))

	return dc, nil
}

func calculateDependencyCouplingForPackages(pkgs []*packages.PackageInfo, dependency string) *DependencyCoupling {
	dc := &DependencyCoupling{
		Dependency: dependency,
	}
	for _, p := range pkgs {
		pc := &PackageCoupling{
			Package: p.Path,
		}
		for _, f := range packageUtils.GetCodeFiles(p.PackageData) {
			fc, err := calculateCouplingForFile(p.PackageData.Dir, f, dependency)
			if err == nil && fc != nil {
				pc.FileDetails = append(pc.FileDetails, fc)
			}
		}
		pc.CouplingLevel = calculateCouplingLevelForFile(pc.FileDetails)
		sort.Sort(SortFilesByDependencyLevel(pc.FileDetails))

		dc.PackageDetails = append(dc.PackageDetails, pc)
	}
	dc.CouplingLevel = calculateCouplingLevelForModule(dc.PackageDetails)
	return dc
}

func calculateCouplingForFile(pkgPath, srcFile, dep string) (*FileCoupling, error) {
	data, err := os.ReadFile(filepath.Join(pkgPath, srcFile))
	if err != nil {
		return nil, err
	}
	fileset := token.NewFileSet()
	astFile, err := parser.ParseFile(fileset, srcFile, data, 0)
	if err != nil {
		return nil, err
	}

	if containsDependency(astFile, dep) {
		fc := &FileCoupling{
			Package:  pkgPath,
			File:     srcFile,
			FilePath: filepath.Join(pkgPath, srcFile),
		}
		fc.Details = calculateCouplingDetails(fileset, astFile, dep)
		fc.CouplingLevel = len(fc.Details)
		fc.Lines = resolveCouplingLines(fc.Details)
		return fc, nil
	}

	return nil, nil
}

func resolveCouplingLines(details []Detail) []int {
	var lines []int
	for _, d := range details {
		if !containLines(lines, d.Line) {
			lines = append(lines, d.Line)
		}
	}

	return lines
}

func containLines(array []int, value int) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

func calculateCouplingLevelForFile(fileCoupling []*FileCoupling) int {
	var total int
	for _, fc := range fileCoupling {
		total += fc.CouplingLevel
	}
	return total
}

func calculateCouplingLevelForModule(packageCoupling []*PackageCoupling) int {
	var total int
	for _, pc := range packageCoupling {
		total += pc.CouplingLevel
	}
	return total
}
