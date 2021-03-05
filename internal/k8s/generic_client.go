// Copyright 2020-2021 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package k8s

import (
	"fmt"
	
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	 "sigs.k8s.io/controller-runtime/pkg/cluster"
)

var (
	scheme   = runtime.NewScheme()
)

func init() {
	_ = clientgoscheme.AddToScheme(scheme)
}

type GenericClusterClient struct {
	cluster.Cluster
}

func NewGenericClusterClient(contextName, kubeconfig string) (*GenericClusterClient, error) {
	config, _, err := NewConfig(contextName, kubeconfig)
    if err != nil {
		return nil, err
	}

	setOpts := func(opts *cluster.Options) {
		opts.Scheme = scheme
	}

	clusterClinet, err := cluster.New(config, setOpts)
	if err != nil {
		return nil, fmt.Errorf("unable create generic client for contex %q: %w", contextName, err)
	}

	return &GenericClusterClient{clusterClinet}, nil
}
