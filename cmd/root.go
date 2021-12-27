package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "test",
	Short: "A generator for Cobra based Applications",
}

func Execute() error {
	return rootCmd.Execute()
}
