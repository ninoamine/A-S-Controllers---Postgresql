package api

import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type PostgreSQLDatabaseSpec struct {
    DBName string `json:"dbName"`
}

type PostgreSQLDatabaseStatus struct {
    Phase string `json:"phase,omitempty"`
    Message string `json:"message,omitempty"`
    Ready bool `json:"ready"`
    LastUpdated metav1.Time `json:"lastUpdated,omitempty"`
}

type PostgreSQLDatabase struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`
    Spec PostgreSQLDatabaseSpec `json:"spec"`
	Status PostgreSQLDatabaseStatus `json:"status,omitempty"`
}

type PostgreSQLDatabaseList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    Items           []PostgreSQLDatabase `json:"items"`
}



func (in *PostgreSQLDatabaseSpec) DeepCopy() *PostgreSQLDatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLDatabaseSpec)
	*out = *in
	return out
}


func (in *PostgreSQLDatabaseStatus) DeepCopy() *PostgreSQLDatabaseStatus {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLDatabaseStatus)
	*out = *in
	out.LastUpdated = *in.LastUpdated.DeepCopy()
	return out
}


func (in *PostgreSQLDatabase) DeepCopy() *PostgreSQLDatabase {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLDatabase)
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = *in.Spec.DeepCopy()
	out.Status = *in.Status.DeepCopy()
	return out
}

func (in *PostgreSQLDatabase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *PostgreSQLDatabaseList) DeepCopy() *PostgreSQLDatabaseList {
    if in == nil {
        return nil
    }
    out := new(PostgreSQLDatabaseList)
    *out = *in
    out.TypeMeta = in.TypeMeta
    in.ListMeta.DeepCopyInto(&out.ListMeta)
    if in.Items != nil {
        in, out := &in.Items, &out.Items
        *out = make([]PostgreSQLDatabase, len(*in))
        for i := range *in {
            (*out)[i] = *(*in)[i].DeepCopy()
        }
    }
    return out
}

func (in *PostgreSQLDatabaseList) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}