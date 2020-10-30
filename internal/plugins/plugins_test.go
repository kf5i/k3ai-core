package plugins

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestValidate(t *testing.T) {
	var p Plugin
	testPluginSpec, err := p.Encode("testdata/defaults/plugin.yaml")

	if err != nil {
		t.Fatal("failed to unmarshal test file")
	}
	var tests = []Plugin{
		Plugin{Namespace: "default"},
		*testPluginSpec,
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
			err = test.validate()
			if err != nil {
				t.Fatalf("expected nil but got %v", err)
			}
		})
	}
}

func TestValidateDefaultValues(t *testing.T) {
	var p Plugin
	testPluginSpec, err := p.Encode("testdata/empty_defaults/plugin.yaml")
	if err != nil {
		t.Fatal("failed to unmarshal test file")
	}
	var tests = []Plugin{
		Plugin{Namespace: "default"},
		*testPluginSpec,
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
			err = test.validate()
			if err != nil {
				t.Fatalf("expected nil but got %v", err)
			}
		})
	}
}

func getTestSpecFile(t *testing.T, filePath string) []byte {
	var file, err = ioutil.ReadFile("testdata/" + filePath)
	if err != nil {
		t.Fatal("failed to setup the test")
	}
	return file
}
