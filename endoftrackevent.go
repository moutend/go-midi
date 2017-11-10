package midi

// EndOfTrackEvent corresponds to end of track event.
type EndOfTrackEvent struct {
	deltaTime *DeltaTime
}

// DeltaTime returns delta time of end of track event.
func (e *EndOfTrackEvent) DeltaTime() *DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &DeltaTime{}
	}
	return e.deltaTime
}

// String returns string representation of end of track event.
func (e *EndOfTrackEvent) String() string {
	return "&EndOfTrackEvent{}"
}

// Serialize serializes end of track event.
func (e *EndOfTrackEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, e.DeltaTime().Quantity().Value()...)
	bs = append(bs, 0xff, 0x2f, 0x00)

	return bs
}

// NewEndOfTrackEvent returns EndOfTrackEvent with the given parameter.
func NewEndOfTrackEvent(deltaTime *DeltaTime) (*EndOfTrackEvent, error) {
	event := &EndOfTrackEvent{}
	event.deltaTime = deltaTime

	return event, nil
}
