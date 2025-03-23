package main

import (
	"context"
	cloudnative "controller/internal/clientset/versioned"
	cloudnativeInformer "controller/internal/informers/externalversions"
	"controller/pkg/controller/controller"
	"flag"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const defaultSyncTime = time.Second * 30

var (
	kubeconfig string
	threads    int
)

func parseFlags() {
	flag.StringVar(&kubeconfig, "kubeconfig", "~/.kube/config", "")
	flag.IntVar(&threads, "threads", 2, "")
	flag.Parse()
}

func restConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
		return cfg, nil
	}

	cfg, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func main() {
	parseFlags()

	k8scfg, err := restConfig(kubeconfig)
	if err != nil {
		klog.Fatalf("Error to build rest config: %s", err.Error())
	}

	clientset, err := cloudnative.NewForConfig(k8scfg)
	if err != nil {
		klog.Fatalf("Error to build cloudnative clientset: %s", err.Error())
	}

	informer := cloudnativeInformer.NewSharedInformerFactory(clientset, defaultSyncTime)
	controller := controller.New(clientset, informer)
	ctx, cancel := context.WithCancel(context.Background())
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	if err := controller.Run(ctx, threads); err != nil {
		klog.Fatalf("Error to run the controller instance: %s.", err)
	}

	<-signalChan
	cancel()
	controller.Stop()
}
