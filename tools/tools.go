// Code generated by github.com/izumin5210/gex. DO NOT EDIT.

// +build tools

package tools

// tool dependencies
import (
	_ "github.com/go-swagger/go-swagger/cmd/swagger"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/izumin5210/gex/cmd/gex"
)

// If you want to use tools, please run the following command:
//  go generate ./tools.go
//
//go:generate go build -v -o=./bin/swagger github.com/go-swagger/go-swagger/cmd/swagger
//go:generate go build -v -o=./bin/mockgen github.com/golang/mock/mockgen
//go:generate go build -v -o=./bin/gex github.com/izumin5210/gex/cmd/gex
