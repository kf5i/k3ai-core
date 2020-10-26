package k8s

import (
	"sync"

	"k8s.io/kubectl/pkg/util/openapi"
)

type factory struct {
	KubeConfig            string
	Context               string
	initOpenAPIGetterOnce sync.Once
	openAPIGetter         openapi.Getter
}

// GetKubeConfig returns the kubeconfig path
func GetKubeConfig() string {
	return "/etc/rancher/k3s/k3s.yaml"
}
