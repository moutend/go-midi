package event

import (
	"github.com/moutend/go-midi/constant"
	"github.com/moutend/go-midi/deltatime"
)

// EndOfTrackEvent corresponds to end of track event.
type EndOfTrackEvent struct {
	deltaTime *deltatime.DeltaTime
}

// deltatime.DeltaTime returns delta time of end of track event.
func (e *EndOfTrackEvent) DeltaTime() *deltatime.DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &deltatime.DeltaTime{}
	}
	return e.deltaTime
}

// Serialize serializes end of track event.
func (e *EndOfTrackEvent) Serialize() []byte {
	return []byte{constant.Meta, constant.EndOfTrack, 0}
}

// RunningStatus is fake method.
// It returns always false because the end of track event cannot omit its event type.
func (e *EndOfTrackEvent) RunningStatus() bool {
	return false
}

// SetRunningStatus is fake method.
// It does nothing because the end of track event cannot omit its event type.
func (e *EndOfTrackEvent) SetRunningStatus(status bool) {
	return
}

// String returns string representation of end of track event.
func (e *EndOfTrackEvent) String() string {
	return "&EndOfTrackEvent{}"
}

// NewEndOfTrackEvent returns EndOfTrackEvent with the given parameter.
func NewEndOfTrackEvent(deltaTime *deltatime.DeltaTime) (*EndOfTrackEvent, error) {
	event := &EndOfTrackEvent{}
	event.deltaTime = deltaTime

	return event, nil
}
