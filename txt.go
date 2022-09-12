package main

import (
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("data/aprenda.txt")
	if err != nil {
		panic(err)
	}

	txt := strings.Replace(string(data), "bacana", "legal", 1)

	err = os.WriteFile("data/aprenda.txt", []byte(txt), 0755)
	if err != nil {
		panic(err)
	}

}
