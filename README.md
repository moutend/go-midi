go-midi
========

`go-midi` provides reading and writing standard MIDI file.

## Installation

```console
go get -u github.com/moutend/go-midi
```

## Examples

### Print All Text Events

The following program reads a file named `music.mid` and then prints all text data.

```go
package main

import (
	"fmt"
	"io/ioutil"

	midi "github.com/moutend/go-midi"
)

func main() {
	file, err := ioutil.ReadFile("music.mid")
	if err != nil {
		panic(err)
	}

	midiData, err := midi.Parse(file)
	if err != nil {
		panic(err)
	}

	for _, track := range midiData.Tracks {
		for _, event := range track.Events {
			switch event.(type) {
			case *midi.TextEvent:
				fmt.Printf("%s\n", event.(*midi.TextEvent).Text())
			}
		}
	}
}
```

### Normalize Velocity of Note-on Event

The following program reads a file named `input.mid`, set the velocity of all note-on events to `127` and creates a new file named `output.mid`.

```go
package main

import (
	"io/ioutil"

	midi "github.com/moutend/go-midi"
)

func main() {
	file, err := ioutil.ReadFile("input.mid")
	if err != nil {
		panic(err)
	}

	midiData, err := midi.Parse(file)
	if err != nil {
		panic(err)
	}

	for _, track := range midiData.Tracks {
		for _, event := range track.Events {
			switch event.(type) {
			case *midi.NoteOnEvent:
				event.(*midi.NoteOnEvent).SetVelocity(127)
			}
		}
	}
	err = ioutil.WriteFile("output.mid", midiData.Serialize(), 0644)
	if err != nil {
		panic(err)
	}
}
```

### Create C Major Chord

The following program creates a MIDI file which plays C major chord.

```go
package main

import (
	"io/ioutil"

	midi "github.com/moutend/go-midi"
)

func main() {
	deltaTime1, _ := midi.NewDeltaTime(0)
	deltaTime2, _ := midi.NewDeltaTime(960)

	noteOnC4, _ := midi.NewNoteOnEvent(deltaTime1, 0, midi.C4, 127)
	noteOnE4, _ := midi.NewNoteOnEvent(deltaTime1, 0, midi.E4, 127)
	noteOnG4, _ := midi.NewNoteOnEvent(deltaTime1, 0, midi.G4, 127)

	noteOffC4, _ := midi.NewNoteOffEvent(deltaTime2, 0, midi.C4, 0)
	noteOffE4, _ := midi.NewNoteOffEvent(deltaTime1, 0, midi.E4, 0)
	noteOffG4, _ := midi.NewNoteOffEvent(deltaTime1, 0, midi.G4, 0)

	endOfTrack, _ := midi.NewEndOfTrackEvent(deltaTime2)

	track := &midi.Track{
		Events: make([]midi.Event, 0),
	}

	track.Events = append(track.Events, noteOnC4)
	track.Events = append(track.Events, noteOnE4)
	track.Events = append(track.Events, noteOnG4)

	track.Events = append(track.Events, noteOffC4)
	track.Events = append(track.Events, noteOffE4)
	track.Events = append(track.Events, noteOffG4)

	track.Events = append(track.Events, endOfTrack)

	m := midi.MIDI{}
	m.TimeDivision().SetBPM(240)
	m.Tracks = make([]*midi.Track, 1)
	m.Tracks[0] = track

	ioutil.WriteFile("oo.mid", m.Serialize(), 0644)
}
```

## Numbering of nNotes

There are two conventions for notes in MIDI. The most common is where C3 is `0x3c` and the another is where C4 is `0x3c`. In this package, where C3 is `0x3c`.

## MIDI Files for Testing

The MIDI files located at `testdata` were composed by Nao. Check her great works:

- [星のカービィ　MIDI - みかんの旅](http://mikannotabi.blog31.fc2.com/blog-entry-6.html)

## About MIDI

- [MIDI File Format Specifications · colxi/midi-parser-js Wiki](https://github.com/colxi/midi-parser-js/wiki/MIDI-File-Format-Specifications)
- [0xff21 and 0xff20 in MIDI](https://groups.google.com/forum/#!topic/comp.music.midi/_MIjgi-8xQQ)
- [OpenMIDIProject - Documentations](http://openmidiproject.osdn.jp/documentations_en.html)

## Contributing

1. Fork ([https://github.com/moutend/go-midi/fork](https://github.com/moutend/go-midi/fork))
1. Create a feature branch
1. Add changes
1. Run `go fmt`
1. Commit your changes
1. Open a new Pull Request

## LICENSE

MIT

## Author

[Yoshiyuki Koyanagi](https://github.com/moutend)
