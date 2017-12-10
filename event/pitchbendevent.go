package event

import (
	"fmt"

	"github.com/moutend/go-midi/constant"
	"github.com/moutend/go-midi/deltatime"
)

// PitchBendEvent corresponds to pitch bend event.
type PitchBendEvent struct {
	deltaTime     *deltatime.DeltaTime
	runningStatus bool
	channel       uint8
	pitch         uint16
}

// deltatime.DeltaTime returns delta time of pitch bend event.
func (e *PitchBendEvent) DeltaTime() *deltatime.DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &deltatime.DeltaTime{}
	}
	return e.deltaTime
}

// Serialize serializes pitch bend event.
func (e *PitchBendEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, constant.PitchBend+e.channel)

	msb := byte(e.pitch >> 7)
	lsb := byte(e.pitch & 0x7f)
	bs = append(bs, msb, lsb)

	return bs
}

// SetRunningStatus sets running status.
func (e *PitchBendEvent) SetRunningStatus(status bool) {
	e.runningStatus = status
}

// RunningStatus returns running status.
func (e *PitchBendEvent) RunningStatus() bool {
	return e.runningStatus
}

// SetChannel sets channel.
func (e *PitchBendEvent) SetChannel(channel uint8) error {
	if channel > 0x0f {
		return fmt.Errorf("midi: maximum channel number is 15 (0x0f)")
	}
	e.channel = channel

	return nil
}

// Channel returns channel.
func (e *PitchBendEvent) Channel() uint8 {
	return e.channel
}

// SetPitch sets pitch.
func (e *PitchBendEvent) SetPitch(pitch uint16) error {
	if pitch > 0x3fff {
		return fmt.Errorf("midi: maximum value of pitch is 16384 (0x3fff)")
	}
	e.pitch = pitch

	return nil
}

// Pitch returns pitch.
func (e *PitchBendEvent) Pitch() uint16 {
	return e.pitch
}

// String returns string representation of pitch bend event.
func (e *PitchBendEvent) String() string {
	return fmt.Sprintf("&PitchBendEvent{channel: %v, pitch: %v}", e.channel, e.pitch)
}

// NewPitchBendEvent returns PitchBendEvent with the given parameter.
func NewPitchBendEvent(deltaTime *deltatime.DeltaTime, channel uint8, pitch uint16) (*PitchBendEvent, error) {
	var err error

	event := &PitchBendEvent{}
	event.deltaTime = deltaTime

	err = event.SetChannel(channel)
	if err != nil {
		return nil, err
	}
	err = event.SetPitch(pitch)
	if err != nil {
		return nil, err
	}
	return event, nil
}
