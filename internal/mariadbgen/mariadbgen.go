package mariadbgen

import (
	"InceptionDeployer/config"
	"InceptionDeployer/internal/filewriter"
	"path/filepath"

	"github.com/charmbracelet/log"
)

func Exec(scopePath string, cfg config.ProjectConfig) {
	basePath := filepath.Join(scopePath, "requirements", "mariadb")
	if err := filewriter.Write(filepath.Join(basePath, "Dockerfile"), dockerfileContent(), 0644); err != nil {
		log.Error("failed to write mariadb Dockerfile", "err", err)
	}
	if err := filewriter.Write(filepath.Join(basePath, "conf", "50-server.cnf"), confContent(), 0644); err != nil {
		log.Error("failed to write mariadb config", "err", err)
	}
	if err := filewriter.Write(filepath.Join(basePath, "tools", "entrypoint.sh"), entrypointContent(), 0755); err != nil {
		log.Error("failed to write mariadb entrypoint", "err", err)
	}
	if err := filewriter.Write(filepath.Join(basePath, ".dockerignore"), "", 0644); err != nil {
		log.Error("failed to write mariadb dockerignore", "err", err)
	}
}

func dockerfileContent() string {
	return `FROM debian:bookworm

RUN apt-get update && apt-get install -y mariadb-server mariadb-client && rm -rf /var/lib/apt/lists/*

COPY tools/entrypoint.sh /usr/local/bin/entrypoint.sh
COPY conf/50-server.cnf /etc/mysql/mariadb.conf.d/50-server.cnf
RUN chmod +x /usr/local/bin/entrypoint.sh

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
`
}

func confContent() string {
	return `#
# These groups are read by MariaDB server.

# this is read by the standalone daemon and embedded servers
[server]

# this is only for the mysqld standalone daemon
[mysqld]

#
# * Basic Settings
#

#user                    = mysql
pid-file                = /run/mysqld/mysqld.pid
basedir                 = /usr

skip-name-resolve
bind-address            = 0.0.0.0

#
# * Fine Tuning
#

#key_buffer_size        = 128M
#max_allowed_packet     = 1G
#thread_stack           = 192K
#thread_cache_size      = 8

#
# * Logging and Replication
#

expire_logs_days        = 10

#
# * Character sets
#

character-set-server  = utf8mb4
collation-server      = utf8mb4_general_ci

#
# * InnoDB
#

# this is only for embedded server
[embedded]

[mariadb]

[mariadb-10.11]
`
}

func entrypointContent() string {
	return `#!/bin/sh

set -eu

mkdir -p /run/mysqld
chown -R mysql:mysql /run/mysqld /var/lib/mysql

MYSQL_PASSWORD=$(cat "$MYSQL_PASSWORD_FILE")
MYSQL_ROOT_PASSWORD=$(cat "$MYSQL_ROOT_PASSWORD_FILE")

if [ ! -d /var/lib/mysql/mysql ]; then
    mariadb-install-db --user=mysql --datadir=/var/lib/mysql >/dev/null
fi

mariadbd --user=mysql --skip-networking --socket=/run/mysqld/mysqld.sock &
temp_server=$!

until mariadb-admin --socket=/run/mysqld/mysqld.sock ping >/dev/null 2>&1; do
    sleep 1
done

mariadb --socket=/run/mysqld/mysqld.sock -uroot -p"$MYSQL_ROOT_PASSWORD" <<EOF
ALTER USER 'root'@'localhost' IDENTIFIED BY '${MYSQL_ROOT_PASSWORD}';
CREATE DATABASE IF NOT EXISTS ${MYSQL_DATABASE};
CREATE USER IF NOT EXISTS '${MYSQL_USER}'@'%' IDENTIFIED BY '${MYSQL_PASSWORD}';
GRANT ALL PRIVILEGES ON ${MYSQL_DATABASE}.* TO '${MYSQL_USER}'@'%';
FLUSH PRIVILEGES;
EOF

kill "$temp_server"
wait "$temp_server" 2>/dev/null || true

exec mariadbd --user=mysql --bind-address=0.0.0.0
`
}
