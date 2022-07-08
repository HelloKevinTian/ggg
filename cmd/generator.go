package cmd

import (
	"ggg/generator"

	"github.com/spf13/cobra"
)

var generatorCmd = &cobra.Command{
	Use:   "generator",
	Short: "generator test",
	Run: func(cmd *cobra.Command, args []string) {
		generator.Test()
	},
}
