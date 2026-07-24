package builder

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
)

func CreateStructure(basePath string) {
	directories := []string{
		filepath.Join(basePath, "src", "requirements", "wordpress"),
		filepath.Join(basePath, "src", "requirements", "mariadb"),
		filepath.Join(basePath, "src", "requirements", "nginx"),
	}

	for _, dir := range directories {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Error("Failed to create directory", "path", dir, "err", err)
			os.Exit(1)
		}
	}
	log.Info("Project directory structure created successfully!")
}

func GenerateConfigs(basePath string) {
}