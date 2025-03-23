package controller

import (
	"context"
	cloudnative "controller/internal/clientset/versioned"
	cloudnativeInformer "controller/internal/informers/externalversions"
	listerv1 "controller/internal/listers/cloudnative.group/v1"
	"encoding/json"
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
	"time"
)

const (
	resourceName = "VirtualMachine"
)

type Controller struct {
	clientset cloudnative.Interface
	informer  cloudnativeInformer.SharedInformerFactory
	lister    listerv1.VirtualMachineLister
	synced    cache.InformerSynced
	queue     workqueue.TypedRateLimitingInterface[any]
}

func New(clientset cloudnative.Interface, informer cloudnativeInformer.SharedInformerFactory) *Controller {
	vmInformer := informer.Cloudnative().V1().VirtualMachines()
	controller := &Controller{
		clientset: clientset,
		informer:  informer,
		lister:    vmInformer.Lister(),
		synced:    vmInformer.Informer().HasSynced,
		queue:     workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), resourceName),
	}

	vmInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.enqueue,
		UpdateFunc: func(old, new interface{}) {
			controller.enqueue(new)
		},
	})

	return controller
}

func (c *Controller) Run(ctx context.Context, threadiness int) error {
	go c.informer.Start(ctx.Done())
	klog.Info("Starting controller")
	klog.Info("Waiting for informer caches to sync")
	if ok := cache.WaitForCacheSync(ctx.Done(), c.synced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	for i := 0; i < threadiness; i++ {
		go wait.Until(c.runWorker, time.Second, ctx.Done())
	}
	klog.Info("Started workers")
	return nil
}

func (c *Controller) Stop() {
	klog.Info("Stopping controller")
	c.queue.ShutDown()
}

func (c *Controller) runWorker() {
	defer utilruntime.HandleCrash()
	for c.processNextWorkItem() {

	}
}

// 取数据处理
func (c *Controller) processNextWorkItem() bool {
	obj, shutdown := c.queue.Get()
	if shutdown {
		return false
	}

	err := func(obj interface{}) error {
		defer c.queue.Done(obj)
		key, ok := obj.(string)
		if !ok {
			c.queue.Forget(obj)
			utilruntime.HandleError(fmt.Errorf("Controller expected string in workqueue but got %#v", obj))
			return nil
		}

		if err := c.syncHandler(key); err != nil {
			c.queue.AddRateLimited(key)
			return fmt.Errorf("Controller error syncing '%s': %s, requeuing", key, err.Error())
		}

		c.queue.Forget(obj)
		klog.Infof("Controller successfully synced '%s'", key)
		return nil
	}(obj)

	if err != nil {
		utilruntime.HandleError(err)
		return true
	}
	return true
}

func (c *Controller) enqueue(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		utilruntime.HandleError(err)
		return
	}
	c.queue.Add(key)
}

func (c *Controller) syncHandler(key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("invalid resource key: %s", key))
		return err
	}

	vm, err := c.lister.VirtualMachines(namespace).Get(name)
	if err != nil {
		if errors.IsNotFound(err) {
			utilruntime.HandleError(fmt.Errorf("virtual machine '%s' in work queue no longer exists", key))
		}
		return err
	}

	data, err := json.Marshal(vm)
	if err != nil {
		return err
	}

	// TODO do sth to create vm
	klog.Infof("Converting virtual machine %s/%s ,object : %s", vm.Namespace, vm.Name, string(data))
	return nil
}
