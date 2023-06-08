package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ClusterSpec struct {
	// AccessObjectRefs represents references to objects providing access info to the cluster.
	// It could be a kubeconf stored in a secret
	// +optional
	AccessObjectRefs []AccessObjectRef `json:"accessObjectRef,omitempty"`

	// HealthProbe is used to coordinate the heartbeat time of to check the healthiness of the cluster.
	// +optional
	// Some implementations may have build-in health probe configurations.
	HealthProbe HealthProbe `json:"healthProbe,omitempty"`

	// Taints is a property of cluster that allow the cluster to be repelled when scheduling.
	// +optional
	Taints []corev1.Taint `json:"taints,omitempty"`
}

type HealthProbe struct {
	// HeartbeatIntervalSeconds is the interval of the cluster's heartbeat to check the
	// availability of the cluster.
	HeartbeatIntervalSeconds int32 `json:"heatbeatIntervalSeconds"`
}

type AccessObjectRef struct {
	// Type is type of the access info. If the type is KUBECONFIG, the realted object
	// should be a secret containing kubeconfig key.
	Type string `json:"type"`

	// Group is the API Group of the Kubernetes resource,
	// empty string indicates it is in core group.
	// +optional
	Group string `json:"group,omitempty"`

	// Resource is the resource name of the Kubernetes resource.
	// +kubebuilder:validation:Required
	// +required
	Resource string `json:"resource"`

	// Name is the name of the Kubernetes resource.
	// +kubebuilder:validation:Required
	// +required
	Name string `json:"name"`

	// Name is the namespace of the Kubernetes resource, empty string indicates
	// it is a cluster scoped resource.
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

type ClusterStatus struct {
	// Conditions contains the different condition statuses for this cluster.
	Conditions []metav1.Condition `json:"conditions"`

	// Version represents the kubernetes version of the cluster.
	Version ClusterVersion `json:"version,omitempty"`

	// Resource represents the resource of the cluster.
	Resources Resources `json:"resources,omitempty"`

	// Properties represents properties of collected from the cluster,
	// for example a unique cluster identifier (id.k8s.io).
	// The set of properties is not uniform across a fleet, some properties can be
	// vendor or version specific and may not be included from all clusters.
	// +optional
	Properties []Property `json:"properties,omitempty"`
}

// ClusterVersion represents version information about the cluster.
type ClusterVersion struct {
	// Kubernetes is the kubernetes version of managed cluster.
	// +optional
	Kubernetes string `json:"kubernetes,omitempty"`
}

type Resources struct {
	// Capacity represents the total resource capacity from all nodeStatuses
	// on the cluster.
	// +optional
	Capacity corev1.ResourceList `json:"capacity,omitempty"`

	// Allocatable represents the total allocatable resources on the cluster.
	// +optional
	Allocatable corev1.ResourceList `json:"allocatable,omitempty"`
}

// Property represents a Property collected from a cluster.
type Property struct {
	// Name is the name of a propertie resource on cluster. It's a well known
	// or customized name to identify the propertie.
	// +kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`

	// Value is a property-dependent string
	// +kubebuilder:validation:MaxLength=1024
	// +kubebuilder:validation:MinLength=1
	Value string `json:"value"`
}

const (
	// ClusterConditionJoined means the cluster has successfully joined the control.
	ClusterConditionJoined string = "Joined"
	// Healthey means the cluster is healthy.
	ClusterConditionHealthy string = "Healthy"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Cluster is the Schema for the cluster inventory API
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// spec defines the spec of a cluster.
	// +optional
	Spec ClusterSpec `json:"spec,omitempty"`
	// status defines the status of a cluster.
	Status ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata.
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	// List of clusters.
	// +listType=set
	Items []Cluster `json:"items"`
}
