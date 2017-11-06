package midi

import "fmt"

type LyricsEvent struct {
	deltaTime *DeltaTime
	text      []byte
}

func (e *LyricsEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *LyricsEvent) String() string {
	return fmt.Sprintf("&LyricsEvent{}")
}
