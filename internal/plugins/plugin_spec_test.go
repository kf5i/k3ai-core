package plugins

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestValidate(t *testing.T) {
	file := getTestSpecFile(t)
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
	file := getTestSpecFileWithoutDefaults(t)
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

func getTestSpecFileWithoutDefaults(t *testing.T) []byte {
	var file, err = ioutil.ReadFile("testdata/test_plugin_empty_defaults.yaml")
	if err != nil {
		t.Fatal("failed to setup the test")
	}
	return file
}

func getTestSpecFile(t *testing.T) []byte {
	var file, err = ioutil.ReadFile("testdata/test_plugin.yaml")
	if err != nil {
		t.Fatal("failed to setup the test")
	}
	return file
}
