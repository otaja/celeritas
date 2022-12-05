package main

import (
	"embed"
	"errors"
	"os"
)

//go:embed templates
var templateFS embed.FS

func copyFilefromTemplate(templatePath, targetFile string) error {
	if fileExist(targetFile) {
		return errors.New(targetFile + "already exist")
	}

	data, err := templateFS.ReadFile(templatePath)
	if err != nil {
		exitGracefully(err)
	}

	err = copyDataToFile(data, targetFile)
	if err != nil {
		exitGracefully(err)
	}

	return nil
}

func copyDataToFile(data []byte, to string) error {
	err := os.WriteFile(to, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func fileExist(fileToCheck string) bool {
	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		return false
	}
	return true
}
