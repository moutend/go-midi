package midi

import "fmt"

// MIDIChannelPrefix corresponds to MIDI channel prefix meta event.
type MIDIChannelPrefixEvent struct {
	deltaTime *DeltaTime
	channel   uint8
}

// DeltaTime returns delta time of MIDI channel prefix event.
func (e *MIDIChannelPrefixEvent) DeltaTime() *DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &DeltaTime{}
	}
	return e.deltaTime
}

// String returns string representation of MIDI channel prefix meta event.
func (e *MIDIChannelPrefixEvent) String() string {
	return fmt.Sprintf("&MIDIChannelPrefixEvent{channel: %v}", e.channel)
}

// Serialize serializes MIDI channel prefix meta event.
func (e *MIDIChannelPrefixEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, e.DeltaTime().Quantity().Value()...)
	bs = append(bs, Meta, MIDIChannelPrefix)
	bs = append(bs, 0x01, e.channel)

	return bs
}

// SetChannel sets channel.
func (e *MIDIChannelPrefixEvent) SetChannel(channel uint8) error {
	if channel > 0x0f {
		return fmt.Errorf("midi: maximum channel number is 15 (0x0f)")
	}
	e.channel = channel

	return nil
}

// Channel returns channel.
func (e *MIDIChannelPrefixEvent) Channel() uint8 {
	return e.channel
}

// NewMIDIChannelPrefixEvent returns MIDIChannelPrefixEvent with the given parameter.
func NewMIDIChannelPrefixEvent(deltaTime *DeltaTime, channel uint8) (*MIDIChannelPrefixEvent, error) {
	var err error

	event := &MIDIChannelPrefixEvent{}
	event.deltaTime = deltaTime

	err = event.SetChannel(channel)
	if err != nil {
		return nil, err
	}
	return event, nil
}
