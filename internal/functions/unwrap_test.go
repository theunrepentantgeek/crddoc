package functions

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestUnwrap_GivenList_FormatsMarkdown(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	f := &Functions{}

	actual := f.unwrap([]string{
		"Generator information:",
		"",
		"- Generated from: /path/to/source",
		"  continued on another line",
		"* ARM URI: /subscriptions/{subscriptionId}",
		"",
		"Following content.",
	})

	g.Expect(actual).To(Equal(
		"Generator information:\n\n" +
			"- Generated from: /path/to/source\n" +
			"  continued on another line\n" +
			"* ARM URI: /subscriptions/{subscriptionId}\n\n" +
			"Following content."))
}

func TestUnwrap_GivenUnindentedListContinuations_JoinsItems(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	f := &Functions{}

	actual := f.unwrap([]string{
		"Each condition can be of one of the following types:",
		"* __Leaf Condition -__ must contain 'field' and either 'equals' or 'containsAny'.",
		"which may match a field.",
		"* __AnyOf Condition -__ must contain __only__",
		"'anyOf' (which is an array of Leaf Conditions).",
		"set in an AnyOf Condition.",
	})

	g.Expect(actual).To(Equal(
		"Each condition can be of one of the following types:\n\n" +
			"* __Leaf Condition -__ must contain 'field' and either 'equals' or 'containsAny'. " +
			"which may match a field.\n" +
			"* __AnyOf Condition -__ must contain __only__ " +
			"'anyOf' (which is an array of Leaf Conditions). " +
			"set in an AnyOf Condition."))
}

func TestUnwrap_GivenPostListProse_SeparatesFormattedAndUnformattedText(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	f := &Functions{}

	actual := f.unwrap([]string{
		"- first item",
		"Following content.",
		"- second item",
		"_Please note this is separate._",
		"- third item",
		"**Important:** separate paragraph.",
	})

	g.Expect(actual).To(Equal(
		"- first item\n\nFollowing content.\n\n" +
			"- second item\n\n_Please note this is separate._\n\n" +
			"- third item\n\n**Important:** separate paragraph."))
}

func TestUnwrap_GivenNoList_PreservesExistingFormatting(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	f := &Functions{}

	actual := f.unwrap([]string{
		"Content split",
		"across lines.",
		"",
		"Another paragraph.",
	})

	g.Expect(actual).To(Equal("Content split across lines. <br/>Another paragraph."))
}

func TestUnwrapTable_GivenList_FormatsHTML(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	f := &Functions{}

	actual := f.unwrapTable([]string{
		"Generator information:",
		"- Generated from: /path/to/source",
		"  continued on another line",
		"* ARM URI: /subscriptions/{subscriptionId}",
		"Following content.",
	})

	g.Expect(actual).To(Equal(
		"Generator information:" +
			"<ul>" +
			"<li>Generated from: /path/to/source continued on another line</li>" +
			"<li>ARM URI: /subscriptions/{subscriptionId}</li>" +
			"</ul>" +
			"Following content."))
}

func TestUnwrapTable_GivenUnindentedListContinuations_JoinsItems(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	f := &Functions{}

	actual := f.unwrapTable([]string{
		"- first item",
		"with more detail",
		"* second item",
		"'with quoted detail'",
		"Following content.",
	})

	g.Expect(actual).To(Equal(
		"<ul><li>first item with more detail</li>" +
			"<li>second item 'with quoted detail'</li></ul>" +
			"Following content."))
}

func TestUnwrapTable_GivenFormattedPostListProse_SeparatesText(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	f := &Functions{}

	actual := f.unwrapTable([]string{
		"- first item",
		"_Please note this is separate._",
		"- second item",
		"**Important:** separate paragraph.",
	})

	g.Expect(actual).To(Equal(
		"<ul><li>first item</li></ul>_Please note this is separate._" +
			"<ul><li>second item</li></ul>**Important:** separate paragraph."))
}

func TestUnwrapTable_GivenNoList_PreservesExistingFormatting(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	f := &Functions{}

	actual := f.unwrapTable([]string{"Content split", "across lines."})

	g.Expect(actual).To(Equal("Content split across lines."))
}
