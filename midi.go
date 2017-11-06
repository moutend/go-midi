package midi

type MIDI struct {
	Header *Header
	Tracks []*Track
}

func Parse(stream []byte) (*MIDI, error) {
	midi := &MIDI{}
	header, err := parseHeader(stream)
	if err != nil {
		return nil, err
	}
	tracks, err := parseTracks(stream, int(header.tracks))
	if err != nil {
		return nil, err
	}

	midi.Header = header
	midi.Tracks = tracks

	return midi, nil
}
