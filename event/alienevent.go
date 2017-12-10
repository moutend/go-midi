package event

import (
	"fmt"

	"github.com/moutend/go-midi/constant"
	"github.com/moutend/go-midi/deltatime"
	"github.com/moutend/go-midi/quantity"
)

// AlienEvent represents unknown meta event.
type AlienEvent struct {
	deltaTime     *deltatime.DeltaTime
	runningStatus bool
	metaEventType uint8
	data          []byte
}

// deltatime.DeltaTime returns delta time.
func (e *AlienEvent) DeltaTime() *deltatime.DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &deltatime.DeltaTime{}
	}
	return e.deltaTime
}

// Serialize serializes alien event.
func (e *AlienEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, constant.Meta, e.metaEventType)

	q := &quantity.Quantity{}
	q.SetUint32(uint32(len(e.Data())))
	bs = append(bs, q.Value()...)
	bs = append(bs, e.Data()...)

	return bs
}

// SetRunningStatus sets running status.
func (e *AlienEvent) SetRunningStatus(status bool) {
	e.runningStatus = status
}

// RunningStatus returns running status.
func (e *AlienEvent) RunningStatus() bool {
	return e.runningStatus
}

// SetMetaEventType sets meta event type.
func (e *AlienEvent) SetMetaEventType(metaEventType uint8) error {
	e.metaEventType = metaEventType

	return nil
}

// MetaEventType returns meta event type.
func (e *AlienEvent) MetaEventType() uint8 {
	return e.metaEventType
}

// SetData sets data.
func (e *AlienEvent) SetData(data []byte) error {
	if len(data) > 0xfffffff {
		return fmt.Errorf("midi: maximum length of data is 256 MB")
	}
	e.data = data

	return nil
}

// Data returns data.
func (e *AlienEvent) Data() []byte {
	if e.data == nil {
		e.data = []byte{}
	}
	return e.data
}

// String returns string representation of alien event.
func (e *AlienEvent) String() string {
	return fmt.Sprintf("&AlienEvent{metaEventType: 0x%x, data: %v bytes}", e.metaEventType, len(e.Data()))
}

// NewAlienEvent returns AlienEvent with the given parameter.
func NewAlienEvent(deltaTime *deltatime.DeltaTime, metaEventType uint8, data []byte) (*AlienEvent, error) {
	event := &AlienEvent{}
	event.deltaTime = deltaTime
	event.metaEventType = metaEventType

	err := event.SetData(data)
	if err != nil {
		return nil, err
	}
	return event, nil
}
