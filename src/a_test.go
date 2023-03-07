package a

import (
	"testing"
	"github.com/bazelbuild/rules_go/go/runfiles"
)

func TestA(t *testing.T) {
	var s = "asdf"
	var r = A(s)
	if (s != r) {
		t.Fatalf("Expected equal strings")
	}
}
