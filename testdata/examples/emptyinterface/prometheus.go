// Code generated by WrapGen. DO NOT EDIT.
package emptyinterface

import (

	"github.com/prometheus/client_golang/prometheus"
)

type EmptyInterfaceWithPrometheus struct {
	base EmptyInterface
	metric prometheus.ObserverVec
	instanceName string
}

// NewEmptyInterfaceWithPrometheus returns an instance of the EmptyInterface decorated with prometheus summary metric.
func NewEmptyInterfaceWithPrometheus(base EmptyInterface, metric prometheus.ObserverVec, instanceName string) *EmptyInterfaceWithPrometheus {
	return &EmptyInterfaceWithPrometheus{
		base: base,
		metric: metric,
		instanceName: instanceName,
	}
}


