package midi

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
