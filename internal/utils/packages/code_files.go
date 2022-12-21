// Package packages provides utility functions related to package manipulation
package packages

import "go/build"

// GetCodeFiles returns the list of code files in a certain package.
// The code files are: GoFiles, CgoFiles and TestGoFiles
func GetCodeFiles(pkg *build.Package) []string {
	return append(append(pkg.GoFiles, pkg.CgoFiles...), pkg.TestGoFiles...)
}
