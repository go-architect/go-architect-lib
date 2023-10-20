# Go Architecture Library

[![CI Workflow](https://github.com/go-architect/go-architect-lib/actions/workflows/default.yml/badge.svg)](https://github.com/go-architect/go-architect-lib/actions/workflows/default.yml)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-architect/go-architect-lib)](https://pkg.go.dev/github.com/go-architect/go-architect-lib)
[![GoReport](https://goreportcard.com/badge/github.com/go-architect/go-architect-lib)](https://goreportcard.com/report/github.com/go-architect/go-architect-lib)

Go Architecture Library contains a set of functions and data structures
to support the architectural analysis for a Golang Project.


## Table of Contents

- [Getting Started](#getting-started)
  - [Import the library](#importing-the-library)
  - [Load Project Information](#load-project-information)
  - [Calculate Dependency Coupling Level](#calculate-dependency-coupling)
  - [Load Dependency Structure Matrix](#load-dependency-structure-matrix)
  - [Load Dependencies Graph](#load-dependencies-graph)
  - [Get Instability and Abstractness](#get-instability-and-abstractness)
  - [General Metrics](#general-metrics)
    - [Lines of Code](#lines-of-code)
    - [Count types](#count-types)
    - [Interfaces](#interfaces)
    - [Comments](#comments)
  - [Get Packages Information](#get-packages-information)
  - [Git Repository Information](#git-repository-information)
- [License](#license)
- [Credits](#credits)

## Getting Started

```bash
GO111MODULE=on go get -u github.com/go-architect/go-architect-lib
```

### Importing the library

```go
import "github.com/go-architect/go-architect-lib/project"
```

### Load Project Information

```go
package examples

import (
  "fmt"
  "github.com/go-architect/go-architect-lib/project"
)

func project_information_example() {
  projectPath := "full_path_to_golang_project"

  projectInfo, err := project.LoadProjectInfo(projectPath)
  if err != nil {
    fmt.Printf("Error: %+v\n", err)
  }

  fmt.Println("Project Information")
  fmt.Printf("* Project Name: %s\n", projectInfo.Name)
  fmt.Printf("* Project Path: %s\n", projectInfo.Path)
  fmt.Printf("* Project Main Package: %s\n", projectInfo.Package)
}
```

### Calculate Dependency Coupling Level

One of the most powerful features of this library is the option to analyze the coupling level for a
specific dependency, this information is very helpful to estimate the effort for migrate this dependency or for
encapsulate it to produce a more modular codebase.

```go
package examples

import (
	"fmt"
	"github.com/go-architect/go-architect-lib/coupling"
	"github.com/go-architect/go-architect-lib/project"
)

func dependency_coupling() {
	projectPath := "full_path_to_golang_project"

	projectInfo, err := project.LoadProjectInfo(projectPath)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	cl, err := coupling.CalculateCoupling(projectInfo, "fmt")
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	fmt.Println("-------------------------------------")
	fmt.Printf("Dependency: %s\n", cl.Dependency)
	fmt.Printf("CouplingLevel: %v\n", cl.CouplingLevel)
	for _, p := range cl.PackageDetails {
		fmt.Printf("\tPackage: %s\n", p.Package)
		fmt.Printf("\tCouplingLevel: %v\n", p.CouplingLevel)
		for _, f := range p.FileDetails {
			fmt.Printf("\t\tFile: %s\n", f.File)
			fmt.Printf("\t\tCouplingLevel: %v\n", f.CouplingLevel)
			for _, d := range f.Details {
				fmt.Printf("\t\t\tLine: %d \tDetails: %s\n", d.Line, d.Comments)
			}
		}
	}
}
```

### Load Dependency Structure Matrix

Go Architect Library can generate the [Dependency Structure Matrix](https://en.wikipedia.org/wiki/Design_structure_matrix) for a Golang project.

The packages are added to the DSM using the following criteria:
- First the packages are grouped and added based on each packages group: first Internal packages, then Organization packages, then External packages and finally the Standard packages.
- In each group, the packages are sorted alphabetically.
```go
package examples

import (
	"fmt"
	"github.com/go-architect/go-architect-lib/dsm"
	"github.com/go-architect/go-architect-lib/project"
)

func load_dsm() {
	projectPath := "full_path_to_golang_project"

	projectInfo, err := project.LoadProjectInfo(projectPath)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	matrix, err := dsm.GetDependencyStructureMatrix(projectInfo)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	fmt.Println("Dependency Structure Matrix")
	fmt.Printf("%+v\n", matrix)
}
```


### Load Dependencies Graph

Go Architect Library can generate the [Dependencies Graph](https://deepsource.io/glossary/dependency-graph/) for a Golang project.

```go
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
```

### Get Instability and Abstractness

Go Architect Library can generate a set of [Software Package Metrics](https://en.wikipedia.org/wiki/Software_package_metrics) for a Golang project.
These metrics are useful for analyzing the modular level for a set of software packages.

```go
package examples

import (
	"fmt"
	"github.com/go-architect/go-architect-lib/instability"
	"github.com/go-architect/go-architect-lib/project"
)

func get_instability_and_abstractness() {
	projectPath := "full_path_to_golang_project"

	projectInfo, err := project.LoadProjectInfo(projectPath)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	spm, err := instability.GetInstability(projectInfo)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	fmt.Println("Instability & Abstractness Analysis")
	fmt.Printf("%+v\n", spm)
}
```

### General Metrics

#### Lines of Code
Lines of Code is a basic metric for your codebase

```go
package examples

import (
  "fmt"
  "github.com/go-architect/go-architect-lib/metrics/loc"
  "github.com/go-architect/go-architect-lib/project"
)

func count_loc() {
  projectPath := "full_path_to_golang_project"

  projectInfo, err := project.LoadProjectInfo(projectPath)
  if err != nil {
    fmt.Printf("Error: %+v\n", err)
  }

  locInfo, err := loc.CountLoc(projectInfo)
  if err != nil {
    fmt.Printf("Error: %+v\n", err)
  }

  fmt.Println("Project Lines of Code")
  fmt.Printf("Total LOC: %d\n", locInfo.Total)

  fmt.Println("LOC by Package")
  for _, pkgLoc := range locInfo.Packages {
    fmt.Printf("\tPackage: %s \t\t %d\n", pkgLoc.Package, pkgLoc.LOC)
  }

  fmt.Println("LOC by File")
  for _, fileLoc := range locInfo.Files {
    fmt.Printf("\tPackage [%s] \t\tFile [%s]  \t\t %d\n", fileLoc.Package, fileLoc.File, fileLoc.LOC)
  }
}
```

#### Count types
This library offers a easy way to count the number of each type occurrences in a project.

```go
package examples

import (
	"fmt"
	"github.com/go-architect/go-architect-lib/metrics/types"
	"github.com/go-architect/go-architect-lib/project"
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
```

#### Interfaces
There are some interesting metrics about interfaces, like the average number of methods in your interfaces.
Also you can get the list of interfaces that declares the max and min number of methods in your project.

```go
package examples

import (
	"fmt"
	"github.com/go-architect/go-architect-lib/metrics/interfaces"
	"github.com/go-architect/go-architect-lib/project"
)

func interface_metrics() {
	projectPath := "full_path_to_golang_project"

	projectInfo, err := project.LoadProjectInfo(projectPath)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	interfaceMetrics, err := interfaces.ResolveInterfaceMetrics(projectInfo)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	fmt.Println("Interface Metrics")
	fmt.Printf("Average Number of methods: %s\n", interfaceMetrics.AverageMethods)

	fmt.Println("Interfaces with max number of methods")
	for _, i := range interfaceMetrics.InterfaceMaxMethods {
		fmt.Printf("\t[Package: %s][File: %s][InterfaceName:%s][Methods:%d]\n", i.Package, i.File, i.Name, i.Methods)
	}

	fmt.Println("Interfaces with min number of methods")
	for _, i := range interfaceMetrics.InterfaceMinMethods {
		fmt.Printf("\t[Package: %s][File: %s][InterfaceName:%s][Methods:%d]\n", i.Package, i.File, i.Name, i.Methods)
	}
}
```

#### Comments
You can also get some metrics about comments in your project.

```go
package examples

import (
	"fmt"
	"github.com/go-architect/go-architect-lib/metrics/comments"
	"github.com/go-architect/go-architect-lib/project"
)

func comments_metrics() {
	projectPath := "full_path_to_golang_project"

	projectInfo, err := project.LoadProjectInfo(projectPath)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	commentMetrics, err := comments.ResolveCommentsMetrics(projectInfo)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	fmt.Println("Comments Metrics")
	fmt.Printf("Total Lines of Comments: %v\n", commentMetrics.TotalLines)
	fmt.Printf("Total Files with Comments: %v\n", commentMetrics.FilesWithComments)

	fmt.Printf("Ratio Lines of Comments: %v\n", commentMetrics.CommentsRatio)
	fmt.Printf("Ratio Files with Comments: %v\n", commentMetrics.FilesWithCommentsRatio)
}
```

### Get Packages Information

With the `packages` package, you can get general information about all the packages in a Golang project.
Each element in the returned slice contains the package name, import path and a reference to the `built.Package` element so then you can use this
data for deeper analysis.

```go
package examples

import (
	"fmt"
	"github.com/go-architect/go-architect-lib/packages"
	"github.com/go-architect/go-architect-lib/project"
)

func get_packages_information() {
	projectPath := "full_path_to_golang_project"

	projectInfo, err := project.LoadProjectInfo(projectPath)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	pkgs, err := packages.GetBasicPackagesInfo(projectInfo)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	fmt.Println("General Packages Information")
	for idx, pkg := range pkgs {
		fmt.Printf("Packages[%d] -> %+v\n", idx, pkg)
	}
}
```

### Git Repository Information

With the `repository` package, you can get general information about the GIT repository
of a Golang project.
Even if Git repository information is not part of an architectural analysis of a project is useful to
detect: the revision that was analyzed and if the project has uncommitted changes at the moment of the analysis.

```go
package examples

import (
	"fmt"
	"github.com/go-architect/go-architect-lib/repository"
	"github.com/go-architect/go-architect-lib/repository"
)

func git_repository_information() {
	projectPath := "full_path_to_golang_project"

	repo, err := repository.LoadRepositoryInfo(projectPath)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	fmt.Println("GIT Repository Information")
	fmt.Printf("%+v\n", repo)
}
```


## License

MIT

## Credits

- [Francisco Daines](https://github.com/fdaines)
