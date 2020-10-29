package plugins

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestValidate(t *testing.T) {
	file := getTestSpecFile(t, "defaults/plugin.yaml")
	testPluginSpec, err := unmarshalPluginSpec(file)

	if err != nil {
		t.Fatal("failed to unmarshal test file")
	}
	var tests = []PluginSpec{
		PluginSpec{Namespace: "default"},
		*testPluginSpec,
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
			err = test.validatePluginSpec()
			if err != nil {
				t.Fatalf("expected nil but got %v", err)
			}
		})
	}
}

func TestValidateDefaultValues(t *testing.T) {
	file := getTestSpecFile(t, "empty_defaults/plugin.yaml")
	testPluginSpec, err := unmarshalPluginSpec(file)
	if err != nil {
		t.Fatal("failed to unmarshal test file")
	}
	var tests = []PluginSpec{
		PluginSpec{Namespace: "default"},
		*testPluginSpec,
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
			err = test.validatePluginSpec()
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
