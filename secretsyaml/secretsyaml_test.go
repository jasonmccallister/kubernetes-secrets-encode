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

			outputFile, err := ioutil.ReadFile(tC.output)
			if err != nil {
				t.Fatal(err)
			}

			// Assert
			if strings.Compare(string(inputFile), string(outputFile)) == 0 {
				t.Errorf(
					"expected the input and output files to not match, expected: \n%v\n actual: \n%v\n",
					string(inputFile),
					string(outputFile),
				)
			}
			if tC.err != "" {
				if err.Error() != tC.err {
					t.Errorf("expected error to be %v, got %v instead", tC.err, err)
				}
			}
		})
	}
}
