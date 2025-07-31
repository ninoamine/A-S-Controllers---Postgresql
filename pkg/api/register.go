package api

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime")


const (
	GroupName = "ninoamine.com"
	Version = "v1alpha1"
)

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: Version}

func Resource(resource string) schema.GroupResource{
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}


func addKnownTypes(scheme *runtime.Scheme) error {
    scheme.AddKnownTypes(SchemeGroupVersion,
        &PostgreSQLDatabase{},
        &PostgreSQLDatabaseList{},
    )
    
    metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
    return nil
}