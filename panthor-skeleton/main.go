package main

import("fmt"
	"os"
	"github.com/danvasquez/panthor-skeleton/types"
	"github.com/danvasquez/panthor-skeleton/composer"
	"github.com/danvasquez/panthor-skeleton/php"
	"strings"
)

func main() {
	fmt.Println("Panthor Skeleton Generator")

	appSettings := getAppSettings()

	fmt.Printf("Creating application %s \n", appSettings.Name)

	projectDirectory := createBaseDirectoryStructure(appSettings.Name, appSettings.Location)

	os.Chdir(projectDirectory)

	composer.Init(appSettings)
	php.WriteConfigs(appSettings)
}

func getAppSettings() types.AppSettings {

	return types.AppSettings{
		collectStringFromUser("Application Name"),
		collectStringFromUser("Application Namespace"),
		collectStringFromUser("Author"),
		collectPathFromUser("Project's parent directory"),
	collectBoolFromUser("Are you using Twig for HTML templating (Y/N)?")}
}

func collectBoolFromUser(name string) bool {
	input := strings.ToUpper(collectStringFromUser(name))

	if input !="Y" && input !="N" {
		fmt.Println("Please answer Y or N")
		return collectBoolFromUser(name)
	} else if input =="Y" {
		return true
	} else if input =="N" {
		return false
	}

	return false
}

func collectPathFromUser(name string) string {
	input := collectStringFromUser(name)

	if _, err := os.Stat(input); os.IsNotExist(err) {
		fmt.Printf("Path %s could not be found! \n", input)
		return collectPathFromUser(name)
	}

	return input
}

func collectStringFromUser(name string) string {
	var input string
	fmt.Printf( "%s: ", name)
	fmt.Scanln(&input)

	return input
}

func createBaseDirectoryStructure(appName string, parentPath string) string {
	dirString := "%s/%s/%s"
	directoriesToCreate := [...]string{"src", "tests", "bin", "public", "configuration"}

	for _, dir := range directoriesToCreate {

		dirToCreate := fmt.Sprintf(dirString, parentPath, appName, dir)
		err := os.MkdirAll(dirToCreate, os.ModePerm)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	return fmt.Sprintf("%s/%s", parentPath, appName)
}
