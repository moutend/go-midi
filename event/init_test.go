package event

import (
	"os"
	"testing"
)

var bigdata []byte

func setup() {
	bigdata = make([]byte, 0x10000000)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}
