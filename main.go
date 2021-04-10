package main

import (
	"bytes"
  	"fmt"
  	"os"
  	"os/exec"
  	"strings"
	"errors" 

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
			},
			Action: func(c *cli.Context) error {
				path := c.String("path")
				name := c.String("name")
				work_dir := getWorkDirPath(path)

				if _, err := os.Stat(work_dir); !os.IsNotExist(err) {
					fmt.Println("Directory exists")
					err := errors.New("Directry exist")
					return err
				}

				if _, err := os.Stat(work_dir); os.IsNotExist(err) {
					cmd := "mkdir -p " + work_dir
					resultShelCmd(cmd)
				}
				
				cmd := "cd " + work_dir + "&& git clone " + StubGit
				resultShelCmd(cmd)

				cmd = "cd " + work_dir + "/stub && rm -rf .git"
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