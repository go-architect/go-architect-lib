package dsm

import (
	"github.com/fdaines/go-architect-lib/internal/utils/packages"
	"github.com/fdaines/go-architect-lib/project"
	"strings"
)

func sortPackages(matrixPkgs []string, prj *project.ProjectInfo) []string {
	var pkgs []string

	for _, p := range matrixPkgs {
		pkgs = append(pkgs, p)
	}

	for i := 0; i < len(pkgs)-1; i++ {
		for j := i + 1; j < len(pkgs); j++ {
			if (packages.IsInternalPackage(pkgs[j], prj.Package) && packages.IsInternalPackage(pkgs[i], prj.Package)) ||
				(packages.IsOrganizationPackage(pkgs[j], prj.OrganizationPackages) && packages.IsOrganizationPackage(pkgs[i], prj.OrganizationPackages)) ||
				(packages.IsExternalPackage(pkgs[j], prj.Package) && packages.IsExternalPackage(pkgs[i], prj.Package)) ||
				(packages.IsStandardPackage(pkgs[j]) && packages.IsStandardPackage(pkgs[i])) {
				if strings.Compare(pkgs[i], pkgs[j]) > 0 {
					temp := pkgs[j]
					pkgs[j] = pkgs[i]
					pkgs[i] = temp
				}
			} else if packages.IsInternalPackage(pkgs[j], prj.Package) && !packages.IsInternalPackage(pkgs[i], prj.Package) {
				temp := pkgs[j]
				pkgs[j] = pkgs[i]
				pkgs[i] = temp
			} else if packages.IsOrganizationPackage(pkgs[j], prj.OrganizationPackages) &&
				(packages.IsExternalPackage(pkgs[i], prj.Package) || packages.IsStandardPackage(pkgs[i])) {
				temp := pkgs[j]
				pkgs[j] = pkgs[i]
				pkgs[i] = temp
			} else if packages.IsExternalPackage(pkgs[j], prj.Package) && packages.IsStandardPackage(pkgs[i]) {
				temp := pkgs[j]
				pkgs[j] = pkgs[i]
				pkgs[i] = temp
			}
		}
	}

	return pkgs
}
