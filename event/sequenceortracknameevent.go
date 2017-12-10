package event

import (
	"fmt"

	"github.com/moutend/go-midi/constant"
	"github.com/moutend/go-midi/deltatime"
	"github.com/moutend/go-midi/quantity"
)

// SequenceOrTrackNameEvent corresponds to sequence or track name event.
type SequenceOrTrackNameEvent struct {
	deltaTime     *deltatime.DeltaTime
	runningStatus bool
	text          []byte
}

// deltatime.DeltaTime returns delta time of sequence or track name event.
func (e *SequenceOrTrackNameEvent) DeltaTime() *deltatime.DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &deltatime.DeltaTime{}
	}
	return e.deltaTime
}

// Serialize serializes sequence or track name event.
func (e *SequenceOrTrackNameEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, constant.Meta, constant.SequenceOrTrackName)

	q := &quantity.Quantity{}
	q.SetUint32(uint32(len(e.Text())))
	bs = append(bs, q.Value()...)
	bs = append(bs, e.Text()...)

	return bs
}

// SetRunningStatus sets running status.
func (e *SequenceOrTrackNameEvent) SetRunningStatus(status bool) {
	e.runningStatus = status
}

// RunningStatus returns running status.
func (e *SequenceOrTrackNameEvent) RunningStatus() bool {
	return e.runningStatus
}

// SetText sets text.
func (e *SequenceOrTrackNameEvent) SetText(text []byte) error {
	if len(text) > 0xfffffff {
		return fmt.Errorf("midi: maximum size of text is 256 MB")
	}
	e.text = text

	return nil
}

// Text returns text.
func (e *SequenceOrTrackNameEvent) Text() []byte {
	if e.text == nil {
		e.text = []byte{}
	}

	text := make([]byte, len(e.text))
	copy(text, e.text)

	return text
}

// String returns string representation of sequence or track name event.
func (e *SequenceOrTrackNameEvent) String() string {
	return fmt.Sprintf("&SequenceOrTrackNameEvent{text: \"%v\"}", string(e.Text()))
}

// NewSequenceOrTrackNameEvent returns SequenceOrTrackNameEvent with the given parameter.
func NewSequenceOrTrackNameEvent(deltaTime *deltatime.DeltaTime, text []byte) (*SequenceOrTrackNameEvent, error) {
	var err error

	event := &SequenceOrTrackNameEvent{}
	event.deltaTime = deltaTime

	err = event.SetText(text)
	if err != nil {
		return nil, err
	}
	return event, nil
}
