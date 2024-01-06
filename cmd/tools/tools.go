package tools

import (
	"github.com/spf13/cobra"
)

// ToolsCmd represents the tools command
var ToolsCmd = &cobra.Command{
	Use:   "tools",
	Short: "A collection of tools use for manipulating data",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	ToolsCmd.AddCommand(Pdf2htmlCmd)
	ToolsCmd.AddCommand(Tsv2SqlCmd)
}
