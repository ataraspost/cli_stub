package utilities

import (
	"github.com/ataraspost/cli_stub/structures"
	"html/template"
	"os"
)

func CreateDockerComposeFile(path string, name string, domain string) error {
	context := structures.ContextDocker {
		Name: name,
		Master: false,
		Develop: false,
		Domain: domain,
	}

	path_template := path + "/stub/templates/docker/docker-compose.yml"

	tmpl, err := template.ParseFiles(path_template)

	if err != nil {
		return err
	}


	file, err := os.Create(path + "/stub/docker-compose.yml")
	if err != nil {
		return err
	}

	err = tmpl.Execute(file, context)

	if err != nil {
		return err
	}

	if err := file.Close(); err != nil {
		return err
	}

	return nil

}
