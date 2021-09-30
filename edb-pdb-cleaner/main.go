package main

import (
	"context"
	"flag"
	"log"
	"path/filepath"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
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
		log.Fatalln(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln(err.Error())
	}

	ns := "harbor"
	name := "harbor-primary"
	s := time.Duration(30 * time.Second)

	for {
		pdb, err := clientset.PolicyV1().PodDisruptionBudgets(ns).Get(context.TODO(), name, v1.GetOptions{})

		if errors.IsNotFound(err) {
			log.Printf("PodDisruptionBudget %s in namespace %s not found\n", name, ns)
		} else if err != nil {
			log.Fatalln(err.Error())
		} else {
			log.Printf("Found PodDisruptionBudget %s in namespace %s\n", pdb.Name, ns)
			log.Println("Deleting it..")
			err = clientset.PolicyV1().PodDisruptionBudgets(ns).Delete(context.TODO(), name, v1.DeleteOptions{})
			if err != nil {
				log.Fatalln(err.Error())
			}
		}
		log.Printf("Sleeping for %s\n", s)
		time.Sleep(s)

	}
}
