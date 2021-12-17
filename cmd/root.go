package cmd

import (
	"os"

	"github.com/devmattrick/papercli/cli"
	"github.com/spf13/cobra"
)

var Project string
var Version string
var Build string

var rootCmd = &cobra.Command{
	Use:   "papercli [flags] [subcommand]",
	Short: "papercli is a utility to access PaperMC's REST API",
	Args:  cobra.RangeArgs(0, 3),
	Run: func(cmd *cobra.Command, args []string) {
		cli.Download(Project, Version, Build)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&Project, "project", "p", "paper", "the PaperMC project id")
	rootCmd.PersistentFlags().StringVarP(&Version, "version", "V", "latest", "the version of the project")
	rootCmd.PersistentFlags().StringVarP(&Build, "build", "b", "latest", "the build id of the version")
}
