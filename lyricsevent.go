package midi

import "fmt"

// LyricsEvent corresponds to lyrics event.
type LyricsEvent struct {
	deltaTime     *DeltaTime
	runningStatus bool
	text          []byte
}

// DeltaTime returns delta time of lyrics event.
func (e *LyricsEvent) DeltaTime() *DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &DeltaTime{}
	}
	return e.deltaTime
}

// Serialize serializes lyrics event.
func (e *LyricsEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, e.DeltaTime().Quantity().Value()...)
	bs = append(bs, Meta, Lyrics)

	q := &Quantity{}
	q.SetUint32(uint32(len(e.Text())))
	bs = append(bs, q.Value()...)
	bs = append(bs, e.Text()...)

	return bs
}

// SetRunningStatus sets running status.
func (e *LyricsEvent) SetRunningStatus(status bool) {
	e.runningStatus = status
}

// RunningStatus returns running status.
func (e *LyricsEvent) RunningStatus() bool {
	return e.runningStatus
}

// SetText sets text.
func (e *LyricsEvent) SetText(text []byte) error {
	if len(text) > 0xfffffff {
		return fmt.Errorf("midi: maximum size of text is 256 MB")
	}
	e.text = text

	return nil
}

// Text returns text.
func (e *LyricsEvent) Text() []byte {
	if e.text == nil {
		e.text = []byte{}
	}

	text := make([]byte, len(e.text))
	copy(text, e.text)

	return text
}

// String returns string representation of lyrics event.
func (e *LyricsEvent) String() string {
	return fmt.Sprintf("&LyricsEvent{text: \"%v\"}", string(e.Text()))
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
