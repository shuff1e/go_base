package main

import (
	"context"
	"flag"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

func main() {
	h := os.Getenv("HOME")
	f := filepath.Join(h, ".kube", "config")
	kubeconfig := flag.String("kubeconfig", f, "(optional) absolute path to the kubeconfig file")

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("asaqweqwe:", err)
	}

	deploys,err := clientset.AppsV1().Deployments("sre").Get(context.TODO(),"wakanda",v1.GetOptions{})
	if err != nil {
		fmt.Println("get deploy list err:", err)
	}

	fmt.Println("kind::",deploys.GetObjectKind().GroupVersionKind())
	fmt.Printf("one:%#v\n",deploys)
}
