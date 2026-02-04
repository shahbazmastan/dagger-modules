package main

import (
	"context"

	"github.com/shahbazmastan/dagger-modules/snyk/internal/dagger"
)

type Snyk struct {
	Src *dagger.Directory
}

// New is the constructor.
// Its parameters become flags on `dagger call` (NOT `dagger call new`).
func New(
	// Source directory to work with
	src *dagger.Directory,
) *Snyk {
	return &Snyk{Src: src}
}

func (m *Snyk) Test(ctx context.Context) (string, error) {
	// Just demonstrating that we have Src now
	return "i am running snyk test commands in snyk module (src provided)\n", nil
}

func (m *Snyk) Monitor(ctx context.Context) (string, error) {
	return "i am running snyk monitor commands in snyk module (src provided)\n", nil
}
