// Package dsm provides functions to create the Dependency Structure Matrix for a Golang project
package dsm

import (
	"github.com/fdaines/go-architect-lib/internal/utils/arrays"
	"github.com/fdaines/go-architect-lib/packages"
	"github.com/fdaines/go-architect-lib/project"
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
		//		fmt.Printf("Package: %s\n", pkg.Path)
		if pkg.PackageData != nil {
			for _, d := range pkg.PackageData.Imports {
				//				fmt.Printf("\t* Imports %s\n", d)
				dependencyMatrix.Packages = append(dependencyMatrix.Packages, d)
			}
		}
	}
	dependencyMatrix.Packages = arrays.RemoveDuplicatedStrings(dependencyMatrix.Packages)
	fillDSM(dependencyMatrix, pkgs)
	dependencyMatrix.Packages = sortDSM(*dependencyMatrix, []string{}, []string{})
	fillDSM(dependencyMatrix, pkgs)

	return dependencyMatrix, nil
}

func fillDSM(dependencyMatrix *DependencyStructureMatrix, pkgs []*packages.PackageInfo) {
	dependencyMatrix.Dependencies = make([][]int, len(dependencyMatrix.Packages))
	for i := 0; i < len(dependencyMatrix.Packages); i++ {
		dependencyMatrix.Dependencies[i] = make([]int, len(dependencyMatrix.Packages))
	}
	for _, pkg := range pkgs {
		index1 := arrays.IndexOf(dependencyMatrix.Packages, pkg.Path)
		if pkg.PackageData != nil {
			for _, d := range pkg.PackageData.Imports {
				index2 := arrays.IndexOf(dependencyMatrix.Packages, d)
				dependencyMatrix.Dependencies[index2][index1] += 1
			}
		}
	}
}
