package dsm

import (
	"github.com/fdaines/go-architect-lib/internal/utils/arrays"
	"github.com/fdaines/go-architect-lib/internal/utils/packages"
	"github.com/fdaines/go-architect-lib/project"
	"sort"
)

func reArrangeDSM(dsm DependencyStructureMatrix, prj *project.ProjectInfo) []string {
	return sortDSM(dsm, prj, []string{}, []string{})
}

func sortDSM(dsm DependencyStructureMatrix, prj *project.ProjectInfo, head []string, tail []string) []string {
	if len(dsm.Dependencies) == 0 {
		return append(head, tail...)
	}
	if len(dsm.Dependencies) == 1 {
		return append(append(head, dsm.Packages...), tail...)
	}

	noDependencies, noDependants := resolveCandidatesColumns(dsm, prj)

	for _, n := range noDependants {
		head = append(head, n.packageName)
	}
	for _, n := range noDependencies {
		tail = append([]string{n.packageName}, tail...)
	}
	newMatrix := removeRowsAndColumns(dsm, head, tail)

	return sortDSM(newMatrix, prj, head, tail)
}

func resolveCandidatesColumns(dsm DependencyStructureMatrix, prj *project.ProjectInfo) ([]dependencyDetails, []dependencyDetails) {
	details := make(map[string]dependencyDetails)
	for idx, c := range dsm.Packages {
		var dependencies, dependants int
		for i := 0; i < len(dsm.Packages); i++ {
			dependencies += dsm.Dependencies[i][idx]
			dependants += dsm.Dependencies[idx][i]
		}
		details[c] = dependencyDetails{
			packageName:  c,
			packageRank:  resolvePackageRank(c, prj),
			dependencies: dependencies,
			dependants:   dependants,
		}
	}

	var noDependencies []dependencyDetails
	var noDependants []dependencyDetails
	for _, v := range details {
		if v.dependencies == 0 {
			noDependencies = append(noDependencies, v)
		}
		if v.dependants == 0 {
			noDependants = append(noDependants, v)
		}
	}
	sort.Slice(noDependencies, func(i, j int) bool {
		return (noDependencies[i].packageRank * noDependencies[i].dependants) > (noDependencies[j].packageRank * noDependencies[j].dependants)
	})
	sort.Slice(noDependants, func(i, j int) bool {
		return (noDependants[i].packageRank * noDependants[i].dependencies) > (noDependants[j].packageRank * noDependants[j].dependencies)
	})

	return noDependencies, noDependants
}

func removeRowsAndColumns(dsm DependencyStructureMatrix, head []string, tail []string) DependencyStructureMatrix {
	var newColumns []string
	var matrix [][]int

	for idx, c := range dsm.Packages {
		if !arrays.Contains(head, c) && !arrays.Contains(tail, c) {
			newColumns = append(newColumns, c)
			matrix = append(matrix, dsm.Dependencies[idx])
		}
	}

	for i := 0; i < len(matrix); i++ {
		var newRow []int
		for idx, c := range dsm.Packages {
			if !arrays.Contains(head, c) && !arrays.Contains(tail, c) {
				newRow = append(newRow, matrix[i][idx])
			}
		}
		matrix[i] = newRow
	}

	return DependencyStructureMatrix{
		Module:       dsm.Module,
		Packages:     newColumns,
		Dependencies: matrix,
	}
}

func resolvePackageRank(packageName string, prj *project.ProjectInfo) int {
	if packages.IsInternalPackage(packageName, prj.Package) {
		return 1
	}
	if packages.IsOrganizationPackage(packageName, prj.OrganizationPackages) {
		return 10
	}
	if packages.IsExternalPackage(packageName, prj.Package) {
		return 100
	}
	if packages.IsStandardPackage(packageName) {
		return 1000
	}
	return 10000
}
