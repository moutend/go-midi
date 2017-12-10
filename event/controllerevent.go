package event

import (
	"fmt"

	"github.com/moutend/go-midi/constant"
	"github.com/moutend/go-midi/deltatime"
)

// ControllerEvent corresponds to controller event.
type ControllerEvent struct {
	deltaTime     *deltatime.DeltaTime
	runningStatus bool
	channel       uint8
	control       constant.Control
	value         uint8
}

// deltatime.DeltaTime returns delta time of controller event.
func (e *ControllerEvent) DeltaTime() *deltatime.DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &deltatime.DeltaTime{}
	}
	return e.deltaTime
}

// Serialize serializes controller event.
func (e *ControllerEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, constant.Controller+e.channel)
	bs = append(bs, byte(e.control), e.value)

	return bs
}

// SetRunningStatus sets running status.
func (e *ControllerEvent) SetRunningStatus(status bool) {
	e.runningStatus = status
}

// RunningStatus returns running status.
func (e *ControllerEvent) RunningStatus() bool {
	return e.runningStatus
}

// SetChannel sets channel.
func (e *ControllerEvent) SetChannel(channel uint8) error {
	if channel > 0x0f {
		return fmt.Errorf("midi: maximum channel number is 15 (0x0f)")
	}
	e.channel = channel

	return nil
}

// Channel returns channel.
func (e *ControllerEvent) Channel() uint8 {
	return e.channel
}

// SetControl sets control.
func (e *ControllerEvent) SetControl(control constant.Control) error {
	if control > 0x7f {
		return fmt.Errorf("midi: maximum value of control is 127 (0x7f)")
	}
	e.control = control

	return nil
}

// Control returns control.
func (e *ControllerEvent) Control() constant.Control {
	return e.control
}

// SetValue sets value.
func (e *ControllerEvent) SetValue(value uint8) error {
	if value > 0x7f {
		return fmt.Errorf("midi: maximum value of value is 127 (0x7f)")
	}
	e.value = value

	return nil
}

// Value returns value.
func (e *ControllerEvent) Value() uint8 {
	return e.value
}

// String returns string representation of controller event.
func (e *ControllerEvent) String() string {
	return fmt.Sprintf("&ControllerEvent{channel: %v, control: %v, value: %v}", e.channel, e.control, e.value)
}

// NewControllerEvent returns ControllerEvent with the given parameter.
func NewControllerEvent(deltaTime *deltatime.DeltaTime, channel uint8, control constant.Control, value uint8) (*ControllerEvent, error) {
	var err error

	event := &ControllerEvent{}
	event.deltaTime = deltaTime

	err = event.SetChannel(channel)
	if err != nil {
		return nil, err
	}
	err = event.SetControl(control)
	if err != nil {
		return nil, err
	}
	err = event.SetValue(value)
	if err != nil {
		return nil, err
	}
	return event, nil
}
