package plugins

import (
	"fmt"
	"testing"
)

func TestValidate(t *testing.T) {
	var p Plugin
	err := p.Encode(joinWithRootData("plugins/argo-workflow-ns/plugin.yaml"))

	if err != nil {
		t.Fatal("failed to unmarshal test file")
	}
	var tests = []Plugin{
		{Namespace: "default"},
		p,
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
	err := p.Encode(joinWithRootData("plugins/argo-workflow-no-defaults/plugin.yaml"))
	if err != nil {
		t.Fatal("failed to unmarshal test file")
	}
	var tests = []Plugin{
		{Namespace: "default"},
		p,
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
