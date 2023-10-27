package server

import (
	"fmt"
	"os"
	"os/exec"
)

func shellENV() string {
	s := os.Getenv("SHELL")
	if s == "" {
		s = "/bin/bash"
	}
	log.Infof("Current host has shell env: %s", s)
	return s
}

// Do detailed shell command
func ExecuteShellCmd(name string, param string) (string, error) {
	shellType := shellENV()
	// shell script should have privilege to execute
	exec.Command(shellType, "chmod u+x command/*.sh")
	cmd := exec.Command(fmt.Sprintf("command/%s.sh", name), param)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
