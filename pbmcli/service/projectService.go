package service

import (
	"errors"
	"log"
	"os"
	"pbmcli/pbmcli/resources"
	"pbmcli/pbmcli/utils"
)

var defaultVersion = "1"
var defaultDescription = "No description"

func CreateProject(project *Project) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	if checkIfProjectExists(&dir) {
		log.Println("Project already exists in this directory.")
	}
	log.Println("Creating project at ", dir)

	utils.CopyAllFromFs(resources.MonorepoTmpl, "null")
}

func NewProject(name *string, description *string, version *string) (*Project, error) {
	if name == nil {
		return nil, errors.New("project name is required")
	}

	if version == nil {
		version = &defaultVersion
	}
	if description == nil {
		description = &defaultDescription
	}

	return &Project{
		Name:        *name,
		Version:     *version,
		Description: *description,
	}, nil
}

func checkIfProjectExists(dir *string) bool {
	dirs, err := os.ReadDir(*dir)
	if err != nil {
		log.Fatal(err)
	}

	var containPbmFolder = false
	for _, file := range dirs {
		if file.IsDir() {
			if file.Name() == ".pbm" {
				containPbmFolder = true
			}
		}
	}

	return containPbmFolder
}

type Project struct {
	Name        string
	Version     string
	Description string
}
