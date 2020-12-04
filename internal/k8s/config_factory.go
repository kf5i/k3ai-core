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

// GetK3sKubeConfig returns the kubeconfig path
func GetK3sKubeConfig() string {
	return "/etc/rancher/k3s/k3s.yaml"
}

// GetK0sKubeConfig returns the kubeconfig path
func GetK0sKubeConfig() string {
	return "/var/lib/k0s/pki/admin.conf"
}
