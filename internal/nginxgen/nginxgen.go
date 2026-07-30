package nginxgen

import (
	"InceptionDeployer/config"
	"InceptionDeployer/internal/filewriter"
	"fmt"
	"path/filepath"

	"github.com/charmbracelet/log"
)

func Exec(scopePath string, cfg config.ProjectConfig) {
	basePath := filepath.Join(scopePath, "requirements", "nginx")
	if err := filewriter.Write(filepath.Join(basePath, "Dockerfile"), dockerfileContent(cfg), 0644); err != nil {
		log.Error("failed to write nginx Dockerfile", "err", err)
	}
	if err := filewriter.Write(filepath.Join(basePath, "conf", "default.conf"), confContent(cfg), 0644); err != nil {
		log.Error("failed to write nginx config", "err", err)
	}
	if err := filewriter.Write(filepath.Join(basePath, ".dockerignore"), "", 0644); err != nil {
		log.Error("failed to write nginx dockerignore", "err", err)
	}
}

func dockerfileContent(cfg config.ProjectConfig) string {
	return fmt.Sprintf(`FROM debian:bookworm

RUN apt-get update && apt-get install -y nginx openssl && rm -rf /var/lib/apt/lists/*

RUN mkdir -p /etc/nginx/certs /run/nginx && \
	openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
	-keyout /etc/nginx/certs/incep.key \
	-out /etc/nginx/certs/incep.crt \
	-subj "/C=MA/ST=KHOURIBGA/L=KHOURIBGA/O=1337/OU=STUDENT/CN=%s"

COPY conf/default.conf /etc/nginx/conf.d/default.conf

ENTRYPOINT ["nginx", "-g", "daemon off;"]`, cfg.DomainName)
}

func confContent(cfg config.ProjectConfig) string {
	return fmt.Sprintf(`server {
    listen 443 ssl;
    server_name www.%s %s;

    root /var/www/html;
    index index.php index.html;

    ssl_certificate /etc/nginx/certs/incep.crt;
    ssl_certificate_key /etc/nginx/certs/incep.key;
    ssl_protocols TLSv1.2 TLSv1.3;

    location / {
        try_files $uri $uri/ /index.php$is_args$args;
    }

    location ~ \.php$ {
        include fastcgi_params;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
        fastcgi_pass wordpress:9000;
    }
}
`, cfg.DomainName, cfg.DomainName)
}
