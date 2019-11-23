package main

import (
	"flag"
	"log"
	"os"

	"github.com/jasonmccallister/secretsyaml/secretsyaml"
)

// YAML is the representation of a secrets YAML
type YAML struct {
	APIVersion string            `yaml:"apiVersion"`
	Kind       string            `yaml:"kind"`
	Metadata   map[string]string `yaml:"metadata"`
	Data       map[string]string `yaml:"data"`
}

func main() {
	args := os.Args
	if len(args) == 1 {
		log.Fatal("missing file arg")
	}
	input := args[1]

	var output string
	flag.StringVar(&output, "output", "encoded.yaml", "the output file to save encoded YAML file")

	flag.Parse()

	err := secretsyaml.Encode(input, output)
	if err != nil {
		log.Fatal(err)
	}
}
