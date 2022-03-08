package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zimt",
	Short: "zimt can tell you time",
	Long:  `zimt can tell you everything about time`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("please provide a command to zimt, use `zimt help`")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
