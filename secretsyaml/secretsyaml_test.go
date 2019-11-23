package secretsyaml

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestEncode(t *testing.T) {
	testCases := []struct {
		desc   string
		input  string
		output string
		err    string
	}{
		{
			desc:   "example and encoded example YAML files match",
			input:  "./testdata/example.secrets.yaml",
			output: "./testdata/encoded.yaml",
			err:    "",
		},
		{
			desc:   "YAML that is not a Secret returns an error",
			input:  "./testdata/not.a.secrets.yaml",
			output: "./testdata/not-a-secret-encoded.yaml",
			err:    "you must provide a YAML with the kind Secret",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// Arrange
			inputFile, err := ioutil.ReadFile(tC.input)
			if err != nil {
				t.Fatal(err)
			}

			// Act
			err = Encode(tC.input, tC.output)

			// Assert
			if tC.err != "" {
				if err.Error() != tC.err {
					t.Errorf("expected error to be %v, got %v instead", tC.err, err)
				}
			} else {
				outputFile, err := ioutil.ReadFile(tC.output)
				if err != nil {
					t.Fatal(err)
				}
				if strings.Compare(string(inputFile), string(outputFile)) == 0 {
					t.Errorf(
						"expected the input and output files to not match, expected: \n%v\n actual: \n%v\n",
						string(inputFile),
						string(outputFile),
					)
				}
			}
		})
	}
}
