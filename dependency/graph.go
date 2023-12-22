// Package dependency provides functions to create the Dependency Graph for a Golang project
package dependency

import (
	"fmt"
	"github.com/go-architect/go-architect-lib/internal/utils/arrays"
	packageUtils "github.com/go-architect/go-architect-lib/internal/utils/packages"
	"github.com/go-architect/go-architect-lib/packages"
	"github.com/go-architect/go-architect-lib/project"
)

const allPackages = "ALL"

// GetDependencyGraph retrieves a structure that represents a Dependency Graph for the provided Golang project
// If you pass a startPackage different from "ALL", then this function will generate the Dependency Graph starting
// at this package and including only the referenced packages.
//
// An error is returned when it's not possible to get the packages information
func GetDependencyGraph(prj *project.ProjectInfo, startPackage string) (*ModuleDependencyGraph, error) {
	pkgs, err := packages.GetBasicPackagesInfo(prj)
	if err != nil {
		return nil, err
	}

	fmt.Printf("GetDependencyGraph: %+v\n", pkgs)
	//	fmt.Printf("StartPackage: %s\n", startPackage)
	if startPackage == allPackages {
		return getFullDependencyGraph(pkgs, prj.Package, prj.OrganizationPackages), nil
	}
	return getPartialDependencyGraph(startPackage, pkgs, prj.Package, prj.OrganizationPackages), nil
}

func getPartialDependencyGraph(startPackage string, pkgs []*packages.PackageInfo, mainPackage string, orgModulePatterns []string) *ModuleDependencyGraph {
	var currentPackage string
	var internal []string
	var external []string
	var organization []string
	var standard []string
	relations := make(map[string][]string)

	packagesQueue := []string{startPackage}

	for len(packagesQueue) > 0 {
		currentPackage, _ = arrays.Dequeue(packagesQueue)

		for _, pkg := range pkgs {
			if pkg.Path == currentPackage {
				internal = append(internal, pkg.Path)
				for _, i := range packageUtils.GetImportedPackages(pkg.PackageData) {
					flag := true
					relations[pkg.Path] = append(relations[pkg.Path], i)
					if packageUtils.IsInternalPackage(i, mainPackage) {
						packagesQueue = append(packagesQueue, i)
						flag = false
					}

					standard, organization, external = classifyDependency(flag, i, mainPackage, standard, orgModulePatterns, organization, external)
				}
			}
		}
	}

	return &ModuleDependencyGraph{
		Internal:     arrays.RemoveDuplicatedStrings(internal),
		External:     arrays.RemoveDuplicatedStrings(external),
		Organization: arrays.RemoveDuplicatedStrings(organization),
		StandardLib:  arrays.RemoveDuplicatedStrings(standard),
		Relations:    relations,
	}
}

func getFullDependencyGraph(pkgs []*packages.PackageInfo, mainPackage string, orgModulePatterns []string) *ModuleDependencyGraph {
	var internal []string
	var external []string
	var organization []string
	var standard []string

	relations := make(map[string][]string)

	for _, pkg := range pkgs {
		internal = append(internal, pkg.Path)
		fmt.Printf("Package: %+v\n", pkg.Path)
		if pkg.PackageData != nil {
			for _, i := range packageUtils.GetImportedPackages(pkg.PackageData) {
				flag := true
				relations[pkg.Path] = append(relations[pkg.Path], i)
				if packageUtils.IsInternalPackage(i, mainPackage) {
					flag = false
				}

				standard, organization, external = classifyDependency(flag, i, mainPackage, standard, orgModulePatterns, organization, external)
			}
		}
	}

	return &ModuleDependencyGraph{
		Internal:     arrays.RemoveDuplicatedStrings(internal),
		External:     arrays.RemoveDuplicatedStrings(external),
		Organization: arrays.RemoveDuplicatedStrings(organization),
		StandardLib:  arrays.RemoveDuplicatedStrings(standard),
		Relations:    relations,
	}
}

func classifyDependency(flag bool, dep string, mainPackage string, standard []string, orgModulePatterns []string, organization []string, external []string) ([]string, []string, []string) {
	if flag && packageUtils.IsStandardPackage(dep) {
		standard = append(standard, dep)
		flag = false
	}

	if flag && packageUtils.IsOrganizationPackage(dep, orgModulePatterns) {
		organization = append(organization, dep)
		flag = false
	}

	if flag && packageUtils.IsExternalPackage(dep, mainPackage) {
		external = append(external, dep)
	}
	return standard, organization, external
}
