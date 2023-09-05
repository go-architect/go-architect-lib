package dsm

// DependencyStructureMatrix represents a DSM
type DependencyStructureMatrix struct {
	Module       string   `json:"module"`       // module package
	Packages     []string `json:"packages"`     // list of packages used by the module
	Dependencies [][]int  `json:"dependencies"` // this matrix[i][j] represents an import of Package[i] in Package[j]
}

// For internal use only, represents how many dependencies and dependants a package has.
type dependencyDetails struct {
	packageName  string
	packageRank  int
	dependencies int
	dependants   int
}
