package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of zimt",
	Long:  `All software has versions. This is zimt's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("zimt v0.1.0")
	},
}
