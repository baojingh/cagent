package handler

import (
	logger "agent-server/log"
	"os/exec"
)

var log = logger.New()

func ExecCmd(command string) {
	cmd := exec.Command(command)
	out, err := cmd.Output()
	if err != nil {
		log.Error("")
	}

}
