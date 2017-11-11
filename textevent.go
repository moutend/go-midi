package midi

import "fmt"

// TextEvent corresponds to text meta event.
type TextEvent struct {
	deltaTime *DeltaTime
	text      []byte
}

// DeltaTime returns delta time of text event.
func (e *TextEvent) DeltaTime() *DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &DeltaTime{}
	}
	return e.deltaTime
}

// String returns string representation of text event.
func (e *TextEvent) String() string {
	return fmt.Sprintf("&TextEvent{text: \"%v\"}", string(e.text))
}

// Serialize serializes text event.
func (e *TextEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, e.DeltaTime().Quantity().Value()...)
	bs = append(bs, Meta, Text)
	bs = append(bs, byte(len(e.Text())))
	bs = append(bs, e.Text()...)

	return bs
}

// SetText sets text.
func (e *TextEvent) SetText(text []byte) error {
	if len(text) > 127 {
		return fmt.Errorf("midi: maximum length of text is 127 bytes")
	}
	e.text = text

	return nil
}

// Text returns text.
func (e *TextEvent) Text() []byte {
	if e.text == nil {
		e.text = []byte{}
	}

	return e.text
}

// NewTextEvent returns TextEvent with the given parameter.
func NewTextEvent(deltaTime *DeltaTime, text []byte) (*TextEvent, error) {
	var err error

	event := &TextEvent{}
	event.deltaTime = deltaTime

	err = event.SetText(text)
	if err != nil {
		return nil, err
	}
	return event, nil
}
