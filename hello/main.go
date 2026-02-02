package main

import (
	"context"
	"fmt"

	"dagger.io/dagger"
)

type Hello struct{}

// Called as:
// dagger call hello --release-version=... --tag-prefix=... --dry-run=... --github-token=...
func (m *Hello) Hello(
	ctx context.Context,
	ReleaseVersion string,
	TagPrefix string,
	DryRun bool,
	GithubToken dagger.Secret,
) (string, error) {
	tokenProvided := "no"
	if GithubToken != nil {
		tokenProvided = "yes"
	}

	return fmt.Sprintf(
		"hello from dagger module!\nrelease_version=%s\ntag_prefix=%s\ndry_run=%v\ngithub_token_provided=%s\n",
		ReleaseVersion, TagPrefix, DryRun, tokenProvided,
	), nil
}
