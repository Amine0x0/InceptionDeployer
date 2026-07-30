package makefilegen

import (
	"InceptionDeployer/config"
	"InceptionDeployer/internal/filewriter"
	"path/filepath"

	"github.com/charmbracelet/log"
)

func Exec(basePath string, cfg config.ProjectConfig) {
	_ = cfg
	filename := filepath.Join(basePath, "Makefile")
	if err := filewriter.Write(filename, fillContent(), 0644); err != nil {
		log.Error("failed to write Makefile", "err", err)
	}
}

func fillContent() string {
	return `include srcs/.env

DOCKER_COMPOSE_FILE = ./srcs/docker-compose.yml
ENV_FILE = ./srcs/.env
DATA_DIR = /home/$(LOGIN)/data
COMPOSE = docker compose --env-file $(ENV_FILE) -f $(DOCKER_COMPOSE_FILE)

hosts:
	@if ! grep -q "$(DOMAIN_NAME)" /etc/hosts; then \
		echo "127.0.0.1 $(DOMAIN_NAME) www.$(DOMAIN_NAME)" | sudo tee -a /etc/hosts > /dev/null; \
	else \
		echo "Host entry already exists."; \
	fi

up: hosts
	@mkdir -p $(DATA_DIR)/wordpress $(DATA_DIR)/mariadb
	@$(COMPOSE) up --build

build: hosts
	@$(COMPOSE) build

down:
	@$(COMPOSE) down

clean:
	@$(COMPOSE) down -v

fclean:
	@$(COMPOSE) down -v --rmi all

rmdata:
	@sudo rm -rf $(DATA_DIR)/wordpress $(DATA_DIR)/mariadb

re: clean up
`
}
