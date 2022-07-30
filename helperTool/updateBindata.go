package main

import (
	"fmt"
	"gopkg.in/errgo.v2/errors"
	"os"
	"path/filepath"

	"github.com/kevinburke/go-bindata"
)

type UpdateBindata struct {
	Directory string `long:"directory" description:"source code root directory" default:"../"`
}

var uBindataCommand UpdateBindata

func init() {
	_, err := parser.AddCommand(
		"updateBindata",
		"Create/Update Bindata",
		"Create/Update the Bindata (migrations and templates)",
		&uBindataCommand)

	if err != nil {
		panic(err)
	}
}

func (x *UpdateBindata) Execute(args []string) error {
	if len(x.Directory) == 0 {
		return errors.New("Need the project's root directory")
	}

	templatesDirectory := filepath.Join(x.Directory, "httpServer", "templates")
	stat, err := os.Stat(templatesDirectory)
	if err != nil {
		return errors.New(fmt.Sprintf("templates directory '%s' doesn't exist", templatesDirectory))
	} else if !stat.IsDir() {
		return errors.New(fmt.Sprintf("templates directory '%s' isn't a directory", templatesDirectory))
	}

	templatesBinDirectory := filepath.Join(x.Directory, "httpServer", "templateData")
	stat, err = os.Stat(templatesDirectory)
	if err != nil {
		return errors.New(fmt.Sprintf("templates bin directory '%s' doesn't exist", templatesBinDirectory))
	} else if !stat.IsDir() {
		return errors.New(fmt.Sprintf("templates bin directory '%s' isn't a directory", templatesBinDirectory))
	}

	config := bindata.Config{
		Package: "templateData",
		Input:   []bindata.InputConfig{bindata.InputConfig{Path: templatesDirectory, Recursive: true}},
		Output:  filepath.Join(templatesBinDirectory, "main.go"),
		Prefix:  templatesDirectory,
	}

	err = bindata.Translate(&config)

	if err != nil {
		return errors.Because(err, nil, "Unable to create template bindata")
	}

	return nil
}
