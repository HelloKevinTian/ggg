package cmd

import (
	"ggg/examples"

	"github.com/spf13/cobra"
)

var examplesCmd = &cobra.Command{
	Use:   "examples",
	Short: "examples test",
	Run: func(cmd *cobra.Command, args []string) {
		examples.RunExamples()
	},
}
