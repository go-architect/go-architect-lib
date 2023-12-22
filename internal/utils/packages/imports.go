// Package packages provides utility functions related to package manipulation
package packages

import "go/build"

// GetImportedPackages returns the list of imported packages, considering imports from source and test files.
// The result contains Imports and TestImports
func GetImportedPackages(pkg *build.Package) []string {
	return removeDuplicateStr(append(pkg.Imports, pkg.TestImports...))
}

func removeDuplicateStr(strSlice []string) []string {
	var list []string
	allKeys := make(map[string]bool)
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
