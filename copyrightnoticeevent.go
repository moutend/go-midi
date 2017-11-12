package midi

import "fmt"

// CopyrightNoticeEvent corresponds to copyright notice event.
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
	return fmt.Sprintf("&CopyrightNoticeEvent{text: \"%v\"}", string(e.Text()))
}

// Serialize serializes copyright notice event.
func (e *CopyrightNoticeEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, e.DeltaTime().Quantity().Value()...)
	bs = append(bs, Meta, CopyrightNotice)

	q := &Quantity{}
	q.SetUint32(uint32(len(e.Text())))
	bs = append(bs, q.Value()...)
	bs = append(bs, e.Text()...)

	return bs
}

// SetText sets text.
func (e *CopyrightNoticeEvent) SetText(text []byte) error {
	if len(text) > 0xfffffff {
		return fmt.Errorf("midi: maximum size of text is 256 MB")
	}
	e.text = text

	return nil
}

// Text returns text.
func (e *CopyrightNoticeEvent) Text() []byte {
	if e.text == nil {
		e.text = []byte{}
	}

	text := make([]byte, len(e.text))
	copy(text, e.text)

	return text
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
