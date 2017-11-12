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

	q := &Quantity{}
	q.SetUint32(uint32(len(e.Text())))
	bs = append(bs, q.Value()...)
	bs = append(bs, e.Text()...)

	return bs
}

// SetText sets text.
func (e *CuePointEvent) SetText(text []byte) error {
	if len(text) > 0xfffffff {
		return fmt.Errorf("midi: maximum size of text is 256 MB")
	}
	e.text = text

	return nil
}

// Text returns text.
func (e *CuePointEvent) Text() []byte {
	if e.text == nil {
		e.text = []byte{}
	}

	text := make([]byte, len(e.text))
	copy(text, e.text)

	return text
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
