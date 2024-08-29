// +build tools
//go:build tools

// Package tools tracks dependencies for tools that used in the build process.
// See https://github.com/golang/go/wiki/Modules
package tools

import (
	_ "github.com/google/addlicense"
	_ "golang.org/x/perf/cmd/benchstat"
	_ "k8s.io/kube-openapi/cmd/openapi-gen"
	_ "sigs.k8s.io/logtools/logcheck"
	_ "sigs.k8s.io/mdtoc"
)
