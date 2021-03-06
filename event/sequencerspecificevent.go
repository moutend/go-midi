package event

import (
	"fmt"

	"github.com/moutend/go-midi/constant"
	"github.com/moutend/go-midi/deltatime"
	"github.com/moutend/go-midi/quantity"
)

// SequencerSpecificEvent corresponds to sequencer specific event.
type SequencerSpecificEvent struct {
	deltaTime     *deltatime.DeltaTime
	runningStatus bool
	data          []byte
}

// deltatime.DeltaTime returns delta time.
func (e *SequencerSpecificEvent) DeltaTime() *deltatime.DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &deltatime.DeltaTime{}
	}
	return e.deltaTime
}

// Serialize serializes sequencer specific event.
func (e *SequencerSpecificEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, constant.Meta, constant.SequencerSpecific)

	q := &quantity.Quantity{}
	q.SetUint32(uint32(len(e.Data())))
	bs = append(bs, q.Value()...)
	bs = append(bs, e.Data()...)

	return bs
}

// SetRunningStatus sets running status.
func (e *SequencerSpecificEvent) SetRunningStatus(status bool) {
	e.runningStatus = status
}

// RunningStatus returns running status.
func (e *SequencerSpecificEvent) RunningStatus() bool {
	return e.runningStatus
}

// SetData sets data.
func (e *SequencerSpecificEvent) SetData(data []byte) error {
	if len(data) > 0xfffffff {
		return fmt.Errorf("midi: maximum length of data is 256 MB")
	}
	e.data = data

	return nil
}

// Data returns data.
func (e *SequencerSpecificEvent) Data() []byte {
	if e.data == nil {
		e.data = []byte{}
	}
	return e.data
}

// String returns string representation of sequencer specific event.
func (e *SequencerSpecificEvent) String() string {
	return fmt.Sprintf("&SequencerSpecificEvent{data: %v bytes}", len(e.Data()))
}

// NewSequencerSpecificEvent returns SequencerSpecificEvent with the given parameter.
func NewSequencerSpecificEvent(deltaTime *deltatime.DeltaTime, data []byte) (*SequencerSpecificEvent, error) {
	var err error

	event := &SequencerSpecificEvent{}
	event.deltaTime = deltaTime

	err = event.SetData(data)
	if err != nil {
		return nil, err
	}
	return event, nil
}
