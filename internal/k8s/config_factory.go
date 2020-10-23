package k8s

import (
	"k8s.io/kubectl/pkg/util/openapi"
	"sync"
)

type factory struct {
	KubeConfig            string
	Context               string
	initOpenAPIGetterOnce sync.Once
	openAPIGetter         openapi.Getter
}

func GetKubeConfig() string {
	return "/etc/rancher/k3s/k3s.yaml"
}
