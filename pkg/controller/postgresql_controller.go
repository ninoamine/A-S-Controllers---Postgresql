package controller

import (
    "time"
	"context"
	"fmt"
    
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
	controller.informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
        AddFunc:    controller.onAdd,
        UpdateFunc: controller.onUpdate,
        DeleteFunc: controller.onDelete,
    })
	return controller
}
// listPostgreSQLDatabases liste les ressources selon la configuration namespace
func (c *PostgreSQLController) listPostgreSQLDatabases(options metav1.ListOptions) (runtime.Object, error) {
    if len(c.watchNamespaces) == 0 {
        // Tous les namespaces - utilise cluster-scoped list
        result := &api.PostgreSQLDatabaseList{}
        err := c.restClient.Get().
            Resource("postgresqldatabases").
            VersionedParams(&options, metav1.ParameterCodec).
            Do(context.TODO()).
            Into(result)
        return result, err
    }
    
    // Namespaces spécifiques - combine les résultats
    allDatabases := &api.PostgreSQLDatabaseList{}
    for _, ns := range c.watchNamespaces {
        result := &api.PostgreSQLDatabaseList{}
        err := c.restClient.Get().
            Namespace(ns).
            Resource("postgresqldatabases").
            VersionedParams(&options, metav1.ParameterCodec).
            Do(context.TODO()).
            Into(result)
        if err != nil {
            return nil, err
        }
        allDatabases.Items = append(allDatabases.Items, result.Items...)
    }
    return allDatabases, nil
}

// watchPostgreSQLDatabases watch les changements sur les ressources
func (c *PostgreSQLController) watchPostgreSQLDatabases(options metav1.ListOptions) (watch.Interface, error) {
    if len(c.watchNamespaces) == 0 {
        // Tous les namespaces
        return c.restClient.Get().
            Resource("postgresqldatabases").
            VersionedParams(&options, metav1.ParameterCodec).
            Watch(context.TODO())
    }
    
    // Pour plusieurs namespaces, on va watcher le premier pour l'instant
    // (Une vraie implémentation multi-namespace est plus complexe)
    return c.restClient.Get().
        Namespace(c.watchNamespaces[0]).
        Resource("postgresqldatabases").
        VersionedParams(&options, metav1.ParameterCodec).
        Watch(context.TODO())
}


// onAdd est appelé quand une nouvelle ressource PostgreSQLDatabase est créée
func (c *PostgreSQLController) onAdd(obj interface{}) {
    key, err := cache.MetaNamespaceKeyFunc(obj)
    if err != nil {
        fmt.Printf("Erreur lors de la création de la clé: %v\n", err)
        return
    }
    c.queue.Add(key)
}

// onUpdate est appelé quand une ressource PostgreSQLDatabase est modifiée
func (c *PostgreSQLController) onUpdate(oldObj, newObj interface{}) {
    key, err := cache.MetaNamespaceKeyFunc(newObj)
    if err != nil {
        fmt.Printf("Erreur lors de la création de la clé: %v\n", err)
        return
    }
    c.queue.Add(key)
}

// onDelete est appelé quand une ressource PostgreSQLDatabase est supprimée
func (c *PostgreSQLController) onDelete(obj interface{}) {
    key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
    if err != nil {
        fmt.Printf("Erreur lors de la création de la clé: %v\n", err)
        return
    }
    c.queue.Add(key)
}