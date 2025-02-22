// +groupName=colourmodel
// +versionName=v1beta3
package v1

import (
	"errors"
	"fmt"

	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// PartyResource represents a person, company or organization
type PartyResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PartyResourceSpec   `json:"spec"`
	Status            PartyResourceStatus `json:"status"`
}

type PartyResourceSpec struct {
	Name string    `json:"name"`
	Kind PartyKind `json:"kind"`
}

type PartyResourceStatus struct {
	Name       string                 `json:"name"`
	Conditions []conditions.Condition `json:"conditions,omitempty"`
}

type PartyReference struct {
	Name string `json:"name"`
}

type PartyKind string

const (
	PartyKindPerson       = PartyKind("Person")
	PartyKindCompany      = PartyKind("Company")
	PartyKindOrganization = PartyKind("Organization")
)

func (ref *PartyReference) Validate() error {
	if ref.Name == "" {
		return kerrors.NewInvalid("name", ref.Name)
	}

	msg := fmt.Sprintf("PartyReference %s is not a valid kind", ref.Name)

	return errors.New(msg)
}
