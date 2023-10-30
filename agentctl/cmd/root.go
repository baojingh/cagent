package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hello",
	Short: "thisi s a short",
	Long: `benmiobsdkmp niokvweiop niobsd, nuivbw
			b,mdop[,pbnd]
			`,
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}
