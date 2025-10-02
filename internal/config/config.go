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

	// FileMode controls how documentation files are generated.
	// Defaults to "single-file".
	FileMode string `yaml:"fileMode"`

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

const (
	FileModeSingleFile   = "Single-File"
	FileModeMultipleFile = "Multiple-File"
)

// Standard returns the standard, as a basis for loading other configuration,
// or for export.
func Standard() *Config {
	return &Config{
		FileMode:    FileModeSingleFile,
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

	// Ensure our nested config exists
	if c.ClassDiagrams == nil {
		c.ClassDiagrams = &ClassDiagram{}
	}

	c.ClassDiagrams.Enabled = value
}

// SetFileMode sets the FileMode field to the provided value.
// If the value is nil, the field is not changed.
// The mode value is normalized to handle case variations.
func (c *Config) SetFileMode(fileMode *string) {
	if fileMode == nil || *fileMode == "" {
		// No value passed, do nothing
		return
	}

	c.FileMode = *fileMode
}

// HasFileMode checks if the config has the specified file mode.
// fileMode is the file mode to check for.
func (c *Config) HasFileMode(fileMode string) bool {
	return strings.EqualFold(c.FileMode, fileMode)
}

// Validate checks if the config is valid.
func (c *Config) Validate() error {
	if !c.HasFileMode(FileModeSingleFile) &&
		!c.HasFileMode(FileModeMultipleFile) {
		return errors.Errorf(
			"invalid file mode %q: must be either %q or %q",
			c.FileMode,
			FileModeSingleFile,
			FileModeMultipleFile)
	}

	for _, f := range c.TypeFilters {
		if err := f.validate(); err != nil {
			return err
		}
	}

	return nil
}
