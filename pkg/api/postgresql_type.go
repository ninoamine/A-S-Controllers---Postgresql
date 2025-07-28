package api

import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PostgreSQLDatabaseSpec struct {
    DBName string `json:"dbName"`
}

type PostgreSQLDatabase struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`

    Spec PostgreSQLDatabaseSpec `json:"spec"`
}

type PostgreSQLDatabaseList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    Items           []PostgreSQLDatabase `json:"items"`
}
