package config

import (
	"io"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Editors     []Editor  `yaml:"editors"`
	TypeFilters []*Filter `yaml:"typeFilters"`
	PrettyPrint bool      `yaml:"prettyPrint"`
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

	data, err := io.ReadAll(file)
	if err != nil {
		return errors.Wrapf(err, "reading config file %q", path)
	}

	err = yaml.Unmarshal(data, c)
	if err != nil {
		return errors.Wrapf(err, "parsing config file %q", path)
	}

	return nil
}

func (c *Config) Validate() error {
	for _, f := range c.TypeFilters {
		if err := f.validate(); err != nil {
			return err
		}
	}

	return nil
}
