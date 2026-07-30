package builder

import (
	"os"
	"path/filepath"

	"InceptionDeployer/config"
	"InceptionDeployer/internal/filewriter"

	"github.com/charmbracelet/log"
)

func CreateStructure(basePath string) {
	directories := []string{
		filepath.Join(basePath, "secrets"),
		filepath.Join(basePath, "srcs"),
		filepath.Join(basePath, "srcs", "requirements", "nginx", "conf"),
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

	log.Info("Mandatory Inception scaffolding created successfully!")
}

func GenerateConfigs(basePath string, cfg config.ProjectConfig) {
	files := map[string]string{
		filepath.Join(basePath, "README.md"):                       readmeContent(cfg),
		filepath.Join(basePath, "USER_DOC.md"):                     userDocContent(cfg),
		filepath.Join(basePath, "DEV_DOC.md"):                      devDocContent(cfg),
		filepath.Join(basePath, "srcs", ".env"):                    envContent(cfg),
		filepath.Join(basePath, "secrets", "credentials.txt"):      cfg.WPAdminPassword,
		filepath.Join(basePath, "secrets", "db_password.txt"):      cfg.MysqlPassword,
		filepath.Join(basePath, "secrets", "db_root_password.txt"): cfg.MysqlRootPassword,
	}

	for file, content := range files {
		if err := filewriter.Write(file, content, 0644); err != nil {
			log.Error("Failed to write file", "path", file, "err", err)
		}
	}
}

func readmeContent(cfg config.ProjectConfig) string {
	return "*This project has been created as part of the 42 curriculum by " + cfg.StudentLogin + ".*\n\n" +
		"## Description\n" +
		"This project sets up a small Docker-based infrastructure with NGINX, WordPress, and MariaDB.\n" +
		"The goal is to serve WordPress over HTTPS, keep data persistent with named volumes, and manage the services with Docker Compose.\n\n" +
		"## Instructions\n" +
		"Run `make up` to build and start the stack.\n" +
		"Run `make down` to stop it.\n" +
		"Run `make clean` to remove containers and volumes.\n" +
		"Run `make fclean` to remove containers, volumes, and images.\n\n" +
		"Access the website at `https://" + cfg.DomainName + "` and the WordPress admin panel at `https://" + cfg.DomainName + "/wp-admin`.\n\n" +
		"## Resources\n" +
		"- Docker documentation\n" +
		"- Docker Compose documentation\n" +
		"- WordPress documentation\n" +
		"- MariaDB documentation\n" +
		"- NGINX documentation\n\n" +
		"AI was used to help organize the project files, explain Docker Compose behavior, and draft concise documentation.\n\n" +
		"## Project Description\n" +
		"Docker is used to isolate services into separate containers so each one has its own responsibility.\n" +
		"Compared to virtual machines, Docker is lighter and faster because it shares the host kernel.\n\n" +
		"Secrets are stored in files under `secrets/` instead of being hardcoded in Dockerfiles.\n" +
		"Environment variables in `srcs/.env` are used for non-secret configuration values.\n\n" +
		"Docker networks let the containers communicate privately without exposing internal ports to the host.\n" +
		"Docker volumes are used for persistence, while bind mounts are not used for application storage.\n\n" +
		"The services included in this project are NGINX, WordPress with php-fpm, and MariaDB.\n"
}

func userDocContent(cfg config.ProjectConfig) string {
	return "# User Documentation\n\n" +
		"## Services Provided\n" +
		"- NGINX serves the website over HTTPS.\n" +
		"- WordPress provides the public website and the administration dashboard.\n" +
		"- MariaDB stores the WordPress data.\n\n" +
		"## Start and Stop\n" +
		"- Start the stack with `make up`.\n" +
		"- Stop the stack with `make down`.\n\n" +
		"## Access\n" +
		"- Website: `https://" + cfg.DomainName + "`\n" +
		"- Administration dashboard: `https://" + cfg.DomainName + "/wp-admin`\n\n" +
		"## Credentials\n" +
		"- WordPress admin username: `" + cfg.WPAdminUser + "`\n" +
		"- WordPress admin password: stored in `secrets/credentials.txt`\n" +
		"- WordPress database user password: stored in `secrets/db_password.txt`\n" +
		"- MariaDB root password: stored in `secrets/db_root_password.txt`\n\n" +
		"## Basic Checks\n" +
		"- Use `docker compose ps` to verify that the three services are running.\n" +
		"- Use `docker volume ls` to verify persistent volumes exist.\n" +
		"- If the site does not load, check the containers with `docker compose logs`.\n"
}

func devDocContent(cfg config.ProjectConfig) string {
	return "# Developer Documentation\n\n" +
		"## Prerequisites\n" +
		"- Docker\n" +
		"- Docker Compose\n" +
		"- A Linux virtual machine for the project\n\n" +
		"## Configuration\n" +
		"- Non-secret values are stored in `srcs/.env`.\n" +
		"- Secret values are stored in `secrets/`.\n" +
		"- Persistent host data is stored under `/home/" + cfg.StudentLogin + "/data/`.\n\n" +
		"## Build and Launch\n" +
		"- `make up` builds the images and starts the stack.\n" +
		"- `make build` only builds the images.\n" +
		"- `make down` stops the containers.\n" +
		"- `make clean` removes containers and volumes.\n" +
		"- `make fclean` removes containers, volumes, and images.\n\n" +
		"## Docker Compose Usage\n" +
		"- The compose file is `srcs/docker-compose.yml`.\n" +
		"- NGINX is the only service exposed to the host on port 443.\n" +
		"- WordPress talks to MariaDB on the private Docker network.\n\n" +
		"## Data Persistence\n" +
		"- MariaDB data is stored in the named volume mapped to `/home/" + cfg.StudentLogin + "/data/mariadb`.\n" +
		"- WordPress files are stored in the named volume mapped to `/home/" + cfg.StudentLogin + "/data/wordpress`.\n" +
		"- Data remains after container recreation as long as the volumes are kept.\n"
}

func envContent(cfg config.ProjectConfig) string {
	return "LOGIN=" + cfg.StudentLogin + "\n" +
		"DOMAIN_NAME=" + cfg.DomainName + "\n" +
		"MYSQL_DATABASE=" + cfg.MysqlDatabase + "\n" +
		"MYSQL_USER=" + cfg.MysqlUser + "\n" +
		"WP_SITE_TITLE=" + cfg.WPTitle + "\n" +
		"WP_ADMIN_USER=" + cfg.WPAdminUser + "\n" +
		"WP_USER=" + cfg.WPUser + "\n" +
		"WP_ADMIN_EMAIL=" + cfg.WPAdminEmail + "\n" +
		"WP_USER_EMAIL=" + cfg.WPUserEmail + "\n"
}
