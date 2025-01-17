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

	masterUrl := ""
	config, err := clientcmd.BuildConfigFromFlags(masterUrl, *kubeconfig)
	if err != nil {
		panic(err)
	}


	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("asaqweqwe:", err)
	}

	//clientset.CoreV1().Pods("default").List()

	pods, err := clientset.CoreV1().Pods("sre").Get(context.TODO(),"wakanda-84f6974fcc-5qr4b" ,v1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%#v\n",pods.Kind)


	//for _, v := range pods.Items {
	//	fmt.Println(v.Name)
	//
	//	fmt.Println(v.CreationTimestamp)
	//
	//	//v.Name
	//}
	//fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
}
