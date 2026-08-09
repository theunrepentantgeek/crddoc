package functions

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/theunrepentantgeek/crddoc/internal/model"
)

func TestAsObjectReferenceList_GivenMultiplePropertiesOnTheSameObject_ReturnsEachObjectOnce(
	t *testing.T,
) {
	t.Parallel()
	g := NewGomegaWithT(t)

	first := model.NewPropertyReference("First", "first", "One")
	second := model.NewPropertyReference("Second", "second", "One")
	duplicate := model.NewPropertyReference("First", "first", "Two")

	result := asObjectReferenceList([]model.PropertyReference{first, second, duplicate})

	g.Expect(result).To(Equal([]ListItem[model.PropertyReference]{
		{Item: first, First: true},
		{Item: second, Last: true},
	}))
}
