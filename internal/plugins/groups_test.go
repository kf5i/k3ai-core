package plugins

import (
	"fmt"
	"testing"
)

func TestValidatePluginsGroupSpec(t *testing.T) {
	var group Group
	 err := group.Encode(joinWithRootData("groups/argo-workflow/group.yaml"))

	if err != nil {
		t.Fatalf("failed to unmarshal test file: %s", err)
	}
	var tests = []Group{
		group,
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
