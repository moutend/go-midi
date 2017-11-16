package midi

import "fmt"

// ControllerEvent corresponds to controller event.
type ControllerEvent struct {
	deltaTime *DeltaTime
	channel   uint8
	control   Control
	value     uint8
}

// DeltaTime returns delta time of controller event.
func (e *ControllerEvent) DeltaTime() *DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &DeltaTime{}
	}
	return e.deltaTime
}

// String returns string representation of controller event.
func (e *ControllerEvent) String() string {
	return fmt.Sprintf("&ControllerEvent{channel: %v, control: %v, value: %v}", e.channel, e.control, e.value)
}

// Serialize serializes controller event.
func (e *ControllerEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, e.DeltaTime().Quantity().Value()...)
	bs = append(bs, Controller+e.channel)
	bs = append(bs, byte(e.control), e.value)

	return bs
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
func (e *ControllerEvent) SetControl(control Control) error {
	if control > 0x7f {
		return fmt.Errorf("midi: maximum value of control is 127 (0x7f)")
	}
	e.control = control

	return nil
}

// Control returns control.
func (e *ControllerEvent) Control() Control {
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

// NewControllerEvent returns ControllerEvent with the given parameter.
func NewControllerEvent(deltaTime *DeltaTime, channel uint8, control Control, value uint8) (*ControllerEvent, error) {
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
