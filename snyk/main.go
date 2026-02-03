package main

import "context"

type Snyk struct{}

func New() *Snyk {
	return &Snyk{}
}

func (m *Snyk) Test(ctx context.Context) (string, error) {
	return "i am running snyk test commands in snyk module\n", nil
}

func (m *Snyk) Monitor(ctx context.Context) (string, error) {
	return "i am running snyk monitor commands in snyk module\n", nil
}
