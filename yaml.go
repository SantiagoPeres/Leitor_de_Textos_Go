package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Metadata struct {
	Name string   `yaml:"name"`
	Tags []string `yaml:"tags"`
}

func main() {
	data, err := os.ReadFile("data/metadata.yaml")
	if err != nil {
		panic(err)
	}

	var metadata Metadata
	err = yaml.Unmarshal(data, &metadata)
	if err != nil {
		panic(err)
	}

	metadata.Name += "!!!"
	metadata.Tags = append(metadata.Tags, "minha tag")

	ndata, err := yaml.Marshal(metadata)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("data/metadata.yaml", ndata, 0755)
	if err != nil {
		panic(err)
	}
}
