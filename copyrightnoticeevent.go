package midi

import "fmt"

// CopyrightNoticeEvent corresponds to copyright notice meta event.
type CopyrightNoticeEvent struct {
	deltaTime *DeltaTime
	text      []byte
}

// DeltaTime returns delta time of copyright notice event.
func (e *CopyrightNoticeEvent) DeltaTime() *DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &DeltaTime{}
	}
	return e.deltaTime
}

// String returns string representation of copyright notice event.
func (e *CopyrightNoticeEvent) String() string {
	return fmt.Sprintf("&CopyrightNoticeEvent{text: \"%v\"}", string(e.text))
}

// Serialize serializes copyright notice event.
func (e *CopyrightNoticeEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, e.DeltaTime().Quantity().Value()...)
	bs = append(bs, Meta, CopyrightNotice)
	bs = append(bs, byte(len(e.Text())))
	bs = append(bs, e.Text()...)

	return bs
}

// SetText sets text.
func (e *CopyrightNoticeEvent) SetText(text []byte) error {
	if len(text) > 127 {
		return fmt.Errorf("midi: maximum length of text is 127 bytes")
	}
	e.text = text

	return nil
}

// Text returns text.
func (e *CopyrightNoticeEvent) Text() []byte {
	if e.text == nil {
		e.text = []byte{}
	}

	return e.text
}

// NewCopyrightNoticeEvent returns CopyrightNoticeEvent with the given parameter.
func NewCopyrightNoticeEvent(deltaTime *DeltaTime, text []byte) (*CopyrightNoticeEvent, error) {
	var err error

	event := &CopyrightNoticeEvent{}
	event.deltaTime = deltaTime

	err = event.SetText(text)
	if err != nil {
		return nil, err
	}
	return event, nil
}
