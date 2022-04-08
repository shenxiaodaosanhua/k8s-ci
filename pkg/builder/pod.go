package builder

import (
	"context"
	"fmt"
	"github.com/shenxiaodaosanhua/k8s-ci/pkg/apis/task/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type PodBuilder struct {
	task *v1alpha1.Task
	client.Client
}

func NewPodBuilder(task *v1alpha1.Task, client client.Client) *PodBuilder {
	return &PodBuilder{task: task, Client: client}
}

func (r *PodBuilder) Builder(ctx context.Context) error {
	pod := &corev1.Pod{}
	pod.Namespace = r.task.Namespace
	pod.Name = fmt.Sprintf("task-pod-%s", r.task.Name)
	//重不重启
	pod.Spec.RestartPolicy = corev1.RestartPolicyNever

	c := []corev1.Container{}
	for _, step := range r.task.Spec.Steps {
		step.Container.ImagePullPolicy = corev1.PullIfNotPresent
		c = append(c, step.Container)
	}
	pod.Spec.Containers = c
	pod.OwnerReferences = append(pod.OwnerReferences, metav1.OwnerReference{
		APIVersion: r.task.APIVersion,
		Kind:       r.task.Kind,
		Name:       r.task.Name,
		UID:        r.task.UID,
	})

	return r.Create(ctx, pod)
}
