package midi

import "fmt"

// SequenceOrTrackNameEvent corresponds to sequence or track name event.
type SequenceOrTrackNameEvent struct {
	deltaTime *DeltaTime
	text      []byte
}

// DeltaTime returns delta time of sequence or track name event.
func (e *SequenceOrTrackNameEvent) DeltaTime() *DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &DeltaTime{}
	}
	return e.deltaTime
}

// String returns string representation of sequence or track name event.
func (e *SequenceOrTrackNameEvent) String() string {
	return fmt.Sprintf("&SequenceOrTrackNameEvent{text: \"%v\"}", string(e.Text()))
}

// Serialize serializes sequence or track name event.
func (e *SequenceOrTrackNameEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, e.DeltaTime().Quantity().Value()...)
	bs = append(bs, Meta, SequenceOrTrackName)
	bs = append(bs, byte(len(e.Text())))
	bs = append(bs, e.Text()...)

	return bs
}

// SetText sets text.
func (e *SequenceOrTrackNameEvent) SetText(text []byte) error {
	if len(text) > 127 {
		return fmt.Errorf("midi: maximum length of text is 127 bytes")
	}
	e.text = text

	return nil
}

// Text returns text.
func (e *SequenceOrTrackNameEvent) Text() []byte {
	if e.text == nil {
		e.text = []byte{}
	}

	return e.text
}

// NewSequenceOrTrackNameEvent returns SequenceOrTrackNameEvent with the given parameter.
func NewSequenceOrTrackNameEvent(deltaTime *DeltaTime, text []byte) (*SequenceOrTrackNameEvent, error) {
	var err error

	event := &SequenceOrTrackNameEvent{}
	event.deltaTime = deltaTime

	err = event.SetText(text)
	if err != nil {
		return nil, err
	}
	return event, nil
}
