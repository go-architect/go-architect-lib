package loc

// ProjectLOC represents the total lines of code for a project
type ProjectLOC struct {
	Total    int          `json:"total"`    // number of lines of code for the whole project
	Packages []PackageLOC `json:"packages"` // detailed information by package
	Files    []FileLOC    `json:"files"`    // detailed information by file
}

// PackageLOC represents the total lines of code for a specific package
type PackageLOC struct {
	Package string `json:"package"` // package name
	LOC     int    `json:"loc"`     // number of lines of code for this package
}

// FileLOC represents the total lines of code for a specific code file
type FileLOC struct {
	Package string `json:"package"` // package name
	File    string `json:"file"`    // file name
	LOC     int    `json:"loc"`     // number of lines of code for this file
}
