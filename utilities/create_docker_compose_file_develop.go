package utilities

import (
	"github.com/ataraspost/cli_stub/structures"
	"html/template"
	"os"
)

func CreateDockerComposeFileDevelop(path string, name string, domain string) error {
	context := structures.ContextDocker {
		Name: name,
		Master: false,
		Develop: true,
		Domain: domain,
	}

	path_template := path + "/stub/templates/docker/docker-compose.yml"

	file, err := os.Create(path + "/stub/docker-compose.develop.yml")
	if err != nil {
		return err
	}

	tmplDevelop, err := template.ParseFiles(path_template)

	if err != nil {
		return err
	}

	err = tmplDevelop.Execute(file, context)

	if err != nil {
		return err
	}

	if err := file.Close(); err != nil {
		return err
	}

	return nil

}

