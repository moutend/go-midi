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

type AlienEvent struct {
	deltaTime     *DeltaTime
	metaEventType MetaEventType
	data          []byte
}

func (e *AlienEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

type TextEvent struct {
	deltaTime *DeltaTime
	text      []byte
}

func (e *TextEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

type CopyrightNoticeEvent struct {
	deltaTime *DeltaTime
	text      []byte
}

func (e *CopyrightNoticeEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

type Event interface {
	DeltaTime() *DeltaTime
}

func parseEvent(stream []byte, chunkSize uint32) (Event, error) {
	deltaTime, err := parseDeltaTime(stream)
	if err != nil {
		return nil, err
	}

	var eventType EventType

	data := bytes.NewReader(stream)
	sizeOfDeltaTime := int64(len(deltaTime.value))
	binary.Read(io.NewSectionReader(data, sizeOfDeltaTime, 1), binary.BigEndian, &eventType)

	if eventType == Meta {
		return parseMetaEvent(stream[sizeOfDeltaTime:], deltaTime)
	}

	return parseNormalEvent(stream, deltaTime, eventType)
}

func parseMetaEvent(stream []byte, deltaTime *DeltaTime) (Event, error) {
	var event Event

	metaEventType := MetaEventType(stream[1])
	sizeOfMetaEventData := int64(stream[2])
	metaEventData := stream[2 : sizeOfMetaEventData+2]

	switch metaEventType {
	case Text:
		event = &TextEvent{
			deltaTime: deltaTime,
			text:      metaEventData,
		}
	case CopyrightNotice:
		event = &CopyrightNoticeEvent{
			deltaTime: deltaTime,
			text:      metaEventData,
		}
	default:
		event = &AlienEvent{
			deltaTime:     deltaTime,
			metaEventType: metaEventType,
			data:          metaEventData,
		}
	}

	return event, nil
}

func parseNormalEvent(stream []byte, deltaTime *DeltaTime, eventType EventType) (Event, error) {
	var event Event

	parameter := stream[1:2]
	channel := uint8(eventType) & 0x0f
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
	default:
	}

	return event, nil
}
