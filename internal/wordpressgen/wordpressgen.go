package wordpressgen

import (
	"InceptionDeployer/config"
	"InceptionDeployer/internal/filewriter"
	"path/filepath"

	"github.com/charmbracelet/log"
)

func Exec(scopePath string, cfg config.ProjectConfig) {
	basePath := filepath.Join(scopePath, "requirements", "wordpress")
	if err := filewriter.Write(filepath.Join(basePath, "Dockerfile"), dockerfileContent(), 0644); err != nil {
		log.Error("failed to write wordpress Dockerfile", "err", err)
	}
	if err := filewriter.Write(filepath.Join(basePath, "tools", "entrypoint.sh"), entrypointContent(), 0755); err != nil {
		log.Error("failed to write wordpress entrypoint", "err", err)
	}
}

func dockerfileContent() string {
	return `FROM debian:12

RUN apt-get update -y && apt-get install -y \
	curl \
	php \
	php-fpm \
	php-mysql \
	default-mysql-client \
	&& rm -rf /var/lib/apt/lists/*

RUN curl -fsSL https://raw.githubusercontent.com/wp-cli/builds/gh-pages/phar/wp-cli.phar -o /usr/local/bin/wp \
	&& chmod +x /usr/local/bin/wp

RUN mkdir -p /run/php /var/www/html
RUN chown -R www-data:www-data /var/www/html
RUN sed -i 's|listen = .*|listen = 9000|' /etc/php/*/fpm/pool.d/www.conf

COPY ./tools/entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

ENTRYPOINT ["/app/entrypoint.sh"]

`
}

func entrypointContent() string {
	return `#!/bin/sh

set -eu

MYSQL_PASS=$(cat /run/secrets/db_password)
WP_ADMIN_PASS=$(cat /run/secrets/credentials)
WP_USER_PASS=$(cat /run/secrets/credentials)

mkdir -p /run/php /var/www/html
rm -f /var/www/html/wp-content/mu-plugins/auto-approve-comments.php

until mariadb -h mariadb -u"$MYSQL_USER" -p"$MYSQL_PASS" -e "SELECT 1" >/dev/null 2>&1; do
	echo "Still waiting for mariadb..."
	sleep 2
done

cd /var/www/html

if [ ! -f wp-config.php ]; then
	wp core download --allow-root

	wp config create \
		--allow-root \
		--force \
		--dbname="$MYSQL_DATABASE" \
		--dbuser="$MYSQL_USER" \
		--dbpass="$MYSQL_PASS" \
		--dbhost="mariadb"

	wp core install \
		--allow-root \
		--url="https://$DOMAIN_NAME" \
		--title="$WP_SITE_TITLE" \
		--admin_user="$WP_ADMIN_USER" \
		--admin_password="$WP_ADMIN_PASS" \
		--admin_email="$WP_ADMIN_EMAIL"

	wp user create "$WP_USER" "$WP_USER_EMAIL" \
		--role=subscriber \
		--user_pass="$WP_USER_PASS" \
		--allow-root
fi

exec /usr/sbin/php-fpm8.2 -F
`
}
