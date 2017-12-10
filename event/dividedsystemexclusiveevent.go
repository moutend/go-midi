package event

import (
	"fmt"

	"github.com/moutend/go-midi/constant"
	"github.com/moutend/go-midi/deltatime"
	"github.com/moutend/go-midi/quantity"
)

// DividedSystemExclusiveEvent corresponds to system exclusive meta event.
type DividedSystemExclusiveEvent struct {
	deltaTime     *deltatime.DeltaTime
	runningStatus bool
	data          []byte
}

// deltatime.DeltaTime returns delta time of system exclusive event.
func (e *DividedSystemExclusiveEvent) DeltaTime() *deltatime.DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &deltatime.DeltaTime{}
	}
	return e.deltaTime
}

// Serialize serializes system exclusive event.
func (e *DividedSystemExclusiveEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, constant.DividedSystemExclusive)

	q := &quantity.Quantity{}
	q.SetUint32(uint32(len(e.Data())))
	bs = append(bs, q.Serialize()...)
	bs = append(bs, e.Data()...)

	return bs
}

// SetRunningStatus sets running status.
func (e *DividedSystemExclusiveEvent) SetRunningStatus(status bool) {
	e.runningStatus = status
}

// RunningStatus returns running status.
func (e *DividedSystemExclusiveEvent) RunningStatus() bool {
	return e.runningStatus
}

// SetData sets data.
func (e *DividedSystemExclusiveEvent) SetData(data []byte) error {
	if len(data) > 0xfffffff {
		return fmt.Errorf("midi: maximum size of data is 256 MB")
	}
	e.data = data

	return nil
}

// Data returns data.
func (e *DividedSystemExclusiveEvent) Data() []byte {
	if e.data == nil {
		e.data = []byte{}
	}

	return e.data
}

// String returns string representation of system exclusive event.
func (e *DividedSystemExclusiveEvent) String() string {
	return fmt.Sprintf("&DividedSystemExclusiveEvent{data: %v bytes}", len(e.Data()))
}

// NewDividedSystemExclusiveEvent returns DividedSystemExclusiveEvent with the given parameter.
func NewDividedSystemExclusiveEvent(deltaTime *deltatime.DeltaTime, data []byte) (*DividedSystemExclusiveEvent, error) {
	var err error

	event := &DividedSystemExclusiveEvent{}
	event.deltaTime = deltaTime

	err = event.SetData(data)
	if err != nil {
		return nil, err
	}
	return event, nil
}
