package controller

import (
    "time"
    
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/apimachinery/pkg/runtime"
    "k8s.io/apimachinery/pkg/watch"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/rest"
    "k8s.io/client-go/tools/cache"
    "k8s.io/client-go/util/workqueue"
    
    "github.com/ninoamine/A-S-Controllers---Postgresql/pkg/api"
)


type PostgreSQLController struct {
    kubeClient    kubernetes.Interface
    restClient    rest.Interface
    informer      cache.SharedIndexInformer
    queue         workqueue.TypedRateLimitingInterface[string]   
    postgresURL   string
	watchNamespaces []string
}

func NewPostgreSQLController(kubeClient kubernetes.Interface, restClient rest.Interface, postgresUrl string, watchNamespaces []string) *PostgreSQLController {
	queue := workqueue.NewTypedRateLimitingQueue(workqueue.DefaultTypedControllerRateLimiter[string]())

	controller := &PostgreSQLController{
		kubeClient:    kubeClient,
		restClient:    restClient,
		queue:         queue,
		postgresURL:   postgresUrl,
		watchNamespaces: watchNamespaces,
	}
    controller.informer = cache.NewSharedIndexInformer(
        &cache.ListWatch{
            ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
                return controller.listPostgreSQLDatabases(options)
            },
            WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
                return controller.watchPostgreSQLDatabases(options)
            },
        },
        &api.PostgreSQLDatabase{},
        30*time.Second, // Resync period
        cache.Indexers{},
    )
	return controller
}