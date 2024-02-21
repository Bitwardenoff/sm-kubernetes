/*
Copyright 2024 Bitwarden, Inc..

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BitwardenSecretSpec defines the desired state of BitwardenSecret
type BitwardenSecretSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// The base URL for Secrets Manager API.  When in doubt, leave blank
	// +kubebuilder:Required
	BitwardenApiUrl string `json:"bitwardenApiUrl,omitempty"`
	// The base URL for the Bitwarden Identity API.  When in doubt, leave blank
	// +kubebuilder:Required
	IdentityApiUrl string `json:"identityApiUrl,omitempty"`
	// The organization ID for your organization
	// +kubebuilder:Optional
	OrganizationId string `json:"organizationId"`
	// The projects to sync from the organization
	// +kubebuilder:Required
	Projects []Project `json:"projects"`
	// The secret key reference for the authorization token used to connect to Secrets Manager
	// +kubebuilder:Required
	AuthToken AuthToken `json:"authToken"`
}

type Project struct {
	// The Project ID for your organization
	// +kubebuilder:Required
	ProjectId string `json:"projectId"`
	// The name of the secret for the
	// +kubebuilder:Required
	SecretName string `json:"secretName"`
	// Optional filter to specify that only these IDs are included in the secrets created
	// +kubebuilder:Optional
	IncludedSecretIds []string `json:"includedSecretIds,omitempty"`
	// Optional filter to specify that these secrets should not be included in the secrets created
	// +kubebuilder:Optional
	ExcludedSecretIds []string `json:"excludedSecretIds,omitempty"`
	// The mapping of project secrets to K8s secret keys
	// +kubebuilder:Required
	SecretMap []SecretMap `json:"map"`
}

type AuthToken struct {
	// The name of the secret where the authorization token is stored
	// +kubebuilder:Required
	SecretName string `json:"secretName"`
	// The key of the secret where the authorization token is stored
	// +kubebuilder:Required
	SecretKey string `json:"secretKey"`
}

type SecretMap struct {
	// The ID of the secret in Secrets Manager
	// +kubebuilder:Required
	BwSecretId string `json:"bwSecretId"`
	// The name of the mapped key in the created Kubernetes secret
	// +kubebuilder:Required
	SecretKeyName string `json:"secretKeyName"`
}

// BitwardenSecretStatus defines the observed state of BitwardenSecret
type BitwardenSecretStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// BitwardenSecret is the Schema for the bitwardensecrets API
type BitwardenSecret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BitwardenSecretSpec   `json:"spec,omitempty"`
	Status BitwardenSecretStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// BitwardenSecretList contains a list of BitwardenSecret
type BitwardenSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BitwardenSecret `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BitwardenSecret{}, &BitwardenSecretList{})
}
