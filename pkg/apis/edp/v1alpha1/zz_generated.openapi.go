// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"nexus-operator/pkg/apis/edp/v1alpha1.Nexus":       schema_pkg_apis_edp_v1alpha1_Nexus(ref),
		"nexus-operator/pkg/apis/edp/v1alpha1.NexusSpec":   schema_pkg_apis_edp_v1alpha1_NexusSpec(ref),
		"nexus-operator/pkg/apis/edp/v1alpha1.NexusStatus": schema_pkg_apis_edp_v1alpha1_NexusStatus(ref),
	}
}

func schema_pkg_apis_edp_v1alpha1_Nexus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Nexus is the Schema for the nexus API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("nexus-operator/pkg/apis/edp/v1alpha1.NexusSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("nexus-operator/pkg/apis/edp/v1alpha1.NexusStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta", "nexus-operator/pkg/apis/edp/v1alpha1.NexusSpec", "nexus-operator/pkg/apis/edp/v1alpha1.NexusStatus"},
	}
}

func schema_pkg_apis_edp_v1alpha1_NexusSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NexusSpec defines the desired state of Nexus",
				Properties: map[string]spec.Schema{
					"keycloakSpec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("nexus-operator/pkg/apis/edp/v1alpha1.KeycloakSpec"),
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"volumes": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("nexus-operator/pkg/apis/edp/v1alpha1.NexusVolumes"),
									},
								},
							},
						},
					},
				},
				Required: []string{"version"},
			},
		},
		Dependencies: []string{
			"nexus-operator/pkg/apis/edp/v1alpha1.KeycloakSpec", "nexus-operator/pkg/apis/edp/v1alpha1.NexusVolumes"},
	}
}

func schema_pkg_apis_edp_v1alpha1_NexusStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NexusStatus defines the observed state of Nexus",
				Properties: map[string]spec.Schema{
					"available": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"lastTimeUpdated": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "date-time",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}
