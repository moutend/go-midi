package main

import (
	"io/ioutil"
	"log"

	midi "github.com/moutend/go-midi"
	"github.com/moutend/go-midi/constant"
	"github.com/moutend/go-midi/deltatime"
	"github.com/moutend/go-midi/event"
)

func main() {
	deltaTime1, _ := deltatime.New(0)
	deltaTime2, _ := deltatime.New(960)

	noteOnC4, _ := event.NewNoteOnEvent(deltaTime1, 0, constant.C4, 127)
	noteOnE4, _ := event.NewNoteOnEvent(deltaTime1, 0, constant.E4, 127)
	noteOnG4, _ := event.NewNoteOnEvent(deltaTime1, 0, constant.G4, 127)

	noteOffC4, _ := event.NewNoteOffEvent(deltaTime2, 0, constant.C4, 0)
	noteOffE4, _ := event.NewNoteOffEvent(deltaTime1, 0, constant.E4, 0)
	noteOffG4, _ := event.NewNoteOffEvent(deltaTime1, 0, constant.G4, 0)

	endOfTrack, _ := event.NewEndOfTrackEvent(deltaTime2)

	t := midi.NewTrack(
		noteOnC4,
		noteOnE4,
		noteOnG4,
		noteOffC4,
		noteOffE4,
		noteOffG4,
		endOfTrack,
	)

	m := midi.MIDI{}
	m.TimeDivision().SetBPM(240)
	m.Tracks = append(m.Tracks, t)

	err := ioutil.WriteFile("output.mid", m.Serialize(), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
