package main

import (
	"context"
	"fmt"

	"github.com/shahbazmastan/dagger-modules/snyk/internal/dagger"
)

type Snyk struct {
	Src      *dagger.Directory
	Greeting string
	Name     string
	Age      int
}

func New(
	src *dagger.Directory,
	name string,
	age int,
	// Greeting message
	// +optional
	// +default="Hello"
	greeting string,
) *Snyk {
	if greeting == "" {
		greeting = "Hello"
	}
	return &Snyk{Src: src, Greeting: greeting, Name: name, Age: age}
}

func (m *Snyk) Monitor(ctx context.Context) (string, error) {
	// Use fmt.Sprintf (Age prints correctly)
	msg := fmt.Sprintf("%s, %s! You are %d years old.", m.Greeting, m.Name, m.Age)
	return msg + "\n", nil
}

func (m *Snyk) Test(ctx context.Context) (string, error) {
	msg := fmt.Sprintf("%s, %s! You are %d years old.", m.Greeting, m.Name, m.Age)
	return msg + "\n", nil
}
