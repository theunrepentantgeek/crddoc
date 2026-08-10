package cmd

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/theunrepentantgeek/crddoc/internal/config"
)

func TestDocumentCRDsOptions_ApplyToConfig_SkipFormatter(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	skipFormatter := true
	options := documentCRDsOptions{skipFormatter: &skipFormatter}
	cfg := config.Standard()

	options.applyToConfig(cfg)

	g.Expect(cfg.PrettyPrint).To(BeFalse())
}
