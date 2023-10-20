package coupling

import (
	"github.com/go-architect/go-architect-lib/internal/utils/arrays"
	"github.com/go-architect/go-architect-lib/packages"
	"go/ast"
	"strings"
)

func retrievePrefix(dep string) string {
	split := strings.Split(dep, "/")
	return split[len(split)-1]
}

func containsDependency(file *ast.File, dep string) bool {
	for _, is := range file.Imports {
		if sameDependency(is.Path.Value, dep) {
			return true
		}
	}

	return false
}

func sameDependency(d1, d2 string) bool {
	dx1 := strings.Replace(d1, "\"", "", 2)
	dx2 := strings.Replace(d2, "\"", "", 2)

	return dx1 == dx2
}

func resolveDependantPackages(pkgs []*packages.PackageInfo, dep string) []*packages.PackageInfo {
	var dependants []*packages.PackageInfo
	for _, p := range pkgs {
		if p.PackageData != nil {
			if arrays.Contains(p.PackageData.Imports, dep) {
				dependants = append(dependants, p)
			}
		}
	}
	return dependants
}
