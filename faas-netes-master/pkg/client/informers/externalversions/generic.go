/*
Copyright 2019-2020 OpenFaaS Authors

Licensed under the MIT license. See LICENSE file in the project root for full license information.
*/

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1 "github.com/openfaas/faas-netes/pkg/apis/openfaas/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=openfaas.com, Version=v1
	case v1.SchemeGroupVersion.WithResource("functions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Openfaas().V1().Functions().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("profiles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Openfaas().V1().Profiles().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}