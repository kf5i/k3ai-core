package plugins

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestValidate(t *testing.T) {
	file := getTestSpecFile(t, "test_plugin.yaml")
	testPluginSpec, err := unmarshal(file)
	if err != nil {
		t.Fatal("failed to unmarshal test file")
	}
	var tests = []PluginSpec{
		PluginSpec{},
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
	file := getTestSpecFile(t, "test_plugin_empty_defaults.yaml")
	testPluginSpec, err := unmarshal(file)
	if err != nil {
		t.Fatal("failed to unmarshal test file")
	}
	var tests = []PluginSpec{
		PluginSpec{},
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

func getTestSpecFile(t *testing.T, filename string) []byte {
	var file, err = ioutil.ReadFile("testdata/" + filename)
	if err != nil {
		t.Fatal("failed to setup the test")
	}
	return file
}
