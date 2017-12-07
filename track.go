package midi

// Track represents MIDI track.
type Track struct {
	Events []Event
}

// Serialize serializes track.
func (t *Track) Serialize() []byte {
	data := []byte{}

	for _, event := range t.Events {
		b := event.Serialize()

		if event.RunningStatus() {
			dt := event.DeltaTime().Quantity().Value()
			sizeOfDeltaTime := len(dt)
			switch b[sizeOfDeltaTime] {
			case 0xff:
				data = append(data, b[2:]...)
			default:
				data = append(data, dt...)
				data = append(data, b[sizeOfDeltaTime+1:]...)
			}
		} else {
			data = append(data, b...)
		}
	}

	stream := []byte("MTrk")

	sizeOfData := uint32(len(data))
	stream = append(stream, byte(sizeOfData>>24))
	stream = append(stream, byte((sizeOfData&0xff0000)>>16))
	stream = append(stream, byte((sizeOfData&0xff00)>>8))
	stream = append(stream, byte(sizeOfData&0xff))

	stream = append(stream, data...)

	return stream
}
