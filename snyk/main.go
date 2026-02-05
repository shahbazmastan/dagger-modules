package main

import (
	"fmt"

	"github.com/shahbazmastan/dagger-modules/snyk/internal/dagger"
)

type Snyk struct {
	Src      *dagger.Directory
	Greeting string
	Name     string
	Age      int
}

// New is the constructor.
// Its parameters become flags on `dagger call` (NOT `dagger call new`).
func New(
	// Source directory to work with
	src *dagger.Directory,
	name string,
	age int,
	// Greeting message
	// +optional
	// +default="Hello"
	greeting string,
) *Snyk {
	if greeting == "" {
		greeting = "Hello" // defensive fallback
	}

	return &Snyk{
		Src:      src,
		Greeting: greeting,
		Name:     name,
		Age:      age,
	}
}

func (m *Snyk) Monitor() {
	fmt.Println(m.Greeting + ", " + m.Name + "! You are " + string(rune(m.Age)) + " years old.")

}

func (m *Snyk) Test() {
	fmt.Println(m.Greeting + ", " + m.Name + "! You are " + string(rune(m.Age)) + " years old.")
}

// func (m *Snyk) Test(ctx context.Context) (string, error) {
// 	// Just demonstrating that we have Src now
// 	return "i am running snyk test commands in snyk module (src provided)\n", nil
// }

// func (m *Snyk) Monitor(ctx context.Context) (string, error) {
// 	return "i am running snyk monitor commands in snyk module (src provided)\n", nil
// }
