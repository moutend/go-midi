package event

import (
	"fmt"

	"github.com/moutend/go-midi/constant"
	"github.com/moutend/go-midi/deltatime"
	"github.com/moutend/go-midi/quantity"
)

// MarkerEvent corresponds to marker event.
type MarkerEvent struct {
	deltaTime     *deltatime.DeltaTime
	runningStatus bool
	text          []byte
}

// deltatime.DeltaTime returns delta time of marker event.
func (e *MarkerEvent) DeltaTime() *deltatime.DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &deltatime.DeltaTime{}
	}
	return e.deltaTime
}

// Serialize serializes marker event.
func (e *MarkerEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, constant.Meta, constant.Marker)

	q := &quantity.Quantity{}
	q.SetUint32(uint32(len(e.Text())))
	bs = append(bs, q.Value()...)
	bs = append(bs, e.Text()...)

	return bs
}

// SetRunningStatus sets running status.
func (e *MarkerEvent) SetRunningStatus(status bool) {
	e.runningStatus = status
}

// RunningStatus returns running status.
func (e *MarkerEvent) RunningStatus() bool {
	return e.runningStatus
}

// SetText sets text.
func (e *MarkerEvent) SetText(text []byte) error {
	if len(text) > 0xfffffff {
		return fmt.Errorf("midi: maximum size of text is 256 MB")
	}
	e.text = text

	return nil
}

// Text returns text.
func (e *MarkerEvent) Text() []byte {
	if e.text == nil {
		e.text = []byte{}
	}

	text := make([]byte, len(e.text))
	copy(text, e.text)

	return text
}

// String returns string representation of marker event.
func (e *MarkerEvent) String() string {
	return fmt.Sprintf("&MarkerEvent{text: \"%v\"}", string(e.Text()))
}

// NewMarkerEvent returns MarkerEvent with the given parameter.
func NewMarkerEvent(deltaTime *deltatime.DeltaTime, text []byte) (*MarkerEvent, error) {
	var err error

	event := &MarkerEvent{}
	event.deltaTime = deltaTime

	err = event.SetText(text)
	if err != nil {
		return nil, err
	}
	return event, nil
}
