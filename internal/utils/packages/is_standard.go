package packages

import "strings"

// IsStandardPackage verify if a certain package is a part of golang standard lib or not.
// this function returns true for standard lib packages (https://pkg.go.dev/std) and
// for Sub-repositories (https://pkg.go.dev/golang.org/x)
func IsStandardPackage(pkg string) bool {
	if strings.HasPrefix(pkg, "golang.org/x") {
		return true
	}
	if strings.ContainsAny(pkg, ".") {
		return false
	}
	return true
}
