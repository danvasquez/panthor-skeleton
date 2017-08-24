package main

import("fmt"
	"os"
	"os/exec"
	"io/ioutil"
)

func main() {
	fmt.Println("Panthor Skeleton")

	appSettings := getAppSettings()

	fmt.Printf("Creating application %s \n", appSettings.Name)

	projectDirectory := createBaseDirectoryStructure(appSettings.Name)

	os.Chdir(projectDirectory)

	composerInit(appSettings)
	composerAddPackage("ql/mcp-panthor", "~3.0")
	composerAddPackage("paragonie/random_compat", "~1.1")
	composerAddPackage("twig/twig", "~1.20")
}

func getAppSettings() AppSettings {

	return AppSettings{getAppName(), getNamespace(), getAuthor()}
}

func composerInit(appSettings AppSettings) {
	composerText := createBasicComposerFile(appSettings)
	ioutil.WriteFile("composer.json", []byte(composerText), os.ModePerm)
}

func composerAddPackage(packageName string, version string) {
	cmd := exec.Command("composer", "require", packageName, version)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func getAppName() string {
	var appName string
	fmt.Print("Enter application name: ")
	fmt.Scanln(&appName)

	return appName
}

func getNamespace() string {
	var namespace string
	fmt.Print("App Namespace: ")
	fmt.Scanln(&namespace)

	return namespace
}

func getAuthor() string {
	var author string
	fmt.Print("Author: ")
	fmt.Scanln(&author)

	return author
}

func createBaseDirectoryStructure(appName string) string {
	currentDirectory := os.Getenv("PWD")
	dirString := "%s/%s/%s"
	directoriesToCreate := [...]string{"src", "tests", "bin", "public", "configuration"}

	for _, dir := range directoriesToCreate {

		dirToCreate := fmt.Sprintf(dirString, currentDirectory, appName, dir)
		os.MkdirAll(dirToCreate, os.ModePerm)
	}

	return fmt.Sprintf("%s/%s", currentDirectory, appName)
}

func createBasicComposerFile(settings AppSettings) string {
	var fileTemplate = `{
"name": "%s/%s",
"description": "",
"autoload": {
	"psr-4": {
		"%s\\%s\\": "src/"
		}
	}
}
	`
	return fmt.Sprintf(fileTemplate, settings.Namespace, settings.Name, settings.Namespace, settings.Name)
}

type AppSettings struct {
	Name string
	Namespace string
	Author string
}

