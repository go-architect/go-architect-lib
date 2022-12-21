package types

// ProjectTypes represents the total number of each type of element in a project
type ProjectTypes struct {
	counter
	ProjectPackage string         `json:"project_package"` // project's module package
	Packages       int            `json:"packages"`        // number of packages in the module
	PackageDetails []PackageTypes `json:"details"`         // detailed information for each package
}

// PackageTypes represents the total number of each type of element in a package
type PackageTypes struct {
	counter
	Package string `json:"package"` // package name
}

// counter is used to count the number of elements by type (structs, interfaces, files, variables and so on).
type counter struct {
	SourceFiles      int `json:"source_files"`      // number of source files
	Structs          int `json:"structs"`           // number of structs
	Interfaces       int `json:"interfaces"`        // number of interfaces
	Functions        int `json:"functions"`         // number of functions
	Methods          int `json:"methods"`           // number of methods
	Variables        int `json:"variables"`         // number of variables
	Constants        int `json:"constants"`         // number of constants
	PublicStructs    int `json:"public_structs"`    // number of public structs
	PublicInterfaces int `json:"public_interfaces"` // number of public interfaces
	PublicFunctions  int `json:"public_functions"`  // number of public functions
	PublicMethods    int `json:"public_methods"`    // number of public methods
	PublicVariables  int `json:"public_variables"`  // number of public variables
	PublicConstants  int `json:"public_constants"`  // number of public constants
}
