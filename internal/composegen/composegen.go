package composegen

import (
	"os"
	"path/filepath"
	"github.com/charmbracelet/log"
)

func Exec(scopePath string) {
	filename := filepath.Join(scopePath, "docker-compose.yml")
	f, err := os.Create(filename)
	if err != nil {
		log.Error("failed to create docker-compose", "err", err)
		return
	}
	defer f.Close()

	fillContent(f)
}

func fillContent(file *os.File){
	content:= `
services:
  mariadb:
    build: requirements/mariadb
    container_name: mariadb
    image: mariadb
    restart: always
  
  nginx:
    build: requirements/nginx
    container_name: nginx
    image: nginx
    restart: always
    ports:
      - "443:443"`
	_, err := file.WriteString(content)
		if err != nil {
			log.Error("failed to write docker-compose content", "err", err)
		}
}

