package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	filesDir = "files"
)

func main() {
	filesDirFd, err := os.Open(filesDir)
	if err != nil {
		panic(err)
	}

	fileInfoList, err := filesDirFd.Readdir(0)
	if err != nil {
		panic(err)
	}

	filenames := make([]string, 0, len(fileInfoList))
	for _, fileInfo := range fileInfoList {
		filenames = append(filenames, fileInfo.Name())
	}

	sourceToDestination := make(map[string]string, len(filenames))
	for _, filename := range filenames {
		sourceFilename := filepath.Join(filesDir, filename)
		destinationFilename := filepath.Join("..", "." + filename)
		sourceToDestination[sourceFilename] = destinationFilename
	}

	for source, destination := range sourceToDestination {
		fmt.Printf("linking %v -> %v\n", source, destination)
		err := os.Symlink(source, destination)
		if err != nil {
			panic(err)
		}
	}
}
