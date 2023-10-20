// Package loc provides functions to calculate lines of code for a Golang project
package loc

import (
	"github.com/go-architect/go-architect-lib/internal/utils/loc"
	packageUtils "github.com/go-architect/go-architect-lib/internal/utils/packages"
	packages "github.com/go-architect/go-architect-lib/packages"
	"github.com/go-architect/go-architect-lib/project"
)

// CountLoc retrieves the Lines of Code metrics for a Golang project.
// LocMetric contains lines of code aggregated by: file, package and project.
//
// An error is returned when it's not possible to get the packages information
func CountLoc(prj *project.ProjectInfo) (*ProjectLOC, error) {
	pkgs, err := packages.GetBasicPackagesInfo(prj)
	if err != nil {
		return nil, err
	}

	locByFile := calculateLocByFile(pkgs)
	locByPackage := calculateLocByPackage(locByFile)
	locTotal := calculateLocTotal(locByPackage)

	return &ProjectLOC{
		Total:    locTotal,
		Packages: locByPackage,
		Files:    locByFile,
	}, nil
}

func calculateLocTotal(locByPackage []PackageLOC) int {
	var loc int

	for _, lp := range locByPackage {
		loc += lp.LOC
	}

	return loc
}

func calculateLocByPackage(files []FileLOC) []PackageLOC {
	aggregator := map[string]int{}
	var locByPackage []PackageLOC

	for _, lf := range files {
		_, ok := aggregator[lf.Package]
		if ok {
			aggregator[lf.Package] += lf.LOC
		} else {
			aggregator[lf.Package] = lf.LOC
		}
	}

	for key, value := range aggregator {
		locByPackage = append(locByPackage, PackageLOC{
			Package: key,
			LOC:     value,
		})
	}

	return locByPackage
}

func calculateLocByFile(packages []*packages.PackageInfo) []FileLOC {
	var locByFile []FileLOC

	for _, pkg := range packages {
		if pkg.PackageData != nil {
			for _, f := range packageUtils.GetCodeFiles(pkg.PackageData) {
				lines, err := loc.CountLinesOfCode(pkg.PackageData.Dir, f)
				if err == nil {
					locByFile = append(locByFile, FileLOC{
						Package: pkg.PackageData.ImportPath,
						File:    f,
						LOC:     lines,
					})
				}
			}
		}
	}

	return locByFile
}
