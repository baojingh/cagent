package cmd

import (
	"agentctl/cmd/fluentbit"
	"agentctl/cmd/suricata"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var buildTime string
var buildVersion string

var RootCmd = &cobra.Command{
	Use:     "agentctl",
	Short:   "agentctl helps update configuration from server side",
	Long:    `agentctl is client to communicate with grpc server with modules.`,
	Version: fmt.Sprintf("%s, build %s", buildVersion, buildTime),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	RootCmd.AddCommand(fluentbit.FluentbitCmd)
	RootCmd.AddCommand(suricata.SuricataCmd)
}
