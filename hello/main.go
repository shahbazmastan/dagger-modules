package main

import (
	"context"
	"fmt"

	// This is the GENERATED package created by `dagger develop`
	"github.com/shahbazmastan/dagger-modules/hello/internal/dagger"
)

type Hello struct{}

// This becomes:
// dagger call hello --release-version=... --tag-prefix=... --dry-run=... --github-token=...
func (m *Hello) Hello(
	ctx context.Context,
	ReleaseVersion string,
	TagPrefix string,
	DryRun bool,
	GithubToken dagger.Secret,
) (string, error) {

	return fmt.Sprintf(
		"hello from dagger module!\nrelease_version=%s\ntag_prefix=%s\ndry_run=%v\ngithub_token_received=yes\n",
		ReleaseVersion,
		TagPrefix,
		DryRun,
	), nil
}

func (m *Hello) Snyk(ctx context.Context) (string, error) {
	return "i am running snyk commands\n", nil
}

func (m *Hello) SNYKMONITOR(ctx context.Context) (string, error) {
	return "i am running snyk monitor commands\n", nil
}
