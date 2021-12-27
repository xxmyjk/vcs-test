package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
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
	p := promptui.Prompt{
		Label: "vcs >>",
	}

	for {
		rs, err := p.Run()
		fmt.Println(rs, err)
		if rs == "exit" {
			break
		}
	}
}

func init() {
	rootCmd.AddCommand(vcsCmd)
}
