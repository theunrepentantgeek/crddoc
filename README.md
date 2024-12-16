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

Documentation will be generated into `crd.md`.

## Features

It's been said [**Performance** is a feature](https://blog.codinghorror.com/performance-is-a-feature/) and **crddoc** generates full documentation in just seconds.

**Built-in templates** for a quick start, but you aren't forced to use them. Use `crddoc export templates` to save the built in templates and customize them to suit your needs. These are standard [Go templates](https://pkg.go.dev/text/template).

**Markdown output** for maximum compatibility, and it's properly formatted (pretty printed) to make for easier reading even before it is rendered as HTML.

For projects using [controller-gen](https://github.com/kubernetes-sigs/kubebuilder/blob/master/docs/book/src/reference/controller-gen.md), **crddoc** will look for relevant [marker comments](https://book.kubebuilder.io/reference/markers) and include them in the model that drives template generation. 

[**Automated editor support**]() allows you to make precise edits as your documentation is generated. For example, these enable insertion of a zero-width-breaking-space after each forward-slash (`/`) in an ARM ID to allow those IDs to [wordwrap properly within table cells]() instead of pushing the table wider than the viewport. 

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

After customization, use the `--template` flag to reference them when generating documentation.

``` bash
crddoc document crds --output package.md --template ./templates ./package
```

### Export configuration

Exporting the default configuration of **crddoc** is a good place to start customizing its behaviour.

``` bash
crddoc export configuration --output crddoc.yaml
```

After changing the configuration, use the `--config` flag when generating documentation.

``` bash
crddoc document crds --output package.md --config crddoc.yaml ./package
```

For possible options, see the [configuration reference](https://github.com/theunrepentantgeek/crddoc/blob/main/docs/config/config.md).

# Contributing

Contributions are welcome.

## Devcontainer

The easiest way to get started is to use Visual Studio Code (or any other IDE that supports devcontainers), as all the necessary tools will automatically be installed.

## Setup

If you don't want to (or can't) use the devcontainer, you have two options.

The simplest is to use just the Go CLI, as the entire project can be built with `go build` and tested with `go test ./...`. 

Alternatively, you can use the `install-dependencies.sh` script to install all the recommended tools. I suggest you review the script before running it. 

## Linters

I'm a big believer in the value of static analysis tools (aka linters) to discover potential bugs and other issues early in the development process, and this project uses a considerable number through [golangci-lint](https://golangci-lint.run/). These linters are run automatically for all pull-requests. To run the same checks yourself, use `task lint` within the devcontainer. 