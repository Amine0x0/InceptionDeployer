package config

import(
	"path/filepath"
)

const DefaultProjectName string = "Inception"

func ScopeResolver(basePath string, scope string) string {
	switch scope {
	case "root":
		return basePath
	case "secrets":
		return filepath.Join(basePath, "secrets")
	case "srcs":
		return filepath.Join(basePath, "srcs")
	case "reqs":
		return filepath.Join(basePath, "srcs", "requirements")
	default:
		return basePath
	}
}