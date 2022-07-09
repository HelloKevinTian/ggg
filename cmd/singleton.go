package cmd

import (
	"ggg/singleton"

	"github.com/spf13/cobra"
)

var singletonCmd = &cobra.Command{
	Use:   "singleton",
	Short: "singleton test",
	Run: func(cmd *cobra.Command, args []string) {
		singleton.Test()
	},
}
