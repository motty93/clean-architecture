// tools.go
// go:build tools
//go:build tools
// +build tools

package main

import (
	_ "github.com/cosmtrek/air"
	_ "github.com/golang/mock/mockgen"
)
