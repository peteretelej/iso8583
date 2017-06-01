package iso8583

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestServer(t *testing.T) {
	_, err := WebServer("localhost:9999")
	if err != nil {
		t.Errorf("WebServer returned wrong error value: expected nil, got %v", err)
	}
	tmpdir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("unable to get temp dir for testing: %v", err)
	}
	defer func() {
		if err := os.RemoveAll(tmpdir); err != nil {
			t.Fatalf("unable to delete tmpdir: %v", err)
		}
	}()
	if err := os.Mkdir(filepath.Join(tmpdir, "web"), 0755); err != nil {
		t.Fatalf("unable to create web/ in temp dir: %v", err)
	}
	if err := os.Chdir(tmpdir); err != nil {
		t.Fatalf("unable to change into tempdir: %v", err)
	}
	_, err = WebServer("localhost:9999")
	if err == nil {
		t.Fatalf("Webserver returned wrong error value, expected error, got %v", err)
	}
	if !strings.Contains(err.Error(), "missing web/") {
		t.Errorf("Webserver returned wrong error got %v, expected 'missing web/'", err)
	}
}
