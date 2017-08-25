package composer

import (
	"io/ioutil"
	"os"
	"github.com/danvasquez/panthor-skeleton/types"
	"fmt"
	"os/exec"
)

func Init(appSettings types.AppSettings) {
	fileContents := createBasicComposerFile(appSettings)
	ioutil.WriteFile("composer.json", []byte(fileContents), os.ModePerm)

	addPackage("ql/mcp-panthor", "~3.0")
	addPackage("paragonie/random_compat", "~1.1")

	if appSettings.UsingTwig ==true {
		addPackage("twig/twig", "~1.20")
	}
}

func createBasicComposerFile(settings types.AppSettings) string {
	var fileTemplate = `{
"name": "%[1]s/%[2]s",
"description": "",
"autoload": {
	"psr-4": {
		"%[1]s\\%[2]s\\": "src/"
		}
	},
"authors": [
	{
		"name": "%[3]s"
	}
],
"scripts": {
  "start": "php -S 0.0.0.0:8080 public/index.php",
  "test": "phpunit"
}

}
	`
	return fmt.Sprintf(fileTemplate, settings.Namespace, settings.Name, settings.Author)
}

func addPackage(packageName string, version string) {
	cmd := exec.Command("composer", "require", packageName, version)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("NOTE: this program assumes you have composer installed and executable as 'composer'! See https://getcomposer.org/ for more.")
	}
}