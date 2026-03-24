package organizer

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "organize",
	Short: "Organize - A lightweight Go-based CLI tool to organize files in a directory",
	Long:  `Organize is a lightweight Go-based CLI tool that organizes files in a directory based on their extensions. It moves files to subdirectories named after their extensions, helping developers maintain a clean and organized file system.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
