package main

import (
	"encoding/json"
	"os"

	"gopkg.in/yaml.v2"
)

type Metadata struct {
	Name string   `yaml:"name" json:"name"`
	Tags []string `yaml:"tags" json:"tags"`
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
	metadata.Tags = append(metadata.Tags, "json")

	ndata, err := json.Marshal(metadata)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("data/metadata.json", ndata, 0755)
	if err != nil {
		panic(err)
	}
}
