package suricata

/*
	command entrypoint

*/

import (
	"agentctl/constant"
	logger "agentctl/log"

	"github.com/spf13/cobra"
)

var log = logger.New()
var (
	filepath  string
	batchSize int
)

var SuricataCmd = &cobra.Command{
	Use:   "suricata",
	Short: "Manage suricata service",
	Long:  ``,
	Run:   nil,
}

var SuricataUploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload suricata files",
	Run: func(cmd *cobra.Command, args []string) {
		ip, _ := cmd.Flags().GetString("ip")
		port, _ := cmd.Flags().GetString("port")
		service := NewService(ip, port, filepath, batchSize)
		service.UploadSuricataFile()
	},
}

func init() {
	SuricataUploadCmd.MarkFlagRequired(constant.SURICATA_PARAM_FILE)
	SuricataUploadCmd.Flags().StringVarP(&filepath, "file", "f", "", "suricata upload file")
	SuricataUploadCmd.Flags().IntVarP(&batchSize, "batch", "b", 1024*1024, "buffer batch size for upload")
	SuricataCmd.AddCommand(SuricataUploadCmd)
}
