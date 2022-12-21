package project

// A ProjectInfo represents data about a Golang project
type ProjectInfo struct {
	Name                 string   `json:"name"`                  // name of the golang project
	Path                 string   `json:"path"`                  // path where the project is stored locally
	Package              string   `json:"package"`               // project's module package
	OrganizationPackages []string `json:"organization_packages"` // list of packages patterns (prefix) maintained by the organization
}

// NewProjectInfo creates a ProjectInfo structure based on received parameters
func NewProjectInfo(name, path, mainPackage string) *ProjectInfo {
	return &ProjectInfo{
		Name:    name,
		Path:    path,
		Package: mainPackage,
	}
}
