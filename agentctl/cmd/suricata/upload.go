package suricata

import (
	"github.com/spf13/cobra"
)

var SuricataUploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload suricata files",
	Long:  ``,
	Run:   UploadSuricataFile,
}

func UploadSuricataFile(cmd *cobra.Command, args []string) {

}
