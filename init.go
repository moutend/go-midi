package midi

import (
	"io/ioutil"
	"log"
)

var logger *log.Logger

func init() {
	logger = log.New(ioutil.Discard, "discard logging messages", log.LstdFlags)
}
