package suricata_test

import (
	"agent-server/server/suricata"
	"testing"
)

func TestSetFile(t *testing.T) {
	f := suricata.File{}
	f.SetFile("test.txt", "/data/tmp/rules")

}
