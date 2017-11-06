package midi

import "fmt"

type SequenceOrTrackNameEvent struct {
	deltaTime *DeltaTime
	text      []byte
}

func (e *SequenceOrTrackNameEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *SequenceOrTrackNameEvent) String() string {
	return fmt.Sprintf("&SequenceOrTrackNameEvent{text: \"%v\"}", string(e.text))
}
