package examples

import (
	"fmt"
	"github.com/fdaines/go-architect-lib/metrics/interfaces"
	"github.com/fdaines/go-architect-lib/project"
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
