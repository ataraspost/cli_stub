package utilities

import (
	"github.com/ataraspost/cli_stub/structures"
	"os"
	"html/template"
)

func CreateNginxSettings(path string, name string, domain string) error {
	path_dir_template := path + "stub/templates/nginx/"

	path_dir_nginx := path + "stub/nginx/"

	err := os.MkdirAll(path_dir_nginx, os.ModePerm)

	if err != nil {
		return err
	}

	err = Copy(path_dir_template+"Dockerfile", path_dir_nginx+"Dockerfile")
	if err != nil {
		return err
	}
	err = Copy(path_dir_template+".htpasswd", path_dir_nginx+".htpasswd")
	if err != nil {
		return err
	}

	context := structures.ContextNginx {
		Domain: domain,
	}

	tmpl, _ := template.ParseFiles(path_dir_template+"service.conf")

	file, err := os.Create(path + "stub/nginx/service.conf")
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
