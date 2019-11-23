package main

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type secretsYAML struct {
	APIVersion string            `yaml:"apiVersion"`
	Kind       string            `yaml:"kind"`
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

	secrets := secretsYAML{}

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

	log.Println("showing encoded values")
	for _, v := range secrets.Data {
		log.Println(v)
	}

	// write to output file
	encoded, err := yaml.Marshal(&secrets)
	if err != nil {
		log.Fatal(err)
	}

	// write to file
	// TODO convert to dynamic named file
	err = ioutil.WriteFile("./testdata/example.secrets.encoded.yaml", []byte(encoded), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
