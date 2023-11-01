package command

import "testing"

func TestExec(t *testing.T) {
	// agent-server/command/test.sh
	ExecuteShellCmd("fluentbit_host", "host=11.22.99.00")
}
