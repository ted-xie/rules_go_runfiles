package a

import (
	"os"
	"path"
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


func TestRunFiles(t *testing.T) {
	var a_txt = "src/a.txt"
	var rpath, err = runfiles.Rlocation(path.Join(os.Getenv("TEST_WORKSPACE"), a_txt))
	if err != nil {
		t.Fatalf("Could not find runfile at %v", rpath)
	}
}
