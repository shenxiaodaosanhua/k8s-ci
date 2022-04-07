package k8sconfig

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
)

func K8sRestConfigInPod() *rest.Config {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err)
	}

	return config
}

func K8sRestConfig() *rest.Config {
	if os.Getenv("release") == "1" {
		log.Println("run in cluster")
		return K8sRestConfigInPod()
	}
	config, err := clientcmd.BuildConfigFromFlags("", "./resource/config")
	if err != nil {
		log.Fatal(err)
	}
	config.Insecure = true
	return config
}
