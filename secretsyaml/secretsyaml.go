package secretsyaml

import (
	"encoding/base64"
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Secret is the representation of a secrets.yaml file in Go
// used for easy marshalling and unmarshalling
type Secret struct {
	APIVersion string            `yaml:"apiVersion"`
	Kind       string            `yaml:"kind"`
	Metadata   map[string]string `yaml:"metadata"`
	Data       map[string]string `yaml:"data"`
}

// Encode takes a filename and output as a string
// it will open the file, parse as YAML
// and return a struct representation
// of the file provided with the secrets
// data base64 encoded
func Encode(filename, output string) error {
	// open the provided file
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	// unmarshal file
	secret := Secret{}
	err = yaml.Unmarshal(file, &secret)
	if err != nil {
		return err
	}

	// make sure this is a secret
	if secret.Kind != "Secret" {
		return errors.New("you must provide a YAML with the kind Secret")
	}

	// encode each value in the .Data map, placing the encoded value
	// back into the struct
	for k, v := range secret.Data {
		secret.Data[k] = base64.StdEncoding.EncodeToString([]byte(v))
	}

	// marshal into struct
	encoded, err := yaml.Marshal(&secret)
	if err != nil {
		return err
	}

	// write the struct to the output file
	err = ioutil.WriteFile(output, []byte(encoded), 0644)
	if err != nil {
		return err
	}

	return nil
}
