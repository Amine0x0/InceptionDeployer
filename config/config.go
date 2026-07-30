package config

import (
	"path/filepath"
)

const DefaultProjectName string = "Inception"

type ProjectConfig struct {
	ProjectName       string
	ProjectPath       string
	StudentLogin      string
	DomainName        string
	MysqlDatabase     string
	MysqlUser         string
	MysqlPassword     string
	MysqlRootPassword string
	WPTitle           string
	WPAdminUser       string
	WPAdminPassword   string
	WPAdminEmail      string
	WPUser            string
	WPUserPassword    string
	WPUserEmail       string
}

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
