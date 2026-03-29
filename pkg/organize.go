package organizer

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"text/tabwriter"
)

func OrganizeFiles(path string, w *tabwriter.Writer, excludes []string, log bool, dryRun bool) {
	files, err := os.ReadDir(path)

	if err != nil {
		fmt.Println("Error reading path : ", err)
	}

	fileNameParts := strings.Split(path, "/")
	folderName := fileNameParts[len(fileNameParts)-1]

	for _, file := range files {
		fileName := file.Name()
		if slices.Contains(excludes, fileName) {
			if log {
				fmt.Println("Excluded file : ", fileName)
			}
			return
		}

		fileNameParts := strings.Split(fileName, ".")
		fileType := fileNameParts[len(fileNameParts)-1]

		if fileType+"s" == folderName {
			return
		}

		if !file.IsDir() {
			oldPath := path + "/" + fileName
			newPath := path + "/" + fileType + "s/" + fileName

			os.Chdir(path)

			if !dryRun {
				err := os.MkdirAll(fileType+"s", os.ModePerm)
				if err != nil {
					fmt.Println("Error creating directory:", err)
					return
				}
			}

			if !dryRun {
				err = os.Rename(oldPath, newPath)
				if err != nil {
					fmt.Println("Error moving file:", err)
					return
				}
			}

			if log || dryRun {
				fmt.Fprintf(w, "%s\t%s\t%s\n", "MOVE", oldPath, newPath)
			}
		} else {
			OrganizeFiles(path+"/"+fileName, w, excludes, log, dryRun)
		}

	}
}
