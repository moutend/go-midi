package midi

import "fmt"

// ContinuousControllerEvent corresponds to continuous controller event.
type ContinuousControllerEvent struct {
	deltaTime *DeltaTime
	control   uint8
	value     uint8
}

// DeltaTime returns delta time of continuous controller event.
func (e *ContinuousControllerEvent) DeltaTime() *DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &DeltaTime{}
	}
	return e.deltaTime
}

// String returns string representation of continuous controller event.
func (e *ContinuousControllerEvent) String() string {
	return fmt.Sprintf("&ContinuousControllerEvent{control: %v, value: %v}", e.control, e.value)
}

// Serialize serializes continuous controller event.
func (e *ContinuousControllerEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, e.DeltaTime().Quantity().Value()...)
	bs = append(bs, e.control, e.value)

	return bs
}

// SetControl sets control.
func (e *ContinuousControllerEvent) SetControl(control uint8) error {
	if control > 0x7f {
		return fmt.Errorf("midi: maximum value of control is 127 (0x7f)")
	}
	e.control = control

	return nil
}

// Control returns control.
func (e *ContinuousControllerEvent) Control() uint8 {
	return e.control
}

// SetValue sets value.
func (e *ContinuousControllerEvent) SetValue(value uint8) error {
	if value > 0x7f {
		return fmt.Errorf("midi: maximum value of value is 127 (0x7f)")
	}
	e.value = value

	return nil
}

// Value returns value.
func (e *ContinuousControllerEvent) Value() uint8 {
	return e.value
}

// NewContinuousControllerEvent returns ContinuousControllerEvent with the given parameter.
func NewContinuousControllerEvent(deltaTime *DeltaTime, control, value uint8) (*ContinuousControllerEvent, error) {
	var err error

	event := &ContinuousControllerEvent{}
	event.deltaTime = deltaTime

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
