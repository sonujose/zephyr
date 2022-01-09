package kubernetes

import (
	"flag"
	"log"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	//Vendor specific authentication (Azure)
	_ "k8s.io/client-go/plugin/pkg/client/auth/azure"
)

// Create new kubernetes client based on the runtime
// Create kubernetes client while running locally in machine using kubeconfig for incluster use service account
// Use vendor specific sdk for auth if you are not using cluster certificates in your kubeconfig
// Here we are using k8s.io/client-go/plugin/pkg/client/auth/azure since the connection is only from aks
func NewClient() (client *kubernetes.Clientset, err error) {

	var kubeconfig *string

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		// handle error
		log.Printf("Unable to build kubernetes client config from flags - %s, using in-cluster service account.\n", err.Error())
		config, err = rest.InClusterConfig()

		if err != nil {
			log.Printf("Kubernetes client Error fetching getting inclusterconfig - %s", err.Error())
			return client, err
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		// handle error
		log.Printf("Error creating clientset %v \n", err.Error())
	}

	return clientset, nil
}
