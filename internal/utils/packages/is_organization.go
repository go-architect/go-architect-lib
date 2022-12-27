package packages

import "strings"

// IsOrganizationPackage checks if a package is maintained for this organization.
// orgPatterns is a list of packages prefixes, so then this function checks if the provided package
// starts with any of this package patterns.
func IsOrganizationPackage(pkg string, orgPatterns []string) bool {
	for _, op := range orgPatterns {
		if strings.HasPrefix(pkg, op) {
			return true
		}
	}

	return false
}
