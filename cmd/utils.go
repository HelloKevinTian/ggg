package cmd

import (
	"ggg/utils"

	"github.com/spf13/cobra"
)

var utilsCmd = &cobra.Command{
	Use:   "utils",
	Short: "utils test",
	Run: func(cmd *cobra.Command, args []string) {
		utils.Test()
	},
}
