package midi

import (
	"os"
	"path/filepath"
	"testing"
)

var (
	bigdata   []byte
	pathToMid string
)

func setup() {
	bigdata = make([]byte, 0x10000000)
	pathToMid = filepath.Join("testdata", "vegetable_valley.mid")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}
