package config

type Config struct {
	Filters []*Filter `yaml:"filters"`
}

func (c *Config) Validate() error {
	for _, f := range c.Filters {
		if err := f.validate(); err != nil {
			return err
		}
	}

	return nil
}

func (c *Config) Filter(name string) FilterResult {
	for _, f := range c.Filters {
		if result := f.Applies(name); result != FilterResultNone {
			return result
		}
	}

	return FilterResultInclude
}
