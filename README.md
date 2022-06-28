# openapi3 Examples

Couldn't really find any good getting started examples for kin-openapi/openapi3 and Spectral validation of an entire openapi document. Hope this helps someone. ❤

## Contents

* spec_test.go - Loads and validates an openapi3 spec using Go test, also runs a contract test against the spec
* main.go - Builds something very similar to the openapi.yaml spec programmatically using Go
* .spectral.yaml - Spectral config with default rules plus an example function that allows checking to make sure api version (v1) is in EITHER the server or paths section, but not both
* function/version.js - Spectral javascript function example
* package.json - Configuration to publish this spectral config as an npm module for consumption by client apps
* Makefile - Basic commands documented
