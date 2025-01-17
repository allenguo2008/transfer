package catbox

import (
	"fmt"
	"github.com/spf13/cobra"
	"transfer/apis"
	"transfer/utils"
)

var (
	Backend = new(catBox)
)

type catBox struct {
	apis.Backend
	resp     string
	Commands [][]string
}

func (b *catBox) SetArgs(cmd *cobra.Command) {
	cmd.Long = fmt.Sprintf("CatBox - https://catbox.moe/\n\n" +
		utils.Spacer("  Size Limit: 200M\n") +
		utils.Spacer("  Upload Service: Psychz Networks, Los Angeles California, USA\n") +
		utils.Spacer("  Download Service: Psychz Networks, Los Angeles California, USA\n"))
}
