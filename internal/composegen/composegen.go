package composegen

import (
	"InceptionDeployer/config"
	"InceptionDeployer/internal/filewriter"
	"path/filepath"

	"github.com/charmbracelet/log"
)

func Exec(scopePath string, cfg config.ProjectConfig) {
	_ = cfg
	filename := filepath.Join(scopePath, "docker-compose.yml")
	if err := filewriter.Write(filename, fillContent(), 0644); err != nil {
		log.Error("failed to write docker-compose", "err", err)
	}
}

func fillContent() string {
	return `services:
  mariadb:
    build: requirements/mariadb
    image: mariadb
    restart: always
    env_file:
      - .env
    environment:
      MYSQL_ROOT_PASSWORD_FILE: /run/secrets/db_root_password
      MYSQL_PASSWORD_FILE: /run/secrets/db_password
    secrets:
      - db_root_password
      - db_password
    volumes:
      - mariadb_data:/var/lib/mysql
    networks:
      - inception

  wordpress:
    build: requirements/wordpress
    image: wordpress
    restart: always
    env_file:
      - .env
    environment:
      MYSQL_HOST: mariadb
      MYSQL_PASSWORD_FILE: /run/secrets/db_password
      WP_PASSWORD_FILE: /run/secrets/credentials
    secrets:
      - db_password
      - credentials
    volumes:
      - wordpress_data:/var/www/html
    depends_on:
      - mariadb
    networks:
      - inception

  nginx:
    build: requirements/nginx
    image: nginx
    restart: always
    ports:
      - "443:443"
    volumes:
      - wordpress_data:/var/www/html
    depends_on:
      - wordpress
    networks:
      - inception

volumes:
  mariadb_data:
    driver: local
    driver_opts:
      type: none
      o: bind
      device: /home/${LOGIN}/data/mariadb
  wordpress_data:
    driver: local
    driver_opts:
      type: none
      o: bind
      device: /home/${LOGIN}/data/wordpress

networks:
  inception:
    driver: bridge

secrets:
  db_root_password:
    file: ../secrets/db_root_password.txt
  db_password:
    file: ../secrets/db_password.txt
  credentials:
    file: ../secrets/credentials.txt
`
}
