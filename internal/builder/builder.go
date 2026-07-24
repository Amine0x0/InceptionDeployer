package builder

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
)

func CreateStructure(basePath string) {
	directories := []string{
		filepath.Join(basePath, "secrets"),
		filepath.Join(basePath, "srcs"),
		filepath.Join(basePath, "srcs", "requirements", "nginx", "conf"),
		filepath.Join(basePath, "srcs", "requirements", "nginx", "tools"),
		filepath.Join(basePath, "srcs", "requirements", "wordpress", "conf"),
		filepath.Join(basePath, "srcs", "requirements", "wordpress", "tools"),
		filepath.Join(basePath, "srcs", "requirements", "mariadb", "conf"),
		filepath.Join(basePath, "srcs", "requirements", "mariadb", "tools"),
	}

	for _, dir := range directories {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Error("Failed to create directory", "path", dir, "err", err)
			os.Exit(1)
		}
	}

	files := []string{
		filepath.Join(basePath, "Makefile"),

		filepath.Join(basePath, "secrets", "credentials.txt"),
		filepath.Join(basePath, "secrets", "db_password.txt"),
		filepath.Join(basePath, "secrets", "db_root_password.txt"),

		filepath.Join(basePath, "srcs", "docker-compose.yml"),
		filepath.Join(basePath, "srcs", ".env"),

		filepath.Join(basePath, "srcs", "requirements", "nginx", "Dockerfile"),
		filepath.Join(basePath, "srcs", "requirements", "nginx", ".dockerignore"),
		filepath.Join(basePath, "srcs", "requirements", "nginx", "conf", "nginx.conf"),
		filepath.Join(basePath, "srcs", "requirements", "nginx", "tools", "entrypoint.sh"),

		filepath.Join(basePath, "srcs", "requirements", "wordpress", "Dockerfile"),
		filepath.Join(basePath, "srcs", "requirements", "wordpress", ".dockerignore"),
		filepath.Join(basePath, "srcs", "requirements", "wordpress", "conf", "www.conf"),
		filepath.Join(basePath, "srcs", "requirements", "wordpress", "tools", "entrypoint.sh"),

		filepath.Join(basePath, "srcs", "requirements", "mariadb", "Dockerfile"),
		filepath.Join(basePath, "srcs", "requirements", "mariadb", ".dockerignore"),
		filepath.Join(basePath, "srcs", "requirements", "mariadb", "conf", "50-server.cnf"),
		filepath.Join(basePath, "srcs", "requirements", "mariadb", "tools", "entrypoint.sh"),
	}

	for _, file := range files {
		f, err := os.Create(file)
		if err != nil {
			log.Error("Failed to create file", "path", file, "err", err)
			continue
		}
		f.Close()
	}

	log.Info("Mandatory Inception scaffolding created successfully!")
}

func GenerateConfigs(basePath string) {
}