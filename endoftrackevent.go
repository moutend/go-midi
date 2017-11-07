package midi

// EndOfTrackEvent corresponds to end of track event (0xff21) in MIDI.
type EndOfTrackEvent struct {
	deltaTime *DeltaTime
}

// DeltaTime returns delta time of this event as DeltaTime.
func (e *EndOfTrackEvent) DeltaTime() *DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &DeltaTime{}
	}

	return e.deltaTime
}

// String returns string representation of this event.
func (e *EndOfTrackEvent) String() string {
	return "&EndOfTrackEvent{}"
}

// Serialize serializes this event according to the SMF specification.
func (e *EndOfTrackEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, e.DeltaTime().Quantity().Value()...)
	bs = append(bs, 0xff, 0x2f, 0x00)

	return bs
}
