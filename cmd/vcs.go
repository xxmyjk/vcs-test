package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var vcsCmd = &cobra.Command{
	Use:   "vcs",
	Short: "vcs test....",
	Run: func(cmd *cobra.Command, args []string) {
		test()
	},
}

func test() {
	fmt.Println("hi~")
}

func init() {
	rootCmd.AddCommand(vcsCmd)
}
