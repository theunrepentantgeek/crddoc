package v1

import "time"

// +groupName=crm.example.com
// +version=v1alpha
// +kubebuilder:validation:Optional

type PersonResource struct {
	Spec   PersonResourceSpec   `json:"spec"`
	Status PersonResourceStatus `json:"status"`
}

type PersonResourceSpec struct {
	// +kubebuilder:validation:Required
	FullName string `json:"fullName"`

	// +kubebuilder:validation:Optional
	KnownAs string

	// +kubebuilder:validation:Required
	FamilyName string `json:"familyName,inline"`

	// +kubebuilder:validation:Required
	FamilyKey string `json:"familyKey,omitempty"`

	Aliases []string `json:"aliases,omitempty"`

	Children []PersonReference `json:"children,omitempty"`

	Friends map[string]PersonReference `json:"friends,omitempty"`

	BirthDate *time.Time `json:"birthDate,omitempty"`
}

type PersonResourceStatus struct {
	Age *int `json:"age,omitempty"`
}

type PersonReference struct {
	Name string `json:"name"`
}
