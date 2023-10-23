// +groupName=colourmodel
// +versionName=v1beta3
package v1

// PartyResource represents a person, company or organization
type PartyResource struct {
	Spec   PartyResourceSpec   `json:"spec"`
	Status PartyResourceStatus `json:"status"`
}

type PartyResourceSpec struct {
	Name string    `json:"name"`
	Kind PartyKind `json:"kind"`
}

type PartyResourceStatus struct {
	Name string `json:"name"`
}

type PartyReference struct {
	Name string `json:"name"`
}

type PartyKind string

const (
	PartyKindPerson       = "Person"
	PartyKindCompany      = "Company"
	PartyKindOrganization = "Organization"
)
