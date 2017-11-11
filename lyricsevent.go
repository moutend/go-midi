package midi

import "fmt"

// LyricsEvent corresponds to lyrics event.
type LyricsEvent struct {
	deltaTime *DeltaTime
	text      []byte
}

// DeltaTime returns delta time of lyrics event.
func (e *LyricsEvent) DeltaTime() *DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &DeltaTime{}
	}
	return e.deltaTime
}

// String returns string representation of lyrics event.
func (e *LyricsEvent) String() string {
	return fmt.Sprintf("&LyricsEvent{text: \"%v\"}", string(e.Text()))
}

// Serialize serializes lyrics event.
func (e *LyricsEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, e.DeltaTime().Quantity().Value()...)
	bs = append(bs, Meta, Lyrics)
	bs = append(bs, byte(len(e.Text())))
	bs = append(bs, e.Text()...)

	return bs
}

// SetText sets text.
func (e *LyricsEvent) SetText(text []byte) error {
	if len(text) > 127 {
		return fmt.Errorf("midi: maximum length of text is 127 bytes")
	}
	e.text = text

	return nil
}

// Text returns text.
func (e *LyricsEvent) Text() []byte {
	if e.text == nil {
		e.text = []byte{}
	}

	return e.text
}

// NewLyricsEvent returns LyricsEvent with the given parameter.
func NewLyricsEvent(deltaTime *DeltaTime, text []byte) (*LyricsEvent, error) {
	var err error

	event := &LyricsEvent{}
	event.deltaTime = deltaTime

	err = event.SetText(text)
	if err != nil {
		return nil, err
	}
	return event, nil
}
