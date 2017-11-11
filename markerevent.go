package midi

import "fmt"

// MarkerEvent corresponds to marker meta event.
type MarkerEvent struct {
	deltaTime *DeltaTime
	text      []byte
}

// DeltaTime returns delta time of marker event.
func (e *MarkerEvent) DeltaTime() *DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &DeltaTime{}
	}
	return e.deltaTime
}

// String returns string representation of marker event.
func (e *MarkerEvent) String() string {
	return fmt.Sprintf("&MarkerEvent{text: \"%v\"}", string(e.text))
}

// Serialize serializes marker event.
func (e *MarkerEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, e.DeltaTime().Quantity().Value()...)
	bs = append(bs, Meta, Marker)
	bs = append(bs, byte(len(e.Text())))
	bs = append(bs, e.Text()...)

	return bs
}

// SetText sets text.
func (e *MarkerEvent) SetText(text []byte) error {
	if len(text) > 127 {
		return fmt.Errorf("midi: maximum length of text is 127 bytes")
	}
	e.text = text

	return nil
}

// Text returns text.
func (e *MarkerEvent) Text() []byte {
	if e.text == nil {
		e.text = []byte{}
	}

	return e.text
}

// NewMarkerEvent returns MarkerEvent with the given parameter.
func NewMarkerEvent(deltaTime *DeltaTime, text []byte) (*MarkerEvent, error) {
	var err error

	event := &MarkerEvent{}
	event.deltaTime = deltaTime

	err = event.SetText(text)
	if err != nil {
		return nil, err
	}
	return event, nil
}
