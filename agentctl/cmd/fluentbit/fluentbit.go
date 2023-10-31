package fluentbit

import (
	logger "agentctl/log"

	"github.com/spf13/cobra"
)

var log = logger.New()

var FluentbitCmd = &cobra.Command{
	Use:   "fluentbit",
	Short: "Manage fluentbit service",
	Long:  ``,
	Run:   UpdateFluentbitHost,
}

func UpdateFluentbitHost(cmd *cobra.Command, args []string) {
	log.Info("Fluentbit is startting")
	// Do sth
}
