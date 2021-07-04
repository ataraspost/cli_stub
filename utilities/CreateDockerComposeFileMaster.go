package utilities

import (
	"github.com/ataraspost/cli_stub/structures"
	"html/template"
	"os"
)

func CreateDockerComposeFileMaster(path string, name string, domain string) error {
	context := structures.ContextDocker {
		Name: name,
		Master: true,
		Develop: false,
		Domain: domain,
	}

	path_template := path + "/stub/templates/docker/docker-compose.yml"

	file, err := os.Create(path + "/stub/docker-compose.master.yml")
	if err != nil {
		return err
	}

	tmplMaster, err := template.ParseFiles(path_template)

	if err != nil {
		return err
	}


	err = tmplMaster.Execute(file, context)

	if err != nil {
		return err
	}

	if err := file.Close(); err != nil {
		return err
	}

	return nil

}

