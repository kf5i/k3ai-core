package plugins

import (
	"fmt"
	"testing"
)

func TestValidatePluginsGroupSpec(t *testing.T) {
	testPluginsGroupsSpec, err := LoadPluginsGroupSpecFormFile("testdata/plugins_group/standard_two_plugins/plugins_group.yaml")

	if err != nil {
		t.Fatalf("failed to unmarshal test file: %s", err)
	}
	var tests = []PluginsGroupSpec{
		PluginsGroupSpec{},
		*testPluginsGroupsSpec,
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
			err = test.validatePluginsGroupSpec()
			if err != nil {
				t.Fatalf("expected nil but got %v", err)
			}
		})
	}
}
