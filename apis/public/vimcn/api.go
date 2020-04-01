package vimcn

import (
	"fmt"
	"github.com/spf13/cobra"
	"transfer/apis"
	"transfer/utils"
)

var (
	Backend = new(vimcn)
)

type vimcn struct {
	apis.Backend
	resp     string
	Config   cbOptions
	Commands [][]string
}

func (b *vimcn) SetArgs(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&b.Config.DebugMode, "verbose", "", false, "Verbose mode to debug")
	cmd.Long = fmt.Sprintf("vim-cn - https://img.vim-cn.com/\n\n" +
		utils.Spacer("  Size Limit: 100M\n") +
		utils.Spacer("  Upload Service: Cloudflare, Global\n") +
		utils.Spacer("  Download Service: Cloudflare, Global\n"))
}