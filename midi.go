package midi

import (
	"fmt"
	"log"
)

// MIDI represents standard MIDI data.
type MIDI struct {
	formatType   uint16
	timeDivision *TimeDivision
	Tracks       []*Track
}

// Serialize serializes MIDI data.
func (m *MIDI) Serialize() []byte {
	stream := []byte("MThd")
	stream = append(stream, 0x00, 0x00, 0x00, 0x06)
	stream = append(stream, 0x00, byte(m.formatType&0xff))

	var numberOfTracks uint16
	if len(m.Tracks) > 0xffff {
		numberOfTracks = uint16(0xffff)
	} else {
		numberOfTracks = uint16(len(m.Tracks))
	}

	stream = append(stream, byte((numberOfTracks>>8)&0xff))
	stream = append(stream, byte(numberOfTracks&0xff))
	stream = append(stream, m.TimeDivision().Serialize()...)

	for _, track := range m.Tracks {
		stream = append(stream, track.Serialize()...)
	}

	return stream
}

// TimeDivision returns time division.
func (m *MIDI) TimeDivision() *TimeDivision {
	if m.timeDivision == nil {
		m.timeDivision = &TimeDivision{}
	}
	return m.timeDivision
}

// Parse parses standard MIDI (*.mid) stream.
func Parse(stream []byte) (*MIDI, error) {
	logger.parsedBytes = 0
	logger.Logger.Printf("midi: start parsing %v bytes\n", len(stream))

	formatType, numberOfTracks, timeDivision, err := parseHeader(stream)
	if err != nil {
		return nil, err
	}

	tracks, err := parseTracks(stream[14:], int(numberOfTracks))
	if err != nil {
		return nil, err
	}

	midi := &MIDI{
		formatType:   formatType,
		timeDivision: &TimeDivision{value: timeDivision},
		Tracks:       tracks,
	}

	logger.Println("successfully done")

	return midi, nil
}

// SetLogger sets logger for debugging.
func SetLogger(l *log.Logger) {
	if l != nil {
		logger.Logger = l
	}
}

// parseHeader parses stream begins with MThd.
func parseHeader(stream []byte) (formatType, numberOfTracks, timeDivision uint16, err error) {
	var start int

	logger.Println("start parsing MThd")

	if string(stream[:4]) != "MThd" {
		return formatType, numberOfTracks, timeDivision, fmt.Errorf("midi: invalid chunk ID %v", stream[:4])
	}

	start += 4
	logger.parsedBytes += 4
	logger.Println("parsing MThd completed")

	start += 4 // skip read header size
	logger.parsedBytes += 4
	logger.Println("skip parsing size of header chunk")

	logger.Println("start parsing format type")

	formatType = uint16(stream[start+1])

	start += 2
	logger.parsedBytes += 2
	logger.Printf("parsing format type completed (formatType=%v)", formatType)

	logger.Println("start parsing number of tracks")

	numberOfTracks = uint16(stream[start])
	numberOfTracks = numberOfTracks << 8
	numberOfTracks += uint16(stream[start+1])

	start += 2
	logger.parsedBytes += 2
	logger.Printf("parsing number of tracks completed (%v)", numberOfTracks)

	logger.Println("start parsing time division")

	timeDivision = uint16(stream[start])
	timeDivision = timeDivision << 8
	timeDivision += uint16(stream[start+1])

	start += 2
	logger.parsedBytes += 2
	logger.Printf("parsing time division completed (%v)", timeDivision)

	return formatType, numberOfTracks, timeDivision, nil
}
