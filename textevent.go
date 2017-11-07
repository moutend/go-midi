package midi

import "fmt"

// TextEvent corresponds to text event (0xff01) in MIDI.
type TextEvent struct {
	deltaTime *DeltaTime
	text      []byte
}

// DeltaTime returns delta time of this event as DeltaTime.
func (e *TextEvent) DeltaTime() *DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &DeltaTime{}
	}
	return e.deltaTime
}

func (e *TextEvent) Text() []byte {
	if e.text == nil {
		e.text = []byte{}
	}

	return e.text
}

// String returns string representation of this event.
func (e *TextEvent) String() string {
	return fmt.Sprintf("&TextEvent{text: \"%v\"}", string(e.text))
}

// Serialize serializes this event.
func (e *TextEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, e.DeltaTime().Quantity().Value()...)
	bs = append(bs, 0xff, 0x01)
	bs = append(bs, byte(len(e.Text())))
	bs = append(bs, e.Text()...)

	return bs
}
