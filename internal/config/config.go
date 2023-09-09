package config

type Config struct {
	Editors     []Editor  `yaml:"editors"`
	TypeFilters []*Filter `yaml:"typefilters"`
}

func (c *Config) Validate() error {
	for _, f := range c.TypeFilters {
		if err := f.validate(); err != nil {
			return err
		}
	}

	return nil
}

func (c *Config) Filter(name string) FilterResult {
	for _, f := range c.TypeFilters {
		if result := f.Applies(name); result != FilterResultNone {
			return result
		}
	}

	return FilterResultInclude
}
