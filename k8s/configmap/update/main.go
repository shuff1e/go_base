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
	"strconv"
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

	cm,err := clientset.CoreV1().ConfigMaps("test").Get(context.TODO(),"test-cm",v1.GetOptions{})
	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Println("old:",cm.Data["age"])

	age := cm.Data["age"]
	oldAge,_ := strconv.Atoi(age)

	cm.Data["age"] = strconv.Itoa(oldAge+1)

	_,err =clientset.CoreV1().ConfigMaps("test").Update(context.TODO(),cm,v1.UpdateOptions{})
	if err != nil {
		fmt.Println("err:", err)
	}

}
