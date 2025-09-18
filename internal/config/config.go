package config

import (
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type Config struct {
	// Editors allow you to make precision changes to the documentation output.
	// Editors are applied in the order specified.
	Editors []Editor `yaml:"editors"`

	// ExternalLinks allow you to add links to external documentation.
	ExternalLinks []*ExternalLink `yaml:"externalLinks"`

	// Mode controls how documentation files are generated.
	// Defaults to "single-file".
	Mode string `yaml:"mode"`

	// PrettyPrint controls whether the Markdown output is pretty-printed or not.
	// Defaults to true.
	PrettyPrint bool `yaml:"prettyPrint"`

	// TemplatePath is the path to a folder containing templates to use for
	// rendering the documentation.
	TemplatePath string `yaml:"templatePath"`

	// TypeFilters allow you to filter out types from the output.
	// Filters are applied in the order specified,
	// with earlier filters taking priority over later ones.
	TypeFilters []*Filter `yaml:"typeFilters"`

	// ClassDiagrams allow you to add class diagrams to the documentation.
	ClassDiagrams *ClassDiagram `yaml:"classDiagrams"`
}

// Standard returns the standard, as a basis for loading other configuration,
// or for export.
func Standard() *Config {
	return &Config{
		Mode:        "single-file",
		PrettyPrint: true,
	}
}

// Load populates this config from the given path.
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

// writeTo writes the current config as YAML to the provided writer.
func (c *Config) writeTo(writer io.Writer) error {
	encoder := yaml.NewEncoder(writer)
	encoder.SetIndent(2)

	err := encoder.Encode(c)
	if err != nil {
		return errors.Wrap(err, "writing config")
	}

	return nil
}

// Save writes the current config to the provided path.
func (c *Config) Save(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.Wrapf(err, "creating config file %q", path)
	}

	defer file.Close()

	return c.writeTo(file)
}

func (c *Config) SetTemplatePath(path *string) {
	if path == nil || *path == "" {
		// No value passed, do nothing
		return
	}

	c.TemplatePath = *path
}

// EnableClassDiagrams sets the ClassDiagrams field to the provided value.
// If the value is nil, the field is not changed.
func (c *Config) EnableClassDiagrams(value *bool) {
	if value == nil {
		// No value passed, do nothing
		return
	}

	// Ensure we nested config exists
	if c.ClassDiagrams == nil {
		c.ClassDiagrams = &ClassDiagram{}
	}

	c.ClassDiagrams.Enabled = value
}

// SetMode sets the Mode field to the provided value.
// If the value is nil, the field is not changed.
// The mode value is normalized to handle case variations.
func (c *Config) SetMode(mode *string) {
	if mode == nil || *mode == "" {
		// No value passed, do nothing
		return
	}

	// Apply Postel's Law: normalize case variations
	normalizedMode := strings.ToLower(strings.TrimSpace(*mode))
	
	switch normalizedMode {
	case "single-file":
		c.Mode = "single-file"
	case "multiple-file":
		c.Mode = "multiple-file"
	default:
		// Set the original value and let validation catch invalid modes
		c.Mode = *mode
	}
}

// validateMode validates and normalizes the mode field.
func (c *Config) validateMode() error {
	// Handle empty mode by setting default
	if c.Mode == "" {
		c.Mode = "single-file"
		return nil
	}

	// Apply Postel's Law: normalize case variations  
	normalizedMode := strings.ToLower(strings.TrimSpace(c.Mode))
	
	switch normalizedMode {
	case "single-file":
		c.Mode = "single-file"
	case "multiple-file":
		c.Mode = "multiple-file"
	default:
		return errors.Errorf("invalid mode %q: must be either 'single-file' or 'multiple-file'", c.Mode)
	}

	return nil
}

func (c *Config) Validate() error {
	// Validate and normalize the mode
	if err := c.validateMode(); err != nil {
		return err
	}

	for _, f := range c.TypeFilters {
		if err := f.validate(); err != nil {
			return err
		}
	}

	return nil
}
