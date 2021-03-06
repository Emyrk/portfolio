package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	var (
		eg      = flag.Bool("eg", false, "Output example project yaml")
		outfile = flag.String("out", "test.html", "Output file")
	)

	flag.Parse()

	if *eg {
		egYaml()
		os.Exit(0)
	}

	BuildIndexPage(*outfile)
}

func egYaml() {
	p := new(Project)
	p.Title = "HodlZone"
	p.Description = "Cryptocurrency lending bot service"
	p.Tags = []string{"Algorithms", "Weekend Hack"}
	p.ModalButtons = []ModalButton{{"x", "x", "x"}}

	data, err := yaml.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

/*
	tmpl, err := template.ParseFiles("test.html")
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile("out.html", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(file, nil)
	if err != nil {
		panic(err)
	}
*/
