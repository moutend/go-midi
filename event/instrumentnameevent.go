package event

import (
	"fmt"

	"github.com/moutend/go-midi/constant"
	"github.com/moutend/go-midi/deltatime"
	"github.com/moutend/go-midi/quantity"
)

// InstrumentNameEvent corresponds to instrument name event.
type InstrumentNameEvent struct {
	deltaTime     *deltatime.DeltaTime
	runningStatus bool
	text          []byte
}

// deltatime.DeltaTime returns delta time of instrument name event.
func (e *InstrumentNameEvent) DeltaTime() *deltatime.DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &deltatime.DeltaTime{}
	}
	return e.deltaTime
}

// Serialize serializes instrument name event.
func (e *InstrumentNameEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, constant.Meta, constant.InstrumentName)

	q := &quantity.Quantity{}
	q.SetUint32(uint32(len(e.Text())))
	bs = append(bs, q.Value()...)
	bs = append(bs, e.Text()...)

	return bs
}

// SetRunningStatus sets running status.
func (e *InstrumentNameEvent) SetRunningStatus(status bool) {
	e.runningStatus = status
}

// RunningStatus returns running status.
func (e *InstrumentNameEvent) RunningStatus() bool {
	return e.runningStatus
}

// SetText sets text.
func (e *InstrumentNameEvent) SetText(text []byte) error {
	if len(text) > 0xfffffff {
		return fmt.Errorf("midi: maximum size of text is 256 MB")
	}
	e.text = text

	return nil
}

// Text returns text.
func (e *InstrumentNameEvent) Text() []byte {
	if e.text == nil {
		e.text = []byte{}
	}

	text := make([]byte, len(e.text))
	copy(text, e.text)

	return text
}

// String returns string representation of instrument name event.
func (e *InstrumentNameEvent) String() string {
	return fmt.Sprintf("&InstrumentNameEvent{text: \"%v\"}", string(e.Text()))
}

// NewInstrumentNameEvent returns InstrumentNameEvent with the given parameter.
func NewInstrumentNameEvent(deltaTime *deltatime.DeltaTime, text []byte) (*InstrumentNameEvent, error) {
	var err error

	event := &InstrumentNameEvent{}
	event.deltaTime = deltaTime

	err = event.SetText(text)
	if err != nil {
		return nil, err
	}
	return event, nil
}
