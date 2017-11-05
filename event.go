package midi

import (
	"bytes"
	"encoding/binary"
	"io"
)

type NoteOffEvent struct {
	deltaTime *DeltaTime
	channel   uint8
	note      uint8
	velocity  uint8
}

func (e *NoteOffEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

type NoteOnEvent struct {
	deltaTime *DeltaTime
	channel   uint8
	note      uint8
	velocity  uint8
}

func (e *NoteOnEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

type NoteAfterTouchEvent struct {
	deltaTime *DeltaTime
	channel   uint8
	note      uint8
	velocity  uint8
}

func (e *NoteAfterTouchEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

type ControllerEvent struct {
	deltaTime *DeltaTime
	channel   uint8
	control   uint8
	value     uint8
}

func (e *ControllerEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

type ProgramChangeEvent struct {
	deltaTime *DeltaTime
	channel   uint8
	program   uint8
}

func (e *ProgramChangeEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

type ChannelAfterTouchEvent struct {
	deltaTime *DeltaTime
	channel   uint8
	velocity  uint8
}

func (e *ChannelAfterTouchEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

type PitchBendEvent struct {
	deltaTime *DeltaTime
	channel   uint8
	note      uint8
	velocity  uint16
}

func (e *PitchBendEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

type Event interface {
	DeltaTime() *DeltaTime
}

func parseEvent(stream []byte, chunkSize uint32) (*Event, error) {
	deltaTime, err := parseDeltaTime(stream)
	if err != nil {
		return nil, err
	}

	var event Event
	var eventType EventType
	var channel uint8

	parameter := make([]byte, 2)
	data := bytes.NewReader(stream)
	sizeOfDeltaTime := int64(len(deltaTime.value))

	binary.Read(io.NewSectionReader(data, sizeOfDeltaTime, 1), binary.BigEndian, &eventType)
	binary.Read(io.NewSectionReader(data, sizeOfDeltaTime+1, 2), binary.BigEndian, &parameter[0])

	if eventType == Meta {
		return &event, nil
	}

	channel = uint8(eventType) & 0x0f
	eventType = eventType & 0xf0

	switch eventType {
	case NoteOff:
		event = &NoteOffEvent{
			deltaTime: deltaTime,
			channel:   channel,
			note:      uint8(parameter[0]),
			velocity:  uint8(parameter[1]),
		}
	case NoteOn:
		event = &NoteOnEvent{
			deltaTime: deltaTime,
			channel:   channel,
			note:      uint8(parameter[0]),
			velocity:  uint8(parameter[1]),
		}
	case NoteAfterTouch:
		event = &NoteAfterTouchEvent{
			deltaTime: deltaTime,
			channel:   channel,
			note:      uint8(parameter[0]),
			velocity:  uint8(parameter[1]),
		}
	case Controller:
		event = &ControllerEvent{
			deltaTime: deltaTime,
			channel:   channel,
			control:   uint8(parameter[0]),
			value:     uint8(parameter[1]),
		}
	case ProgramChange:
		event = &ProgramChangeEvent{
			deltaTime: deltaTime,
			channel:   channel,
			program:   uint8(parameter[0]),
		}
	case ChannelAfterTouch:
		event = &ChannelAfterTouchEvent{
			deltaTime: deltaTime,
			channel:   channel,
			velocity:  uint8(parameter[0]),
		}
	case PitchBend:
		event = &NoteOffEvent{
			deltaTime: deltaTime,
			channel:   channel,
			note:      uint8(parameter[0]),
			velocity:  uint8(parameter[1]),
		}
	}

	return &event, nil
}
