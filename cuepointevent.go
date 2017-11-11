package midi

import "fmt"

// CuePointEvent corresponds to cue point event.
type CuePointEvent struct {
	deltaTime *DeltaTime
	text      []byte
}

// DeltaTime returns delta time of cue point event.
func (e *CuePointEvent) DeltaTime() *DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &DeltaTime{}
	}
	return e.deltaTime
}

// String returns string representation of cue point event.
func (e *CuePointEvent) String() string {
	return fmt.Sprintf("&CuePointEvent{text: \"%v\"}", string(e.Text()))
}

// Serialize serializes cue point event.
func (e *CuePointEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, e.DeltaTime().Quantity().Value()...)
	bs = append(bs, Meta, CuePoint)
	bs = append(bs, byte(len(e.Text())))
	bs = append(bs, e.Text()...)

	return bs
}

// SetText sets text.
func (e *CuePointEvent) SetText(text []byte) error {
	if len(text) > 127 {
		return fmt.Errorf("midi: maximum length of text is 127 bytes")
	}
	e.text = text

	return nil
}

// Text returns text.
func (e *CuePointEvent) Text() []byte {
	if e.text == nil {
		e.text = []byte{}
	}

	return e.text
}

// NewCuePointEvent returns CuePointEvent with the given parameter.
func NewCuePointEvent(deltaTime *DeltaTime, text []byte) (*CuePointEvent, error) {
	var err error

	event := &CuePointEvent{}
	event.deltaTime = deltaTime

	err = event.SetText(text)
	if err != nil {
		return nil, err
	}
	return event, nil
}
