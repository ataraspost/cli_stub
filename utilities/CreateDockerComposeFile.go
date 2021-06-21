package utilities

import (
	"github.com/ataraspost/cli_stub/structures"
	"html/template"
	"os"
)

func CreateDockerComposeFile(path string, name string) error {
	context := structures.ContextDocker {
		Name: name,
	}

	path_template := path + "stub/templates/docker/docker-compose.yml"

	tmpl, _ := template.ParseFiles(path_template)



	file, err := os.Create(path + "stub/docker-compose.yml")
	if err != nil {
		return err
	}

	err = tmpl.Execute(file, context)

	if err != nil {
		return err
	}

	file.Close()

	return nil

}
