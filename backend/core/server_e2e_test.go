//go:build e2e

package core

import "testing"

func TestRunServer(t *testing.T) {
	InitViper("../etc/config.yaml")
	RunServer()
}
