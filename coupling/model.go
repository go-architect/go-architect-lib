package coupling

type DependencyCoupling struct {
	Dependency     string             `json:"dependency"`     // the dependency evaluated for coupling level
	CouplingLevel  int                `json:"coupling_level"` // calculated coupling level for this dependency
	PackageDetails []*PackageCoupling `json:"details"`        // coupling level detailed by package
}

type PackageCoupling struct {
	Package       string          `json:"package"`        // the package evaluated
	CouplingLevel int             `json:"coupling_level"` // calculated coupling level for this package
	FileDetails   []*FileCoupling `json:"details"`        // coupling level detailed by package
}

type FileCoupling struct {
	Package       string   `json:"package"`        // the package
	File          string   `json:"file"`           // the file
	FilePath      string   `json:"file_path"`      // the file path
	Lines         []int    `json:"coupling_lines"` // the line where the dependency is used
	CouplingLevel int      `json:"coupling_level"` // calculated coupling level for this file
	Details       []Detail `json:"details"`        // details about the dependency
}

type Detail struct {
	Line     int    `json:"line"`    // the line where the dependency is used
	Comments string `json:"details"` // details about the use of the dependency in this line
}
