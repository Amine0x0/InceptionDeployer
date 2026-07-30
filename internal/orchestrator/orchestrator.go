package orchestrator

import (
	"InceptionDeployer/config"
	makefilegen "InceptionDeployer/internal/Makefilegen"
	"InceptionDeployer/internal/composegen"
	"InceptionDeployer/internal/mariadbgen"
	"InceptionDeployer/internal/nginxgen"
	"InceptionDeployer/internal/wordpressgen"
)

func Generateall(ProjectPath string, cfg config.ProjectConfig) {
	makefilegen.Exec(ProjectPath, cfg)
	composegen.Exec(config.ScopeResolver(ProjectPath, "srcs"), cfg)
	nginxgen.Exec(config.ScopeResolver(ProjectPath, "srcs"), cfg)
	mariadbgen.Exec(config.ScopeResolver(ProjectPath, "srcs"), cfg)
	wordpressgen.Exec(config.ScopeResolver(ProjectPath, "srcs"), cfg)
}
