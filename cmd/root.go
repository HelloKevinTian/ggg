package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version string

	rootCmd = &cobra.Command{
		Use: "canary-server",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("root go")
		},
	}
)

func init() {
	cobra.OnInitialize(onInit)
	rootCmd.PersistentFlags().StringVar(&version, "version", "latest.", "version")

	rootCmd.AddCommand(utilsCmd)
	rootCmd.AddCommand(generatorCmd)
	rootCmd.AddCommand(singletonCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func onInit() {

}
