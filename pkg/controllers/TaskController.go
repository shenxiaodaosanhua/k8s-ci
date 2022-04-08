package controllers

import (
	"context"
	"github.com/shenxiaodaosanhua/k8s-ci/pkg/apis/task/v1alpha1"
	"github.com/shenxiaodaosanhua/k8s-ci/pkg/builder"
	clientset "github.com/shenxiaodaosanhua/k8s-ci/pkg/client/clientset/versioned"
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
	task := &v1alpha1.Task{}
	err := c.Get(ctx, req.NamespacedName, task)
	if err != nil {
		return reconcile.Result{}, err
	}
	//构建pod
	err = builder.NewPodBuilder(task, c.Client).Builder(ctx)
	if err != nil {
		return reconcile.Result{}, err
	}
	return reconcile.Result{}, nil
}

// InjectClient 注入client
func (c *TaskController) InjectClient(client client.Client) error {
	c.Client = client
	return nil
}
