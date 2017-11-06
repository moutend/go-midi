package midi

import "fmt"

type ChannelAfterTouchEvent struct {
	deltaTime *DeltaTime
	channel   uint8
	velocity  uint8
}

func (e *ChannelAfterTouchEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *ChannelAfterTouchEvent) String() string {
	return fmt.Sprintf("&ChannelAfterTouchEvent{}")
}
