package midi

import "fmt"

// Track represents MIDI track.
type Track struct {
	Events []Event
}

// Serialize serializes track.
func (t *Track) Serialize() []byte {
	data := []byte{}
	for _, event := range t.Events {
		data = append(data, event.Serialize()...)
	}

	sizeOfData := uint32(len(data))
	stream := []byte{0x4d, 0x54, 0x72, 0x6B} // MTrk
	stream = append(stream, byte(sizeOfData>>24))
	stream = append(stream, byte((sizeOfData&0xff0000)>>16))
	stream = append(stream, byte((sizeOfData&0xff00)>>8))
	stream = append(stream, byte(sizeOfData&0xff))
	stream = append(stream, data...)

	return stream
}

func parseTrack(stream []byte) (*Track, error) {
	start := 0
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

// parseTracks parses stream begins with MTrk.
func parseTracks(stream []byte, numberOfTracks int) ([]*Track, error) {
	start := 0
	tracks := make([]*Track, numberOfTracks)

	for n := 0; n < numberOfTracks; n++ {
		logger.Println("start parsing MTrk", start)
		if string(stream[start:start+4]) != "MTrk" {
			return nil, fmt.Errorf("midi: invalid track ID %v", stream[start:start+4])
		}

		start += 4
		logger.parsedBytes += 4
		logger.Println("parsing MTrk completed")

		logger.Println("start parsing size of track")

		chunkSize := uint32(stream[start])
		chunkSize = chunkSize << 8
		chunkSize += uint32(stream[start+1])
		chunkSize = chunkSize << 8
		chunkSize += uint32(stream[start+2])
		chunkSize = chunkSize << 8
		chunkSize += uint32(stream[start+3])

		start += 4
		logger.parsedBytes += 4
		logger.Printf("parsing size of track completed (chunkSize=%v)", chunkSize)

		track, err := parseTrack(stream[start:])
		if err != nil {
			return nil, err
		}

		tracks[n] = track
		start += int(chunkSize)
	}

	return tracks, nil
}
