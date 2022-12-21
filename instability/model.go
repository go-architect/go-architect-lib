package instability

// PackageInstability represents software package metrics for a specific package
// more info: https://en.wikipedia.org/wiki/Software_package_metrics
type PackageInstability struct {
	PackageName          string  `json:"package_name"`          // the package name
	AbstractionsCount    int     `json:"abstractions_count"`    // number of abstractions in the package
	ImplementationsCount int     `json:"implementations_count"` // number of implementations in the package
	AfferentCoupling     int     `json:"afferent_coupling"`     // the number of elements in other packages that depend upon elements within this package is an indicator of the package's responsibility.
	EfferentCoupling     int     `json:"efferent_coupling"`     // the number of elements in other packages that the elements in this package depend upon is an indicator of the package's dependence on externalities
	Instability          float64 `json:"instability"`           // calculated instability ( I = Ce / (Ce+Ca) )
	Abstractness         float64 `json:"abstractness"`          // calculated abstractness ( A = Abstractions/Total )
	Distance             float64 `json:"distance"`              // distance from main diagonal ( D = Abs[I + A - 1] )
}
