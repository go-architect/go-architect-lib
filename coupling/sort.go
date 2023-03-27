package coupling

type SortPackagesByDependencyLevel []*PackageCoupling
type SortFilesByDependencyLevel []*FileCoupling

func (a SortPackagesByDependencyLevel) Len() int { return len(a) }
func (a SortPackagesByDependencyLevel) Less(i, j int) bool {
	return a[i].CouplingLevel > a[j].CouplingLevel
}
func (a SortPackagesByDependencyLevel) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a SortFilesByDependencyLevel) Len() int { return len(a) }
func (a SortFilesByDependencyLevel) Less(i, j int) bool {
	return a[i].CouplingLevel > a[j].CouplingLevel
}
func (a SortFilesByDependencyLevel) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
