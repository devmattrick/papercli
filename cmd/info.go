package cmd

import (
	"github.com/devmattrick/papercli/cli"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info [flags]",
	Short: "obtain information about an subject",
	Run: func(cmd *cobra.Command, args []string) {
		if Version == "latest" {
			cli.PrintVersions(Project)
		} else if Build == "latest" {
			cli.PrintBuilds(Project, Version)
		} else {
			// TODO Print build
		}
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
