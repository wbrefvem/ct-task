package main

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	client := clientset.ExtensionsV1beta1().Deployments("default")

	for {

		deployments, err := client.List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}

		for _, dep := range deployments.Items {
			if dep.ObjectMeta.Labels["app"] == "nginx" {
				minuteDigit := int32(time.Now().Minute() % 10)
				dep.Spec.Replicas = &minuteDigit

				_, err := client.Update(&dep)
				if err != nil {
					panic(err.Error())
				}
			}

		}

		time.Sleep(1 * time.Second)
	}
}
