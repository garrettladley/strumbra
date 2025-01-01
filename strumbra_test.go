package strumbra

import "testing"

func TestEqualDifferentString(t *testing.T) {
	t.Parallel()

	a, err := New("hello")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	b, err := New("world")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if a.Equals(b) {
		t.Fatal("expected strings to be different.")
	}

	if b.Equals(a) {
		t.Fatal("expected strings to be different.")
	}
}

func TestEqualSameString(t *testing.T) {
	t.Parallel()

	a, err := New("")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	b, err := New("")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !a.Equals(b) {
		t.Fatal("expected strings to be the same.")
	}

	if !b.Equals(a) {
		t.Fatal("expected strings to be the same.")
	}
}
