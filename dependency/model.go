package dependency

// ModuleDependencyGraph represents a Dependency Graph for a Golang Module
type ModuleDependencyGraph struct {
	Module       string              `json:"module"`       // module package
	Internal     []string            `json:"internal"`     // Internal dependencies
	External     []string            `json:"external"`     // External dependencies
	Organization []string            `json:"organization"` // dependencies maintained by the same organization
	StandardLib  []string            `json:"standard"`     // dependencies from standard golang library
	Relations    map[string][]string `json:"relations"`    // each key correspond to an internal package and the value stores the imported packages for this package
}
