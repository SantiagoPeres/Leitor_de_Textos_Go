package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Addr     string `json:"addr"`
	Protocol string `json:"protocol"`
	Type     string `json:"type"`
}

func main() {
	data, err := os.ReadFile("data/config.json")
	if err != nil {
		panic(err)
	}

	var configs []Config
	err = json.Unmarshal(data, &configs)
	if err != nil {
		panic(err)
	}

	configs = append(configs, Config{"canaistech.com.br", "https", "website"})

	ndata, err := json.Marshal(configs)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("data/config.json", ndata, 0755)
	if err != nil {
		panic(err)
	}
}
