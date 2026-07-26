package orchestrator

import (
	"InceptionDeployer/internal/composegen"
	"InceptionDeployer/config"
)

func Generateall(ProjectPath string) {
	composegen.Exec(config.ScopeResolver(ProjectPath, "srcs"))
}