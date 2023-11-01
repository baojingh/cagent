package suricata

import (
	"github.com/spf13/cobra"
)

var SuricataUploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload suricata files",

	Run: UploadSuricataFile,
}

func UploadSuricataFile(cmd *cobra.Command, args []string) {

}
