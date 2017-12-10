package event

import (
	"fmt"

	"github.com/moutend/go-midi/constant"
	"github.com/moutend/go-midi/deltatime"
	"github.com/moutend/go-midi/quantity"
)

// SystemExclusiveEvent corresponds to system exclusive meta event.
type SystemExclusiveEvent struct {
	deltaTime     *deltatime.DeltaTime
	runningStatus bool
	data          []byte
}

// deltatime.DeltaTime returns delta time of system exclusive event.
func (e *SystemExclusiveEvent) DeltaTime() *deltatime.DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &deltatime.DeltaTime{}
	}
	return e.deltaTime
}

// Serialize serializes system exclusive event.
func (e *SystemExclusiveEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, constant.SystemExclusive)

	q := &quantity.Quantity{}
	q.SetUint32(uint32(len(e.Data())))
	bs = append(bs, q.Serialize()...)
	bs = append(bs, e.Data()...)

	return bs
}

// SetRunningStatus sets running status.
func (e *SystemExclusiveEvent) SetRunningStatus(status bool) {
	e.runningStatus = status
}

// RunningStatus returns running status.
func (e *SystemExclusiveEvent) RunningStatus() bool {
	return e.runningStatus
}

// SetData sets data.
func (e *SystemExclusiveEvent) SetData(data []byte) error {
	if len(data) > 0xfffffff {
		return fmt.Errorf("midi: maximum size of data is 256 MB")
	}
	e.data = data

	return nil
}

// Data returns data.
func (e *SystemExclusiveEvent) Data() []byte {
	if e.data == nil {
		e.data = []byte{}
	}

	return e.data
}

// String returns string representation of system exclusive event.
func (e *SystemExclusiveEvent) String() string {
	return fmt.Sprintf("&SystemExclusiveEvent{data: %v bytes}", len(e.Data()))
}

// NewSystemExclusiveEvent returns SystemExclusiveEvent with the given parameter.
func NewSystemExclusiveEvent(deltaTime *deltatime.DeltaTime, data []byte) (*SystemExclusiveEvent, error) {
	var err error

	event := &SystemExclusiveEvent{}
	event.deltaTime = deltaTime

	err = event.SetData(data)
	if err != nil {
		return nil, err
	}
	return event, nil
}
