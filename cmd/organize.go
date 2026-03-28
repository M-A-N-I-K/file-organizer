package organizer

import (
	organizer "file-organizer/pkg"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var excludes []string

var organizeCmd = &cobra.Command{
	Use:     "organize",
	Aliases: []string{"organize"},
	Short:   "Organize Folder",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()
		excludes, excludesErr := cmd.Flags().GetStringArray("exclude")
		log, logErr := cmd.Flags().GetBool("log")
		if excludesErr != nil || logErr != nil {
			fmt.Println("Wrong exclude flag format:", logErr, excludesErr)
			return
		}

		organizer.OrganizeFiles(args[0], excludes, log)

		elapsed := time.Since(start)
		fmt.Printf("Organizing files took %s ", elapsed)
	},
}

func init() {
	organizeCmd.Flags().StringArrayVar(&excludes, "exclude", []string{}, "Files to exclude")
	organizeCmd.Flags().Bool("log", false, "Log The Results")
	rootCmd.AddCommand(organizeCmd)
}
