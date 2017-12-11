package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/moutend/go-midi"
	"github.com/moutend/go-midi/event"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	m, err := midi.NewParser(file).Parse()
	if err != nil {
		log.Fatal(err)
	}

	for _, t := range m.Tracks {
		for _, e := range t.Events {
			switch e.(type) {
			case *event.NoteOnEvent:
				noteOn := e.(*event.NoteOnEvent)
				if noteOn.Velocity() > 0 {
					noteOn.SetVelocity(127)
				}
			}
		}
	}

	err = ioutil.WriteFile("output.mid", m.Serialize(), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
