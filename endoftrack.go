package midi

type EndOfTrackEvent struct {
	deltaTime *DeltaTime
}

func (e *EndOfTrackEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *EndOfTrackEvent) String() string {
	return "&EndOfTrackEvent{}"
}
