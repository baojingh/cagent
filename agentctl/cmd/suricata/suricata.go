package suricata

import (
	logger "agentctl/log"

	"github.com/spf13/cobra"
)

var log = logger.New()

var SuricataCmd = &cobra.Command{
	Use:   "suricata",
	Short: "Manage suricata service",
	Long:  ``,
	Run:   UpdateSuricataFile,
}

func UpdateSuricataFile(cmd *cobra.Command, args []string) {
	log.Info("Suricata is startting")
	// Do sth
}
