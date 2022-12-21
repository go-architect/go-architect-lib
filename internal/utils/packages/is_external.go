package packages

import "strings"

// IsExternalPackage verify if a certain package is not part of: the module package, standard lib or sub-repositories
func IsExternalPackage(pkg, mainPackage string) bool {
	if strings.HasPrefix(pkg, mainPackage) {
		return false
	}
	if strings.HasPrefix(pkg, "golang.org/x") {
		return false
	}
	if strings.ContainsAny(pkg, ".") {
		return true
	}
	return false
}
