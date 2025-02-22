package config

type ExternalLink struct {
	// ImportPath is the import path for the package that the link points to.
	ImportPath string `yaml:"importPath"`
	// URLTemplate is a template for the URL to the package documentation.
	URLTemplate string `yaml:"urlTemplate"`
}
