package examples

import (
	"fmt"
	"github.com/fdaines/go-architect-lib/metrics/types"
	"github.com/fdaines/go-architect-lib/project"
)

func count_types() {
	projectPath := "full_path_to_golang_project"

	projectInfo, err := project.LoadProjectInfo(projectPath)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	projectTypes, err := types.ResolveProjectTypes(projectInfo)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	fmt.Println("Project Types")
	fmt.Printf("Main Package: %s\n", projectTypes.ProjectPackage)
	fmt.Printf("\t# of Packages: %d\n", projectTypes.Packages)
	fmt.Printf("\t# of SourceFiles: %d\n", projectTypes.SourceFiles)
	fmt.Printf("\t# of Structs: %d\n", projectTypes.Structs)
	fmt.Printf("\t# of Interfaces: %d\n", projectTypes.Interfaces)
	fmt.Printf("\t# of Functions: %d\n", projectTypes.Functions)
	fmt.Printf("\t# of Methods: %d\n", projectTypes.Methods)
	fmt.Printf("\t# of Variables: %d\n", projectTypes.Variables)
	fmt.Printf("\t# of Constants: %d\n", projectTypes.Constants)
	fmt.Printf("\t# of PublicStructs: %d\n", projectTypes.PublicStructs)
	fmt.Printf("\t# of PublicInterfaces: %d\n", projectTypes.PublicInterfaces)
	fmt.Printf("\t# of PublicFunctions: %d\n", projectTypes.PublicFunctions)
	fmt.Printf("\t# of PublicMethods: %d\n", projectTypes.PublicMethods)
	fmt.Printf("\t# of PublicVariables: %d\n", projectTypes.PublicVariables)
	fmt.Printf("\t# of PublicConstants: %d\n", projectTypes.PublicConstants)

	fmt.Println("Details by Package")
	for _, pd := range projectTypes.PackageDetails {
		// The same type counters are available at package level.
		fmt.Printf("\tPackage: %s \t\t SourceFiles: %d\n", pd.Package, pd.SourceFiles)
	}
}
