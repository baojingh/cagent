package cmd

import "github.com/spf13/cobra"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version number",
	Long:  `thisi s version for log desc`,
	Run: func(cmd *cobra.Command, args []string) {
		println("thisi s verion")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
