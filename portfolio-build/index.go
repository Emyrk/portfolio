package main

import (
	"html/template"
	"path/filepath"

	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var _ = yaml.Decoder{}

type IndexPage struct {
	Projects     []Project
	PageMetaData PageMetaData
}

type Project struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
}

type PageMetaData struct {
	Tile        string
	Description string
}

func BuildState() (*IndexPage, error) {
	state := new(IndexPage)
	projects := make([]Project, 0)

	// Compile projects
	projectYmls, err := filepath.Glob("*.yml")
	if err != nil {
		return nil, err
	}

	for _, yml := range projectYmls {
		var proj Project
		file, err := os.OpenFile(yml, os.O_RDONLY, 0777)
		if err != nil {
			return nil, err
		}

		data, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}

		err = yaml.Unmarshal(data, &proj)
		if err != nil {
			return nil, err
		}

		projects = append(projects, proj)
	}

	state.Projects = projects

	return state, nil
}

func BuildIndexPage(outfile string) error {
	state, err := BuildState()
	if err != nil {
		return err
	}

	out, err := os.OpenFile(outfile, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return err
	}

	tmpls, err := parseTemplates("_tmpls/*.html")
	err = tmpls.Execute(out, state)
	if err != nil {
		return err
	}

	return nil
}

func parseTemplates(tmplsLoc string) (*template.Template, error) {
	return template.ParseGlob(tmplsLoc)
}
