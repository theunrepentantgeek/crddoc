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
		"- Generated from: /path/to/source",
		"  continued on another line",
		"* ARM URI: /subscriptions/{subscriptionId}",
		"Following content.",
	})

	g.Expect(actual).To(Equal(
		"Generator information:\n\n" +
			"- Generated from: /path/to/source\n" +
			"  continued on another line\n" +
			"* ARM URI: /subscriptions/{subscriptionId}\n\n" +
			"Following content."))
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

func TestUnwrapTable_GivenNoList_PreservesExistingFormatting(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	f := &Functions{}

	actual := f.unwrapTable([]string{"Content split", "across lines."})

	g.Expect(actual).To(Equal("Content split across lines."))
}
