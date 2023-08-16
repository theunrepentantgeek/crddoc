# crddoc

Inspired by the [gen-crd-api-reference-docs](https://github.com/ahmetb/gen-crd-api-reference-docs) project, this tool generates documentation for Kubernetes Custom Resource Definitions (CRDs) by parsing comments from the source code.

## Project Status

This is a proof of concept, not yet ready for production use - but feel free to try it out and provide feedback!

## Motivation

Over on the [Azure Service Operator](https://github.com/azure/azure-service-operator) (ASO) project, we've made use of the  [gen-crd-api-reference-docs](https://github.com/ahmetb/gen-crd-api-reference-docs) to generate documentation for our [Supported Resources](https://azure.github.io/azure-service-operator/reference/).  We've made a few contributions back to the tool, but unfortunately, the author updated the [ReadMe](https://github.com/ahmetb/gen-crd-api-reference-docs#alternatives) in June 2023 to indicate the project is no longer actively maintained.

In lieu of adopting one of the other available tools, I thought it might be an interesting personal side project to implement a replacement.
