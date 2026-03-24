package organizer

import (
	"fmt"
	"os"
	"strings"
)

func OrganizeFiles(path string) {
	files, err := os.ReadDir(path)

	if err != nil {
		fmt.Println("Error reading path : ", err)
	}

	for _, file := range files {
		fileNameParts := strings.Split(file.Name(), ".")
		fileType := fileNameParts[len(fileNameParts)-1]
		if !file.IsDir() {
			oldPath := path + "/" + file.Name()
			newPath := path + "/" + fileType + "s/" + file.Name()

			os.Chdir(path)

			err := os.MkdirAll(fileType+"s", os.ModePerm)
			if err != nil {
				fmt.Println("Error creating directory:", err)
				return
			}

			err = os.Rename(oldPath, newPath)
			if err != nil {
				fmt.Println("Error moving file:", err)
				return
			}

		}

	}
}
