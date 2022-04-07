package k8sconfig

import (
	taskv1alpha1 "github.com/shenxiaodaosanhua/k8s-ci/pkg/apis/task/v1alpha1"
	"github.com/shenxiaodaosanhua/k8s-ci/pkg/client/clientset/versioned"
	"github.com/shenxiaodaosanhua/k8s-ci/pkg/controllers"
	"log"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
)

func IntManger() {
	taskClient := versioned.NewForConfigOrDie(K8sRestConfig())
	logf.SetLogger(zap.New())

	manager, err := manager.New(
		K8sRestConfig(),
		manager.Options{
			Logger: logf.Log.WithName("k8s-ci"),
		},
	)
	if err != nil {
		log.Fatal("创建管理器失败:", err.Error())
	}

	err = taskv1alpha1.SchemeBuilder.AddToScheme(manager.GetScheme())
	if err != nil {
		manager.GetLogger().Error(err, "unable add schema")
		os.Exit(1)
	}

	taskController := controllers.NewTaskController(manager.GetEventRecorderFor("k8s-ci"), taskClient)

	err = builder.ControllerManagedBy(manager).For(&taskv1alpha1.Task{}).Complete(taskController)
	if err != nil {
		manager.GetLogger().Error(err, "unable to create manager")
		os.Exit(1)
	}

	err = manager.Start(signals.SetupSignalHandler())
	if err != nil {
		manager.GetLogger().Error(err, "unable to start manager")
		os.Exit(1)
	}
}
