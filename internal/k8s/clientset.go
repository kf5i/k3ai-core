package k8s

import (
	"path/filepath"

	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	// auth providers
	_ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

// NewClientConfig returns a new Kubernetes client config set for a context
func NewClientConfig(configPath string, contextName string) clientcmd.ClientConfig {
	configPathList := filepath.SplitList(configPath)
	configLoadingRules := &clientcmd.ClientConfigLoadingRules{}
	if len(configPathList) <= 1 {
		configLoadingRules.ExplicitPath = configPath
	} else {
		configLoadingRules.Precedence = configPathList
	}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		configLoadingRules,
		&clientcmd.ConfigOverrides{
			CurrentContext: contextName,
		},
	)
}

// NewClientSet returns a new Kubernetes client for a client config
func NewClientSet(clientConfig clientcmd.ClientConfig) (*kubernetes.Clientset, error) {
	c, err := clientConfig.ClientConfig()

	if err != nil {
		return nil, errors.Wrap(err, "failed to get client config")
	}

	clientset, err := kubernetes.NewForConfig(c)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create clientset")
	}

	return clientset, nil
}

// GetK3sClientSet returns the kubernetes Clientset
func GetK3sClientSet() (*kubernetes.Clientset, error) {
	clientConfig := NewClientConfig(GetK3sKubeConfig(), "default")
	return NewClientSet(clientConfig)

}

// GetK0sClientSet returns the kubernetes Clientset
func GetK0sClientSet() (*kubernetes.Clientset, error) {
	clientConfig := NewClientConfig(GetK0sKubeConfig(), "default")
	return NewClientSet(clientConfig)

}
