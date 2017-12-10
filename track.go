package midi

import (
	"github.com/moutend/go-midi/constant"
	"github.com/moutend/go-midi/event"
)

// Track represents MIDI track.
type Track struct {
	Events []event.Event
}

// Serialize serializes track.
func (t *Track) Serialize() []byte {
	data := []byte{}

	for _, event := range t.Events {
		data = append(data, event.DeltaTime().Quantity().Value()...)

		if event.RunningStatus() {
			bs := event.Serialize()

			switch bs[0] {
			case constant.Meta:
				data = append(data, bs[2:]...)
			default:
				data = append(data, bs[1:]...)
			}
		} else {
			data = append(data, event.Serialize()...)
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
