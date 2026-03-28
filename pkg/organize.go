package organizer

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func OrganizeFiles(path string, excludes []string, log bool) {
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

			err := os.MkdirAll(fileType+"s", os.ModePerm)
			if err != nil {
				fmt.Println("Error creating directory:", err)
				return
			}

			err = os.Rename(oldPath, newPath)

			if log {
				fmt.Printf("Moved file from %s to %s\n", oldPath, newPath)
			}
			if err != nil {
				fmt.Println("Error moving file:", err)
				return
			}
		} else {
			OrganizeFiles(path+"/"+fileName, excludes, log)
		}

	}
}
