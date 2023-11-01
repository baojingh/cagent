package suricata

import (
	"agentctl/constant"

	"github.com/spf13/cobra"
)

var SuricataUploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload suricata files",

	Run: UploadSuricataFile,
}

func UploadSuricataFile(cmd *cobra.Command, args []string) {
	log.Info("upload")

}

func init() {
	SuricataUploadCmd.MarkFlagRequired(constant.SURICATA_PARAM_FILE)
	SuricataUploadCmd.Flags().String(constant.SURICATA_PARAM_FILE, "", "suricata upload file")
}
