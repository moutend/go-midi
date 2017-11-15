package midi

import (
	"fmt"
	"io/ioutil"
	"log"
)

type midiLogger struct {
	*log.Logger
	parsedBytes int
}

func (l *midiLogger) Printf(format string, v ...interface{}) {
	format = fmt.Sprintf("midi: [%v] %v", l.parsedBytes, format)
	l.Logger.Printf(format, v...)
}

func (l *midiLogger) Println(v ...interface{}) {
	a := make([]interface{}, len(v)+1)
	a[0] = fmt.Sprintf("midi: [%v]", l.parsedBytes)

	for i := 0; i < len(v); i++ {
		a[i+1] = v[i]
	}

	l.Logger.Println(a...)
}

func newMIDILogger() *midiLogger {
	return &midiLogger{
		log.New(ioutil.Discard, "discard logging messages", 0),
		0,
	}
}
