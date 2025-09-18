package generate

import (
	"context"
)

// In a terminal, run `go generate` in this directory to have
// it generates the generated.go file.

//go:generate moq -out generated.go . MyInterface

// MyInterface is a test interface.
type MyInterface[A ~int, B ~string, C ~float64] interface {
	One(ctx context.Context, in A) bool
	Two(id string, in B) string
	Three(count int) C
}
