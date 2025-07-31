package api

import (
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)


var (
	Scheme = runtime.NewScheme()
)

func init() {
	utilruntime.Must(addKnownTypes(Scheme))
}