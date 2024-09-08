package strumbra_test

import (
	"testing"

	"github.com/garrettladley/strumbra"
)

func TestEqualDifferentString(t *testing.T) {
	t.Parallel()

	a, err := strumbra.New("hello")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	b, err := strumbra.New("world")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if a.Equal(b) {
		t.Fatal("expected strings to be different.")
	}

	if b.Equal(a) {
		t.Fatal("expected strings to be different.")
	}
}

func TestEqualSameString(t *testing.T) {
	t.Parallel()

	a, err := strumbra.New("")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	b, err := strumbra.New("")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !a.Equal(b) {
		t.Fatal("expected strings to be the same.")
	}

	if !b.Equal(a) {
		t.Fatal("expected strings to be the same.")
	}
}
