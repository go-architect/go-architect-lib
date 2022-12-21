package packages

import "go/build"

// PackageInfo contains basic info about a package and a
// reference to the build.Package to describe this Go package.
type PackageInfo struct {
	Name        string         `json:"name"` // package name
	Path        string         `json:"path"` // package path
	PackageData *build.Package // describes the Go package
}
