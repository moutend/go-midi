package midi

import "fmt"

// ProgramChangeEvent corresponds to program change event.
type ProgramChangeEvent struct {
	deltaTime     *DeltaTime
	runningStatus bool
	channel       uint8
	program       GM
}

// DeltaTime returns delta time of program change event.
func (e *ProgramChangeEvent) DeltaTime() *DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &DeltaTime{}
	}

	return e.deltaTime
}

// Serialize serializes program change event.
func (e *ProgramChangeEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, e.DeltaTime().Quantity().Value()...)
	bs = append(bs, ProgramChange+e.channel)
	bs = append(bs, byte(e.program))

	return bs
}

// SetRunningStatus sets running status.
func (e *ProgramChangeEvent) SetRunningStatus(status bool) {
	e.runningStatus = status
}

// RunningStatus returns running status.
func (e *ProgramChangeEvent) RunningStatus() bool {
	return e.runningStatus
}

// SetChannel sets channel.
func (e *ProgramChangeEvent) SetChannel(channel uint8) error {
	if channel > 0x0f {
		return fmt.Errorf("midi: maximum channel number is 15 (0x0f)")
	}
	e.channel = channel

	return nil
}

// Channel returns channel.
func (e *ProgramChangeEvent) Channel() uint8 {
	return e.channel
}

// SetProgram sets program.
func (e *ProgramChangeEvent) SetProgram(program GM) error {
	if program > 0x7f {
		return fmt.Errorf("midi: maximum value of program is 127 (0x7f)")
	}
	e.program = program

	return nil
}

// Program returns program.
func (e *ProgramChangeEvent) Program() GM {
	return e.program
}

// String returns string representation of program change event.
func (e *ProgramChangeEvent) String() string {
	return fmt.Sprintf("&ProgramChangeEvent{channel: %v, program: %v}", e.channel, e.program)
}

// NewProgramChangeEvent returns ProgramChangeEvent with the given parameter.
func NewProgramChangeEvent(deltaTime *DeltaTime, channel uint8, program GM) (*ProgramChangeEvent, error) {
	var err error

	event := &ProgramChangeEvent{}
	event.deltaTime = deltaTime

	err = event.SetChannel(channel)
	if err != nil {
		return nil, err
	}
	err = event.SetProgram(program)
	if err != nil {
		return nil, err
	}
	return event, nil
}
