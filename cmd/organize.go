package organizer

import (
	organizer "file-organizer/pkg"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var organizeCmd = &cobra.Command{
	Use:     "organize",
	Aliases: []string{"organize"},
	Short:   "Organize Folder",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()

		organizer.OrganizeFiles(args[0])

		elapsed := time.Since(start)
		fmt.Printf("Organizing files took %s ", elapsed)
	},
}

func init() {
	rootCmd.AddCommand(organizeCmd)
}
