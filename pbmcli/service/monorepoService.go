package service

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
)

const monorepoFolderName = ".pbm"

func getMonorepo() (*Monorepo, error) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	if !checkIfProjectExists(&dir) {
		log.Println("Project does not exist in this directory.")
	}

	yamlFile, err := os.ReadFile(filepath.Join(monorepoFolderName, "monorepo.yaml"))
	if err != nil {
		log.Fatalf("Error reading monorepo.yaml file: %v", err)
		return nil, err
	}

	var monorepo Monorepo
	err = yaml.Unmarshal(yamlFile, &monorepo)
	if err != nil {
		log.Fatal(err)
	}

	return &monorepo, nil
}

func monorepoFolderExist(dir *string) bool {
	dirs, err := os.ReadDir(*dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range dirs {
		if file.IsDir() {
			if file.Name() == monorepoFolderName {
				return true
			}
		}
	}
	return false
}

type Monorepo struct {
	version     string
	projectName string
	projectUUID string
}
