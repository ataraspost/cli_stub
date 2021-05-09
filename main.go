package main

import (
	"bytes"
  	"fmt"
  	"os"
  	"os/exec"
  	"strings"
	"errors" 
	"io"
	"html/template"

  	"gopkg.in/urfave/cli.v2"
)

const StubGit =  "git@gitlab.com:ataraspost/stub.git"

const ShellToUse = "bash"

func getWorkDirPath(path string) string {

	if path == "." {
		work_dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		return work_dir	
	} else if strings.HasPrefix(path, ".") {
		work_dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		work_dir = work_dir + path[1:]
		return work_dir	
	} else if strings.HasPrefix(path, "/") {
		work_dir := path
		return work_dir	
	} else {
		work_dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		work_dir = work_dir + "/" + path
		return work_dir	
	}
}

func shellOut(command string) (error, string, string) {
    var stdout bytes.Buffer
    var stderr bytes.Buffer
    cmd := exec.Command(ShellToUse, "-c", command)
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    err := cmd.Run()
    return err, stdout.String(), stderr.String()
}

func resultShelCmd(cmd string) {

	err, out, errout := shellOut(cmd)

	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	if len(out) != 0 {
		fmt.Println("--- stdout ---")
		fmt.Println(out)
	}

	if len(errout) != 0 {
		fmt.Println("--- stderr ---")
		fmt.Println(errout)
	}
}

type ContextDocker struct{
	Name string
}

func createDockerComposeFiale(path string, name string) error {
	context := ContextDocker {
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

func Copy(src, dst string) error {

    in, err := os.Open(src)
    if err != nil {
        return err
    }
    defer in.Close()

    out, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer out.Close()

    _, err = io.Copy(out, in)
    if err != nil {
        return err
    }
    return  nil
}

type ContextNginx struct{
	Domain string
}

func createNginxSettings(path string, name string, domain string) error {
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

	context := ContextNginx {
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

func IsEmpty(name string) (bool, error) {
    f, err := os.Open(name)
    if err != nil {
        return false, err
    }
    defer f.Close()

    _, err = f.Readdirnames(1) // Or f.Readdir(1)
    if err == io.EOF {
        return true, nil
    }
    return false, err // Either not empty or error, suits both cases
}

func main() {
  app := &cli.App{
    Name: "cli_stub",
    Usage: "fight the loneliness!",
	
	Commands: []*cli.Command{
		{
			Name: "create",
			Flags: []cli.Flag {
				&cli.StringFlag{
				  Name:        "path",
				  Aliases: []string{"p"},
				  Value:       ".",
				  Usage:       "path to create project",
				},
				&cli.StringFlag{
				  Name:        "name",
				  Aliases: []string{"n"},
				  Value:       "stub",
				  Usage:       "name project",
				},
				&cli.BoolFlag{
					Name:        "nginx",
					Value:       false,
					Usage:       "name project",
				  },
				&cli.StringFlag{
					Name:        "domain",
					Value:       "test.ru",
					Usage:       "name project",
				  }, 
			},
			Action: func(c *cli.Context) error {
				path := c.String("path")
				name := c.String("name")
				nginx := c.Bool("nginx")
				domain := c.String("domain")

				work_dir := getWorkDirPath(path)

				if _, err := os.Stat(work_dir); !os.IsNotExist(err) {

					isEmpty, err :=  IsEmpty(work_dir)
					if err != nil {
						return err
					}

					if  !isEmpty {
						fmt.Println("Directory exists")
						err := errors.New("Directry exist")
						return err
					}
				}

				if _, err := os.Stat(work_dir); os.IsNotExist(err) {
					cmd := "mkdir -p " + work_dir
					resultShelCmd(cmd)
				}
				
				cmd := "cd " + work_dir + "&& git clone " + StubGit
				resultShelCmd(cmd)

				cmd = "cd " + work_dir + "/stub && rm -rf .git"
				resultShelCmd(cmd)

				createDockerComposeFiale(path, name)

				if nginx {
					createNginxSettings(path, name, domain)
				}

				cmd =  "cd " + work_dir + "/stub && rm -rf templates"
				resultShelCmd(cmd)
				
				if name != "stub" {
					cmd := "mv " + work_dir + "/stub " + work_dir + "/" + name
					resultShelCmd(cmd)
				}
				return nil
			},
		},
	},
  }

  app.Run(os.Args)
}