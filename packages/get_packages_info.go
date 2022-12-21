// Package packages provides functions to retrieve infromation about packages in a Golang project
package packages

import (
	"fmt"
	"github.com/fdaines/go-architect-lib/project"
	"go/build"
	"golang.org/x/tools/go/packages"
)

var verbose bool

// GetBasicPackagesInfo retrieves a slice of information about each package in a project.
//
// An error is returned when it's not possible to get the packages information
func GetBasicPackagesInfo(prj project.ProjectInfo) ([]*PackageInfo, error) {
	var packagesInfo []*PackageInfo
	var context = build.Default
	context.Dir = prj.Path

	pkgs, err := getPackages(prj)
	if err != nil {
		return nil, fmt.Errorf("Error: %v\n", err)
	} else {
		for index, packageName := range pkgs {
			if verbose {
				fmt.Printf("Loading packages (%d/%d): %s\n", index+1, len(pkgs), packageName)
			}
			pkg, errX := context.Import(packageName, "", 0)
			if errX != nil {
				return nil, errX
			}
			packagesInfo = append(packagesInfo, &PackageInfo{
				PackageData: pkg,
				Name:        pkg.Name,
				Path:        pkg.ImportPath,
			})
		}
	}

	return packagesInfo, nil
}

func getPackages(prj project.ProjectInfo) ([]string, error) {
	config := &packages.Config{Dir: prj.Path}
	pkgs, err := packages.Load(config, "./...")
	if err != nil {
		return nil, fmt.Errorf("Cannot load module packages: %+v", err)
	}
	var packages []string
	for _, pkg := range pkgs {
		packages = append(packages, pkg.PkgPath)
	}
	if verbose {
		fmt.Printf("%v packages found...\n", len(packages))
	}
	return packages, nil
}
