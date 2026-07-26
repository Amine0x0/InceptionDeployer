package main

import (
	"InceptionDeployer/config"
	"InceptionDeployer/tooling"
	"InceptionDeployer/internal/builder"
	"InceptionDeployer/internal/banner"
	"InceptionDeployer/internal/orchestrator"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"
	"github.com/charmbracelet/log"
	"github.com/sqweek/dialog"
)

func logme(msg string){
	fmt.Printf("%s", Color.Purple + msg + Color.Reset + ": ");
}

type Metadata struct{
	ProjectName string
	ProjectPath string
	StudentLogin string
}

func getProjectName() string {
	scanner := bufio.NewScanner(os.Stdin)
	logme("Enter How you wanna name your Project")
	if scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			return line
		}
	}
	return config.DefaultProjectName
}

func getProjectPath(projectName string) string {
	fmt.Println("Select where you wanna Store your project")
	time.Sleep(1 * time.Second)
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

func getStudentLogin() string{
	scanner := bufio.NewScanner(os.Stdin)
	logme("Enter your 42 login: ")
	if scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			return line
		}
	}
	return "DarthVader"
}

func main(){
	banner.Show()
	name := getProjectName()

	md := Metadata{
		ProjectName: name,
		ProjectPath: getProjectPath(name),
		StudentLogin: getStudentLogin(),
	}

	log.Info("creating folder structure for ", md.StudentLogin);
	log.Info("Project metadata initialized",
		"name", md.ProjectName,
		"path", md.ProjectPath,
	)

	builder.CreateStructure(md.ProjectPath)
	orchestrator.Generateall(md.ProjectPath)
}