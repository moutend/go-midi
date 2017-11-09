package midi

import "fmt"

// ChannelAfterTouchEvent corresponds to channel after touch event (0xd0) in MIDI.
type ChannelAfterTouchEvent struct {
	deltaTime *DeltaTime
	channel   uint8
	velocity  uint8
}

// DeltaTime returns delta time of this event.
func (e *ChannelAfterTouchEvent) DeltaTime() *DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &DeltaTime{}
	}
	return e.deltaTime
}

// String returns string representation of this event.
func (e *ChannelAfterTouchEvent) String() string {
	return fmt.Sprintf("&ChannelAfterTouchEvent{channel: %v, velocity: %v}", e.channel, e.velocity)
}

// Serialize serializes this event.
func (e *ChannelAfterTouchEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, e.DeltaTime().Quantity().Value()...)
	bs = append(bs, ChannelAfterTouch+e.channel)
	bs = append(bs, e.velocity)

	return bs
}

// SetChannel sets channel of this event.
func (e *ChannelAfterTouchEvent) SetChannel(channel uint8) error {
	if channel > 0x0f {
		return fmt.Errorf("midi: maximum channel number is 15 (0x0f)")
	}
	e.channel = channel

	return nil
}

// SetVelocity sets velocity of this event.
func (e *ChannelAfterTouchEvent) SetVelocity(velocity uint8) error {
	if velocity > 0x7f {
		return fmt.Errorf("midi: maximum value of velocity is 127 (0x7f)")
	}
	e.velocity = velocity

	return nil
}

// NewChannelAfterTouchEvent returns ChannelAfterTouchEvent with the given parameter.
func NewChannelAfterTouchEvent(deltaTime *DeltaTime, channel, velocity uint8) (*ChannelAfterTouchEvent, error) {
	var err error

	event := &ChannelAfterTouchEvent{}
	event.deltaTime = deltaTime

	err = event.SetChannel(channel)
	if err != nil {
		return nil, err
	}
	err = event.SetVelocity(velocity)
	if err != nil {
		return nil, err
	}
	return event, nil
}
