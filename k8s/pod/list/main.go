package main

import (
	"context"
	"flag"
	"fmt"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"time"
)

func main() {
	t:= time.Now()

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
		fmt.Println("kubernetes.NewForConfig err:", err)
	}

	list := v12.ListOptions{}
	pl,err := clientset.CoreV1().Pods("").List(context.Background(),list)
	if err  != nil {
		fmt.Println(err)
		return
	}

	for _,v := range pl.Items{
		fmt.Println("name::",v.Name)
		//fmt.Println("kind::",v.Kind)
		//fmt.Println("version::",v.APIVersion)
	}

	fmt.Println(len(pl.Items))

	fmt.Println("time:",time.Now().Sub(t))

	test("",checkOk)

}

type Pod struct {

}

func checkOk(p *Pod) bool {
	return true
}

type podFilter func(pod *Pod) bool

func test(s string, filter ...podFilter) {

}
