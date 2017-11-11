package midi

import "fmt"

// MIDIPortPrefix corresponds to MIDI port prefix meta event.
type MIDIPortPrefixEvent struct {
	deltaTime *DeltaTime
	port      uint8
}

// DeltaTime returns delta time of MIDI port prefix event.
func (e *MIDIPortPrefixEvent) DeltaTime() *DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &DeltaTime{}
	}
	return e.deltaTime
}

// String returns string representation of MIDI port prefix meta event.
func (e *MIDIPortPrefixEvent) String() string {
	return fmt.Sprintf("&MIDIPortPrefixEvent{port: %v}", e.port)
}

// Serialize serializes MIDI port prefix meta event.
func (e *MIDIPortPrefixEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, e.DeltaTime().Quantity().Value()...)
	bs = append(bs, Meta, MIDIPortPrefix)
	bs = append(bs, 0x01, e.port)

	return bs
}

// SetPort sets port.
func (e *MIDIPortPrefixEvent) SetPort(port uint8) error {
	if port > 0x0f {
		return fmt.Errorf("midi: maximum port number is 15 (0x0f)")
	}
	e.port = port

	return nil
}

// Port returns port.
func (e *MIDIPortPrefixEvent) Port() uint8 {
	return e.port
}

// NewMIDIPortPrefixEvent returns MIDIPortPrefixEvent with the given parameter.
func NewMIDIPortPrefixEvent(deltaTime *DeltaTime, port uint8) (*MIDIPortPrefixEvent, error) {
	var err error

	event := &MIDIPortPrefixEvent{}
	event.deltaTime = deltaTime

	err = event.SetPort(port)
	if err != nil {
		return nil, err
	}
	return event, nil
}
