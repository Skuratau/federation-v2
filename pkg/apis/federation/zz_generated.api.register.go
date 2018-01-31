/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by apiregister-gen. Do not edit it manually!

package federation

import (
	"fmt"
	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
)

var (
	InternalFederatedSecret = builders.NewInternalResource(
		"federatedsecrets",
		"FederatedSecret",
		func() runtime.Object { return &FederatedSecret{} },
		func() runtime.Object { return &FederatedSecretList{} },
	)
	InternalFederatedSecretStatus = builders.NewInternalResourceStatus(
		"federatedsecrets",
		"FederatedSecretStatus",
		func() runtime.Object { return &FederatedSecret{} },
		func() runtime.Object { return &FederatedSecretList{} },
	)
	// Registered resources and subresources
	ApiVersion = builders.NewApiGroup("federation.k8s.io").WithKinds(
		InternalFederatedSecret,
		InternalFederatedSecretStatus,
	)

	// Required by code generated by go2idl
	AddToScheme        = ApiVersion.SchemaBuilder.AddToScheme
	SchemeBuilder      = ApiVersion.SchemaBuilder
	localSchemeBuilder = &SchemeBuilder
	SchemeGroupVersion = ApiVersion.GroupVersion
)

// Required by code generated by go2idl
// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Required by code generated by go2idl
// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// +genclient
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedSecret struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   FederatedSecretSpec
	Status FederatedSecretStatus
}

type FederatedSecretSpec struct {
	Template corev1.Secret
}

type FederatedSecretStatus struct {
}

//
// FederatedSecret Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedSecretStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedSecretStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedSecretList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []FederatedSecret
}

func (FederatedSecret) NewStatus() interface{} {
	return FederatedSecretStatus{}
}

func (pc *FederatedSecret) GetStatus() interface{} {
	return pc.Status
}

func (pc *FederatedSecret) SetStatus(s interface{}) {
	pc.Status = s.(FederatedSecretStatus)
}

func (pc *FederatedSecret) GetSpec() interface{} {
	return pc.Spec
}

func (pc *FederatedSecret) SetSpec(s interface{}) {
	pc.Spec = s.(FederatedSecretSpec)
}

func (pc *FederatedSecret) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *FederatedSecret) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc FederatedSecret) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store FederatedSecret.
// +k8s:deepcopy-gen=false
type FederatedSecretRegistry interface {
	ListFederatedSecrets(ctx request.Context, options *internalversion.ListOptions) (*FederatedSecretList, error)
	GetFederatedSecret(ctx request.Context, id string, options *metav1.GetOptions) (*FederatedSecret, error)
	CreateFederatedSecret(ctx request.Context, id *FederatedSecret) (*FederatedSecret, error)
	UpdateFederatedSecret(ctx request.Context, id *FederatedSecret) (*FederatedSecret, error)
	DeleteFederatedSecret(ctx request.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewFederatedSecretRegistry(sp builders.StandardStorageProvider) FederatedSecretRegistry {
	return &storageFederatedSecret{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageFederatedSecret struct {
	builders.StandardStorageProvider
}

func (s *storageFederatedSecret) ListFederatedSecrets(ctx request.Context, options *internalversion.ListOptions) (*FederatedSecretList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*FederatedSecretList), err
}

func (s *storageFederatedSecret) GetFederatedSecret(ctx request.Context, id string, options *metav1.GetOptions) (*FederatedSecret, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*FederatedSecret), nil
}

func (s *storageFederatedSecret) CreateFederatedSecret(ctx request.Context, object *FederatedSecret) (*FederatedSecret, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, true)
	if err != nil {
		return nil, err
	}
	return obj.(*FederatedSecret), nil
}

func (s *storageFederatedSecret) UpdateFederatedSecret(ctx request.Context, object *FederatedSecret) (*FederatedSecret, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil)
	if err != nil {
		return nil, err
	}
	return obj.(*FederatedSecret), nil
}

func (s *storageFederatedSecret) DeleteFederatedSecret(ctx request.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil)
	return sync, err
}