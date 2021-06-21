package main

import (
	"errors"
	"fmt"
	"github.com/ataraspost/cli_stub/utilities"
	"gopkg.in/urfave/cli.v2"
	"os"
)

const StubGit =  "git@gitlab.com:ataraspost/stub.git"


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
					&cli.BoolFlag{
						Name: "ci_cd",
						Value: false,
						Usage: "name project",
					},
					&cli.BoolFlag{
						Name: "master",
						Value: false,
						Usage: "name project",
					},
					&cli.BoolFlag{
						Name: "develop",
						Value: false,
						Usage: "name project",
					},
				},
				Action: func(c *cli.Context) error {
					path := c.String("path")
					name := c.String("name")
					nginx := c.Bool("nginx")
					domain := c.String("domain")

					work_dir := utilities.GetWorkDirPath(path)

					if _, err := os.Stat(work_dir); !os.IsNotExist(err) {

						isEmpty, err :=  utilities.IsEmpty(work_dir)
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
						utilities.ResultShelCmd(cmd)
					}

					cmd := "cd " + work_dir + "&& git clone " + StubGit
					utilities.ResultShelCmd(cmd)

					cmd = "cd " + work_dir + "/stub && rm -rf .git"
					utilities.ResultShelCmd(cmd)

					utilities.CreateDockerComposeFile(path, name)

					if nginx {
						utilities.CreateNginxSettings(path, name, domain)
					}

					cmd =  "cd " + work_dir + "/stub && rm -rf templates"
					utilities.ResultShelCmd(cmd)

					if name != "stub" {
						cmd := "mv " + work_dir + "/stub " + work_dir + "/" + name
						utilities.ResultShelCmd(cmd)
					}
					return nil
				},
			},
		},
	}

	app.Run(os.Args)
}