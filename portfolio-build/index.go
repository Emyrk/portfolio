package main

import (
	"html/template"
	"path/filepath"
	"sort"

	"io/ioutil"
	"os"

	"strings"

	"gopkg.in/yaml.v2"
)

var _ = yaml.Decoder{}

type IndexPage struct {
	Projects         []Project
	ProjectHierarchy [][]Project
	PageMetaData     PageMetaData
	TagList          []string
	TagListSet       map[string]string
}

func (i *IndexPage) ConstructHierarchy() {
	counter := 0
	var arr []Project
	for _, p := range i.Projects {
		if counter >= 3 {
			i.ProjectHierarchy = append(i.ProjectHierarchy, arr)
			arr = []Project{p}
			counter = p.Size
			continue
		}
		arr = append(arr, p)
		counter += p.Size
	}
	if len(arr) > 0 {
		i.ProjectHierarchy = append(i.ProjectHierarchy, arr)
	}
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
	Order       int      `yaml:"order"`
	Size        int      `yaml:"size"`
}

type PageMetaData struct {
	Title string `yaml:"web-title"`
	//Description string `yaml:""`
	Headers []ProjectHeader `yaml:"header"`
}

type ProjectHeader struct {
	Text  string `yaml:"text"`
	Href  string `yaml:"href"`
	Order int    `yaml:"order"`
}

func BuildState() (*IndexPage, error) {
	state := NewIndexPage()
	projects := make([]Project, 0)

	// Compile Meta Data config
	file, err := os.OpenFile("_config/index.yml", os.O_RDONLY, 0777)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &state.PageMetaData)
	if err != nil {
		return nil, err
	}

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

		if proj.Size == 0 {
			proj.Size = 1
		}

		projects = append(projects, proj)
	}

	state.Projects = projects

	// Sort sortable objects
	//	Sort Headers
	sort.SliceStable(state.PageMetaData.Headers, func(i, j int) bool {
		return state.PageMetaData.Headers[i].Order < state.PageMetaData.Headers[j].Order
	})
	// 	Sort Projects
	sort.SliceStable(state.Projects, func(i, j int) bool {
		return state.Projects[i].Order < state.Projects[j].Order
	})

	// Build Hierarchy
	state.ConstructHierarchy()

	return state, nil
}

func BuildIndexPage(outfile string) error {
	state, err := BuildState()
	if err != nil {
		return err
	}

	out, err := os.OpenFile(outfile, os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0777)
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
	"uLine": func(s string) string {
		return strings.Replace(s, " ", "_", -1)
	},
	"add": func(a, b int) int {
		return a + b
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
