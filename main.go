package main

import (
	"errors"
	"fmt"
	"github.com/ataraspost/cli_stub/utilities"
	"github.com/urfave/cli/v2"
	"log"
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
					&cli.StringFlag{
						Name:        "domain",
						Aliases: []string{"d"},
						Value:       "test.ru",
						Usage:       "name domain",
					},
					&cli.BoolFlag{
						Name:        "nginx",
						Value:       false,
						Usage:       "use nginx",
					},
					&cli.BoolFlag{
						Name: "ci_cd",
						Value: false,
						Usage: "create file for ci/cd gitlab",
					},
					&cli.BoolFlag{
						Name: "master",
						Value: false,
						Usage: "create master docker-compose file",
					},
					&cli.BoolFlag{
						Name: "develop",
						Value: false,
						Usage: "create develop docker-compose file",
					},
				},
				Action: func(c *cli.Context) error {
					path := c.String("path")
					name := c.String("name")
					nginx := c.Bool("nginx")
					domain := c.String("domain")
					ci_cd := c.Bool("ci_cd")
					master := c.Bool("master")
					develop := c.Bool("develop")

					work_dir := utilities.GetWorkDirPath(path)

					if _, err := os.Stat(work_dir); !os.IsNotExist(err) {

						isEmpty, err :=  utilities.IsEmpty(work_dir)
						if err != nil {
							return err
						}

						if  !isEmpty {
							fmt.Println("Directory exists")
							err := errors.New("directory exist")
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

					if err := utilities.CreateDockerComposeFile(path, name, domain); err != nil {
						err := errors.New("not create docker-compose.yml settings")
						return err
					}

					if master {
						if err := utilities.CreateDockerComposeFileMaster(path, name, domain); err != nil {
							err := errors.New("not create docker-compose.yml settings")
							return err
						}
					}

					if develop {
						if err := utilities.CreateDockerComposeFileDevelop(path, name, domain); err != nil {
							err := errors.New("not create docker-compose.yml settings")
							return err
						}
					}

					if nginx {
						if err := utilities.CreateNginxSettings(path, domain); err != nil {
							err := errors.New("not create nginx settings")
							return err
						}
					}

					if ci_cd {
						if err := utilities.CreateCiCdSettings(path); err != nil {
							err := errors.New("not create ci/cd settings")
							return err
						}
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

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}