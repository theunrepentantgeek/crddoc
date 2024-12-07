# crddoc

[![Validate Pull Request](https://github.com/theunrepentantgeek/crddoc/actions/workflows/pr-validation.yml/badge.svg)](https://github.com/theunrepentantgeek/crddoc/actions/workflows/pr-validation.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/theunrepentantgeek/crddoc)](https://goreportcard.com/report/github.com/theunrepentantgeek/crddoc)


Inspired by the [gen-crd-api-reference-docs](https://github.com/ahmetb/gen-crd-api-reference-docs) project, this tool generates documentation for Kubernetes Custom Resource Definitions (CRDs) by parsing comments from the source code.

Over on the [Azure Service Operator](https://github.com/azure/azure-service-operator) (ASO) project, we've made use of the  [gen-crd-api-reference-docs](https://github.com/ahmetb/gen-crd-api-reference-docs) to generate documentation for our [Supported Resources](https://azure.github.io/azure-service-operator/reference/).  We've made a few contributions back to the tool, but unfortunately, the author updated the [ReadMe](https://github.com/ahmetb/gen-crd-api-reference-docs#alternatives) in June 2023 to indicate the project is no longer actively maintained.

## Project Status

Beta status, working towards a v1.0 release. Should be stable enough to use for production purposes. No breaking changes are planned.

## Quick Start

Install **crddoc** from source:

``` bash
go install github.com/theunrepentantgeek/crddoc@latest
```

Go to the folder containing the Go source for your CRDs, and generate documentation:

``` bash
crddoc document crds --output crd.md .
```

Dcoumentation will be generated into `crd.md`.

## Features

It's been said [**Performance** is a feature]() and **crddoc** generates full documentation in just seconds.

**Built-in templates** make for a quick start, but you aren't forced to use them. Use `crd export templates` to save the built in templates and customize them to suit your needs. These are standard [Go templates]() and the **crddoc** object model is [fully documented]().

**Markdown output** for maximum compatibility, and it's properly formatted to make for easier reading even before it is rendered as HTML.

[**Editor Support**]() allows you to use regular expressions to make precise edits as your documentation is generated. For example, inserting a zero-width-breaking-space after each forward-slash (`/`) in an ARM ID allows those IDs to [wordwrap properly within table cells]() instead of pushing the table wider than the viewport. 

With **Type Filters** you can elide specific types from the output, allowing your documentation to focus on what's important. This is particularly useful when you have helper types that are a implementation detail, and don't really form a part of the CRDs.

## Installation

### From source

``` bash
go install github.com/theunrepentantgeek/crddoc@latest
```

## Commandline reference

### Generate CRD documentation

Scan a Go package and generate CRD documentation as a markdown file.

``` bash
crddoc document crds --output package.md ./package
```

### Export templates

The built in templates can be easily exported to a folder, allowing you to customize how the generation process works.

``` bash
crddoc export templates --folder ./templates
```

Once you've customized the templates, use the `--template` flag when generating the docs.

``` bash
crddoc document crds --output package.md --template ./templates ./package
```

### Export configuration

Exporting the default configuration of **crddoc** is a good place to start customizing its behaviour.

``` bash
crddoc export configuration --output crddoc.yaml
```

After changing the configuration, use the `--config` flag when generating the docs.

``` bash
crddoc document crds --output package.md --config crddoc.yaml ./package
```

## Configuration reference

