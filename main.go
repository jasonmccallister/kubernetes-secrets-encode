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
		log.Fatal("missing")
	}

	fileArg := args[1]
	outputArg := flag.String("output", "encoded.yaml", "the output file to save encoded YAML file")
	flag.Parse()

	err := secretsyaml.Encode(fileArg, *outputArg)
	if err != nil {
		log.Fatal(err)
	}
}
