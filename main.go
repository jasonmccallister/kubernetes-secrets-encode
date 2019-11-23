package main

import (
	"encoding/base64"
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// YAML is the representation of a secrets yaml
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
	file, err := ioutil.ReadFile(fileArg)
	if err != nil {
		log.Fatal(err)
	}

	output := flag.String("output", "encoded.yaml", "the output file to save encoded YAML file")
	flag.Parse()

	secrets := YAML{}

	yaml.Unmarshal(file, &secrets)

	if secrets.Kind != "Secret" {
		log.Fatal(errors.New("you must provide a secrets yaml"))
	}

	err = yaml.Unmarshal(file, &secrets)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for k, v := range secrets.Data {
		secrets.Data[k] = base64.StdEncoding.EncodeToString([]byte(v))
	}

	// write to output file
	encoded, err := yaml.Marshal(&secrets)
	if err != nil {
		log.Fatal(err)
	}

	// write to file
	err = ioutil.WriteFile(*output, []byte(encoded), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
