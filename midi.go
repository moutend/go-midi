package midi

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
