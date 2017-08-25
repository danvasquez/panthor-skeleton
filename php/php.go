package php

import (
	"github.com/danvasquez/panthor-skeleton/types"
	"fmt"
	"io/ioutil"
	"os"
)

func WriteConfigs(settings types.AppSettings) {
	configDirectory := fmt.Sprintf("%s/%s/configuration", settings.Location, settings.Name)
	publicDirectory := fmt.Sprintf("%s/%s/public", settings.Location, settings.Name)
	srcDirectory := fmt.Sprintf("%s/%s/src", settings.Location, settings.Name)

	writeConfigYml(configDirectory)
	writeDiYml(configDirectory, settings.Namespace, settings.Name)
	writeRoutesYml(configDirectory)
	writeBootstrapPhp(configDirectory,settings.Namespace, settings.Name)
	writeIndexPhp(publicDirectory, settings.Namespace, settings.Name)
	writeHomeControllerPhp(srcDirectory, settings.Namespace, settings.Name)
}

func writeHomeControllerPhp(directory string, namespace string, name string) {
	stringSlug :=`<?php

namespace %s\%s;

use Psr\Http\Message\ResponseInterface;
use Psr\Http\Message\ServerRequestInterface;
use QL\Panthor\ControllerInterface;

class HomeController implements ControllerInterface
{
    public function __invoke(ServerRequestInterface $request, ResponseInterface $response)
    {
        $response->getBody()->write('Hello World!');
        return $response;
    }
}
`
	configToWrite := fmt.Sprintf(stringSlug, namespace, name)

	writeFile(directory, "HomeController.php", configToWrite)
}

func writeIndexPhp(directory string, namespace string, name string) {
	stringSlug :=`<?php

namespace %s\%s\Bootstrap;

if (!$container = @include __DIR__ . '/../configuration/bootstrap.php') {
    http_response_code(500);
    echo "The application failed to start.\n";
    exit;
};

$slim = $container->get('slim');
$routes = $container->get('router.loader');

$routes($slim);
$slim->run();
`
	configToWrite := fmt.Sprintf(stringSlug, namespace, name)

	writeFile(directory, "index.php", configToWrite)
}

func writeBootstrapPhp(directory string, namespace string, name string) {
	stringSlug :=`<?php

namespace %[1]s\%[2]s\Bootstrap;

use QL\Panthor\Bootstrap\Di;
use %[1]s\%[2]s\CachedContainer;

$root = __DIR__ . '/..';
require_once $root . '/vendor/autoload.php';

return Di::getDi($root, CachedContainer::class);
`
	configToWrite := fmt.Sprintf(stringSlug, namespace, name)

	writeFile(directory, "bootstrap.php", configToWrite)
}

func writeConfigYml(directory string) {
	configString :=`imports:
  - resource: ../vendor/ql/mcp-panthor/configuration/panthor-slim.yml
  - resource: ../vendor/ql/mcp-panthor/configuration/panthor.yml
  - resource: di.yml
  - resource: routes.yml
`
	writeFile(directory, "config.yml", configString)
}

func writeDiYml(directory string, namespace string, name string) {
	configString :=`services:
  page.home:
    class: '%s\%s\HomeController'
    `

	configToWrite := fmt.Sprintf(configString, namespace, name)

	writeFile(directory, "di.yml", configToWrite)
}

func writeRoutesYml(directory string) {
	configString :=`parameters:
  routes:
    home:
      route: '/'
      stack: ['page.home']
    `
	writeFile(directory, "routes.yml", configString)
}

func writeFile(directory string, filename string, contents string) {
	fileToWrite := fmt.Sprintf("%s/%s", directory, filename)

	ioutil.WriteFile(fileToWrite, []byte(contents), os.ModePerm)
}
