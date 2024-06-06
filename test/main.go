package main

import (
	"bufio"
	"context"
	"fmt"
	mv1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/hexin/.kube/config")
	if err != nil {
		fmt.Println(err)
		return
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx, _ := context.WithCancel(context.Background())
	pod, err := client.CoreV1().Pods("default").Get(ctx, "log-generator", v1.GetOptions{})
	if err != nil {
		fmt.Println(pod, err)
		return
	}
	podLogOpts := mv1.PodLogOptions{
		Follow: true,
	}
	req := client.CoreV1().Pods("default").GetLogs("log-generator", &podLogOpts)
	s, err := req.Stream(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer s.Close()
	reader := bufio.NewScanner(s)
	for reader.Scan() {
		fmt.Println(reader.Text())
	}
}
