package examples

import (
	"fmt"
	"github.com/go-architect/go-architect-lib/dependency"
	"github.com/go-architect/go-architect-lib/project"
)

func load_dependencies_graph() {
	projectPath := "full_path_to_golang_project"

	projectInfo, err := project.LoadProjectInfo(projectPath)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	// startPackage="ALL" will generate the dependencies graph for the whole project
	graph, err := dependency.GetDependencyGraph(projectInfo, "ALL")
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	// Te returned structure contains:
	// * The main package name for the project
	// * a list of internal, external and standardlib packages used in the project
	// * optionally, a list of packages that are external but are maintained by the same organization
	// * a map that represents a dependency relation ([key] -> [value]) where key is the dependant package and value is a list of dependencies used by this package.
	fmt.Println("Dependencies Graph")
	fmt.Printf("%+v\n", graph)
}

func load_dependencies_graph_with_organization_packages() {
	projectPath := "full_path_to_golang_project"

	projectInfo, err := project.LoadProjectInfo(projectPath)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}
	// this assign implies that every dependency that starts with `github.com/fdaines` should be considered
	// as a dependency maintained by the same organization (it means an external dependency but with a higher
	// level of maintainability than a 3rd party dependency)
	projectInfo.OrganizationPackages = []string{"github.com/fdaines"}

	// startPackage="ALL" will generate the dependencies graph for the whole project
	graph, err := dependency.GetDependencyGraph(projectInfo, "ALL")
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	// Te returned structure contains:
	// * The main package name for the project
	// * a list of internal, external and standardlib packages used in the project
	// * a list of packages maintained by the same organization (starting with `github.com/fdaines`)
	// * a map that represents a dependency relation ([key] -> [value]) where key is the dependant package and value is a list of dependencies used by this package.
	fmt.Println("Dependencies Graph")
	fmt.Printf("%+v\n", graph)
}
