package client

import (

    
    "k8s.io/client-go/rest"
    "k8s.io/client-go/kubernetes/scheme"
    
    "github.com/ninoamine/A-S-Controllers---Postgresql/pkg/api"
)

type PostgreSQLDatabaseClient struct {
    restClient rest.Interface
    ns         string
}

// NewPostgreSQLDatabaseClient crée un nouveau client
func NewPostgreSQLDatabaseClient(cfg *rest.Config, namespace string) (*PostgreSQLDatabaseClient, error) {
    // Configure le client REST pour notre API
    cfg.ContentConfig.GroupVersion = &api.SchemeGroupVersion
    cfg.APIPath = "/apis"
    cfg.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
    cfg.UserAgent = rest.DefaultKubernetesUserAgent()
    
    // Ajoute notre scheme au scheme global
    if err := api.AddToScheme(scheme.Scheme); err != nil {
        return nil, err
    }
    
    client, err := rest.RESTClientFor(cfg)
    if err != nil {
        return nil, err
    }
    
    return &PostgreSQLDatabaseClient{
        restClient: client,
        ns:         namespace,
    }, nil
}

