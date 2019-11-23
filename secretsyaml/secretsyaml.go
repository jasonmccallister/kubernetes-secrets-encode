package secretsyaml

import "errors"

// Encode takes a file and output as a string
// it will open the file, parse as YAML
// and return a struct representation
// of the file provided with the secrets
// data base64 encoded
func Encode(file, output string) error {
	return errors.New("not yet implemented")
}
