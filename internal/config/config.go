package config

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type Config struct {
	// Editors allow you to make precision changes to the documentation output. Editors are applied in the order specified.
	Editors []Editor `yaml:"editors"`

	// TypeFilters allow you to filter out types from the output.
	// Filters are applied in the order specified, with earlier filters taking priority over later ones.
	TypeFilters []*Filter `yaml:"typeFilters"`

	// PrettyPrint controls whether the Markdown output is pretty-printed or not. Defaults to true.
	PrettyPrint bool `yaml:"prettyPrint"`

	// TemplatePath is the path to folder containing templates to use for rendering the documentation.
	TemplatePath string `yaml:"templatePath"`
}

// Default returns the default configuration, as a basis for loading
// or for export.
func Default() *Config {
	return &Config{
		PrettyPrint: true,
	}
}

// Load populates this config from the given path
func (c *Config) Load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return errors.Wrapf(err, "opening config file %q", path)
	}

	defer file.Close()

	decoder := yaml.NewDecoder(file)
	decoder.KnownFields(true)

	err = decoder.Decode(c)
	if err != nil {
		return errors.Wrapf(err, "parsing config file %q", path)
	}

	return nil
}

func (c *Config) OverrideTemplatePath(path *string) {
	if path != nil && *path != "" {
		c.TemplatePath = *path
	}
}

func (c *Config) Validate() error {
	for _, f := range c.TypeFilters {
		if err := f.validate(); err != nil {
			return err
		}
	}

	return nil
}
