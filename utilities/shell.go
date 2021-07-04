package utilities

import (
	"bytes"
	"fmt"
	"os/exec"
)

const ShellToUse = "bash"

func ShellOut(command string) (error, string, string) {
    var stdout bytes.Buffer
    var stderr bytes.Buffer
    cmd := exec.Command(ShellToUse, "-c", command)
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    err := cmd.Run()
    return err, stdout.String(), stderr.String()
}

func ResultShelCmd(cmd string) {

	err, out, errout := ShellOut(cmd)

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