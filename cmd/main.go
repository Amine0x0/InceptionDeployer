package main

import (
	"InceptionDeployer/config"
	"InceptionDeployer/internal/banner"
	"InceptionDeployer/internal/builder"
	"InceptionDeployer/internal/orchestrator"
	Color "InceptionDeployer/tooling"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/sqweek/dialog"
	"golang.org/x/term"
)

func logme(msg string) {
	fmt.Printf("%s", Color.Purple+msg+Color.Reset+": ")
}

func promptText(label string, fallback string) string {
	scanner := bufio.NewScanner(os.Stdin)
	logme(label)
	if scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			return line
		}
	}
	return fallback
}

func getProjectPath(projectName string) string {
	fmt.Println("Select where you wanna Store your project")
	directory, err := dialog.Directory().
		Title("Select Where to Create Your Inception Project").
		Browse()
	if err != nil || directory == "" {
		log.Warn("Folder selection cancelled or failed. Falling back to current directory.", "err", err)
		cwd, _ := os.Getwd()
		return filepath.Join(cwd, projectName)
	}
	return filepath.Join(directory, projectName)
}

func promptSecret(label string, fallback string) string {
	fmt.Printf("%s", Color.Purple+label+Color.Reset+": ")
	value, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	if err == nil {
		line := strings.TrimSpace(string(value))
		if line != "" {
			return line
		}
	}
	return fallback
}

func main() {
	banner.Show()
	name := promptText("Enter how you wanna name your Project", config.DefaultProjectName)
	login := promptText("Enter your 42 login", "DarthVader")
	path := getProjectPath(name)
	domain := promptText("Enter your domain name", login+".42.fr")
	mysqlDatabase := promptText("Enter your database name", "wordpress")
	mysqlUser := promptText("Enter your database user", "wp_user")
	mysqlPassword := promptSecret("Enter your database user password", "dbpassword123")
	mysqlRootPassword := promptSecret("Enter your database root password", "rootpassword123")
	wpTitle := promptText("Enter your WordPress site title", config.DefaultProjectName)
	wpAdminUser := promptText("Enter your WordPress admin username", "supervisor")
	wpAdminEmail := promptText("Enter your WordPress admin email", wpAdminUser+"@"+domain)
	wpAdminPassword := promptSecret("Enter your WordPress admin password", "wordpress123")
	wpUser := promptText("Enter your WordPress subscriber username", "subscriber")
	wpUserEmail := promptText("Enter your WordPress subscriber email", wpUser+"@"+domain)
	wpUserPassword := promptSecret("Enter your WordPress subscriber password", "wordpress123")

	md := config.ProjectConfig{
		ProjectName:       name,
		ProjectPath:       path,
		StudentLogin:      login,
		DomainName:        domain,
		MysqlDatabase:     mysqlDatabase,
		MysqlUser:         mysqlUser,
		MysqlPassword:     mysqlPassword,
		MysqlRootPassword: mysqlRootPassword,
		WPTitle:           wpTitle,
		WPAdminUser:       wpAdminUser,
		WPAdminPassword:   wpAdminPassword,
		WPAdminEmail:      wpAdminEmail,
		WPUser:            wpUser,
		WPUserPassword:    wpUserPassword,
		WPUserEmail:       wpUserEmail,
	}

	log.Info("creating folder structure for ", md.StudentLogin)
	log.Info("Project metadata initialized",
		"name", md.ProjectName,
		"path", md.ProjectPath,
	)

	builder.CreateStructure(md.ProjectPath)
	builder.GenerateConfigs(md.ProjectPath, md)
	orchestrator.Generateall(md.ProjectPath, md)
}
