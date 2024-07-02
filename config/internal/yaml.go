package internal

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func RunYAML() {
	file, err := os.ReadFile("config/config.yaml")
	if err != nil {
		panic(err)
	}

	var cfg Config
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", cfg)
}
