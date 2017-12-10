package event

import (
	"fmt"

	"github.com/moutend/go-midi/constant"
	"github.com/moutend/go-midi/deltatime"
)

// MIDIChannelPrefix corresponds to MIDI channel prefix meta event.
type MIDIChannelPrefixEvent struct {
	deltaTime     *deltatime.DeltaTime
	runningStatus bool
	channel       uint8
}

// deltatime.DeltaTime returns delta time of MIDI channel prefix event.
func (e *MIDIChannelPrefixEvent) DeltaTime() *deltatime.DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &deltatime.DeltaTime{}
	}
	return e.deltaTime
}

// Serialize serializes MIDI channel prefix meta event.
func (e *MIDIChannelPrefixEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, constant.Meta, constant.MIDIChannelPrefix)
	bs = append(bs, 0x01, e.channel)

	return bs
}

// SetRunningStatus sets running status.
func (e *MIDIChannelPrefixEvent) SetRunningStatus(status bool) {
	e.runningStatus = status
}

// RunningStatus returns running status.
func (e *MIDIChannelPrefixEvent) RunningStatus() bool {
	return e.runningStatus
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

// String returns string representation of MIDI channel prefix meta event.
func (e *MIDIChannelPrefixEvent) String() string {
	return fmt.Sprintf("&MIDIChannelPrefixEvent{channel: %v}", e.channel)
}

// NewMIDIChannelPrefixEvent returns MIDIChannelPrefixEvent with the given parameter.
func NewMIDIChannelPrefixEvent(deltaTime *deltatime.DeltaTime, channel uint8) (*MIDIChannelPrefixEvent, error) {
	var err error

	event := &MIDIChannelPrefixEvent{}
	event.deltaTime = deltaTime

	err = event.SetChannel(channel)
	if err != nil {
		return nil, err
	}
	return event, nil
}
