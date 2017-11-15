package midi

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var (
	bigdata    []byte
	pathsToMid []string
)

func walkFunc(path string, info os.FileInfo, err error) error {
	if strings.HasSuffix(path, ".mid") {
		pathsToMid = append(pathsToMid, path)
	}
	return nil
}

func setup() {
	bigdata = make([]byte, 0x10000000)
	filepath.Walk("testdata", walkFunc)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}
