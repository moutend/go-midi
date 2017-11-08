package midi

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type Track struct {
	Events []Event
}

func parseTrack(stream []byte) (*Track, error) {
	var start int

	sizeOfStream := len(stream)
	track := &Track{
		Events: []Event{},
	}
	for {
		if start >= sizeOfStream {
			break
		}
		event, sizeOfEvent, err := parseEvent(stream[start:])
		if err != nil {
			return nil, err
		}
		track.Events = append(track.Events, event)
		start += sizeOfEvent

		switch event.(type) {
		case *EndOfTrackEvent:
			return track, nil
		}
	}

	return nil, fmt.Errorf("midi: missing end of track event")
}

func parseTracks(stream []byte, numberOfTracks int) ([]*Track, error) {
	const MTrk uint32 = 0x4d54726B
	var chunkId uint32
	var start int64
	var chunkSize uint32

	tracks := make([]*Track, numberOfTracks)

	for n := 0; n < numberOfTracks; n++ {
		data := bytes.NewReader(stream[start:])
		binary.Read(io.NewSectionReader(data, 0, 4), binary.BigEndian, &chunkId)
		if chunkId != MTrk {
			return nil, fmt.Errorf("midi: invalid track ID %v", chunkId)
		}

		binary.Read(io.NewSectionReader(data, 4, 4), binary.BigEndian, &chunkSize)
		track, err := parseTrack(stream[start+8:])
		if err != nil {
			return nil, err
		}
		tracks[n] = track
		start += int64(chunkSize + 8)
	}

	return tracks, nil
}
