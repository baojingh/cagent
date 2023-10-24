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
	log.Infof("Current host has shell eenv: %s", s)
	return s
}

// Do detailed shell command
func Exec(name string) (string, error) {
	shellType := shellENV()
	// shell script should have privilege to execute
	exec.Command(shellType, "chmod u+x command/*.sh")
	cmd := exec.Command(shellType, "-c", fmt.Sprintf("command/%s.sh", name))
	out, err := cmd.Output()
	if err != nil {
		log.Info("Failed to execute the shell,", err)
		return "", err
	}
	log.Infof("Command is executed success, cmd: %s", cmd.String())
	return string(out), nil
}
