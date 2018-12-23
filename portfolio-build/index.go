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
	TagList      []string
	TagListSet   map[string]string
}

func NewIndexPage() *IndexPage {
	i := new(IndexPage)
	i.TagListSet = make(map[string]string)

	return i
}

type Project struct {
	Title       string   `yaml:"title"`
	Description string   `yaml:"description"`
	Tags        []string `yaml:"tags"`
	Background  string   `yaml:"background"`
}

type PageMetaData struct {
	Tile        string
	Description string
}

func BuildState() (*IndexPage, error) {
	state := NewIndexPage()
	projects := make([]Project, 0)

	// Compile projects
	projectYmls, err := filepath.Glob("_projects/*.yml")
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

		for _, tag := range proj.Tags {
			if _, ok := state.TagListSet[tag]; !ok {
				state.TagListSet[tag] = tag
				state.TagList = append(state.TagList, tag)
			}
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

// Custom funcMap for building the index page
var funcMap = template.FuncMap{
	"modN": func(n, v int) int {
		return v % n
	},
	"test": func(item interface{}) (int, error) {
		return 0, nil
	},
	"safeCSS": func(s string) template.CSS {
		return template.CSS(s)
	},
	"safeJS": func(s string) template.JS {
		return template.JS(s)
	},
	"safeHTML": func(s string) template.HTML {
		return template.HTML(s)
	},
}

func parseTemplates(tmplsLoc string) (*template.Template, error) {
	tmpl := template.Must(template.New("index.html").Funcs(funcMap).ParseGlob(tmplsLoc))

	// Include the main templates
	//tmpl, err := template.ParseGlob(tmplsLoc)
	//if err != nil {
	//	return nil, err
	//}

	return tmpl, nil
}
