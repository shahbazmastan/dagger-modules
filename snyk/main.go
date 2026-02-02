package main

import (
	"context"
)

type Snyk struct{}

// dagger call test
func (m *Snyk) Test(ctx context.Context) (string, error) {
	return "i am running snyk test commands in snyk module\n", nil
}

// dagger call monitor
func (m *Snyk) Monitor(ctx context.Context) (string, error) {
	return "i am running snyk monitor commands in snyk module\n", nil
}
