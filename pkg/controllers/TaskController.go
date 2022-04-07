package controllers

import (
	"context"
	"fmt"
	clientset "github.com/shenxiaodaosanhua/k8s-ci/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type TaskController struct {
	client.Client
	*clientset.Clientset
	E record.EventRecorder
}

func NewTaskController(e record.EventRecorder, clientset *clientset.Clientset) *TaskController {
	return &TaskController{E: e, Clientset: clientset}
}

func (c *TaskController) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	task, err := c.ApiV1alpha1().Tasks(req.Namespace).Get(ctx, req.Name, metav1.GetOptions{})
	if err != nil {
		return reconcile.Result{}, err
	}
	fmt.Println(task.Name)
	return reconcile.Result{}, nil
}

// InjectClient 注入client
func (c *TaskController) InjectClient(client client.Client) error {
	c.Client = client
	return nil
}
