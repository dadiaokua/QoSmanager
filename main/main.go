package main

import (
	"context"
	"flag"
	"fmt"
	"time"
)

func main() {
	var kubeconfig *string
	kubeconfig = flag.String("kubeconfig", "~/.kube/config", "absolute path to the kubeconfig file")
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		// Examples for error handling:
		// - Use helper functions like e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		namespace := "default"
		pod := "example-xxxxx"
		if _, err = clientset.CoreV1().Pods(namespace).Get(context.TODO(), pod, metav1.GetOptions{}); err != nil {

			if errors.IsNotFound(err) {
				fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
			}

		} else {
			fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
		}

		time.Sleep(5 * time.Second)
	}
}
