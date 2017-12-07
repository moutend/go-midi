package midi

import "fmt"

// TextEvent corresponds to text event.
type TextEvent struct {
	deltaTime     *DeltaTime
	runningStatus bool
	text          []byte
}

// DeltaTime returns delta time of text event.
func (e *TextEvent) DeltaTime() *DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &DeltaTime{}
	}
	return e.deltaTime
}

// Serialize serializes text event.
func (e *TextEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, e.DeltaTime().Quantity().Value()...)
	bs = append(bs, Meta, Text)

	q := &Quantity{}
	q.SetUint32(uint32(len(e.Text())))
	bs = append(bs, q.Value()...)
	bs = append(bs, e.Text()...)

	return bs
}

// SetRunningStatus sets running status.
func (e *TextEvent) SetRunningStatus(status bool) {
	e.runningStatus = status
}

// RunningStatus returns running status.
func (e *TextEvent) RunningStatus() bool {
	return e.runningStatus
}

// SetText sets text.
func (e *TextEvent) SetText(text []byte) error {
	if len(text) > 0xfffffff {
		return fmt.Errorf("midi: maximum size of text is 256 MB")
	}
	e.text = text

	return nil
}

// Text returns text.
func (e *TextEvent) Text() []byte {
	if e.text == nil {
		e.text = []byte{}
	}

	text := make([]byte, len(e.text))
	copy(text, e.text)

	return text
}

// String returns string representation of text event.
func (e *TextEvent) String() string {
	return fmt.Sprintf("&TextEvent{text: \"%v\"}", string(e.Text()))
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
