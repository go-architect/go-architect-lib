package interfaces

// InterfaceMetrics represents general metrics about interfaces in a Golang project
type InterfaceMetrics struct {
	AverageMethods      float64         `json:"average_methods"` // average number of methods declared in all the interfaces
	InterfaceMaxMethods []InterfaceInfo `json:"max_methods"`     // interfaces which declares the max number of methods
	InterfaceMinMethods []InterfaceInfo `json:"min_methods"`     // interfaces which declares the min number of methods
}

// InterfaceInfo represents general information about a specific interface
type InterfaceInfo struct {
	Package string `json:"package"` // package name where the interface is declared
	File    string `json:"file"`    // file name where the interface is declared
	Name    string `json:"name"`    // interface name
	Methods int    `json:"methods"` // number of method in this interface
}
