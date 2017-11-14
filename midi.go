package midi

import "log"

// MIDI represents standard MIDI data.
type MIDI struct {
	Header *Header
	Tracks []*Track
}

// Serialize serializes MIDI data.
func (m *MIDI) Serialize() []byte {
	stream := []byte{}
	stream = append(stream, m.Header.Serialize()...)
	for _, track := range m.Tracks {
		stream = append(stream, track.Serialize()...)
	}

	return stream
}

// Parse parses standard MIDI (*.mid) stream.
func Parse(stream []byte) (*MIDI, error) {
	header, err := parseHeader(stream)
	if err != nil {
		return nil, err
	}

	tracks, err := parseTracks(stream[14:], int(header.tracks))
	if err != nil {
		return nil, err
	}

	midi := &MIDI{
		Header: header,
		Tracks: tracks,
	}

	return midi, nil
}

// SetLogger sets logger for debugging.
func SetLogger(l *log.Logger) {
	if l != nil {
		logger = l
	}
}
