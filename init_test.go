package midi

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var pathsToMid []string

func walkFunc(path string, info os.FileInfo, err error) error {
	if strings.HasSuffix(path, ".mid") {
		pathsToMid = append(pathsToMid, path)
	}
	return nil
}

func setup() {
	filepath.Walk("testdata", walkFunc)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}
