package k8s

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func NewClient() (*kubernetes.Clientset, error) {
	pathOptions := clientcmd.NewDefaultPathOptions()
	pathOptions.LoadingRules.DoNotResolvePaths = false
	config, err := pathOptions.GetStartingConfig()
	if err != nil {
		return nil, err
	}

	configOverrides := clientcmd.ConfigOverrides{}
	clientConfig := clientcmd.NewDefaultClientConfig(*config, &configOverrides)
	restConfig, err := clientConfig.ClientConfig()
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}
