package packages

import "strings"

// IsInternalPackage verify if a certain package is a part of the golang module
// represented by `mainPackage`
func IsInternalPackage(pkg, mainPackage string) bool {
	return strings.HasPrefix(pkg, mainPackage)
}
