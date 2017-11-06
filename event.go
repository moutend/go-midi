package midi

import (
	"bytes"
	"encoding/binary"
	"fmt"
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

func (e *NoteOffEvent) String() string {
	return fmt.Sprintf("&NoteOffEvent{}")
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

func (e *TextEvent) Text() string {
	return string(e.text)
}

func (e *TextEvent) String() string {
	return fmt.Sprintf("&TextEvent{text: \"%v\"}", string(e.text))
}

type SequenceOrTrackNameEvent struct {
	deltaTime *DeltaTime
	text      []byte
}

func (e *SequenceOrTrackNameEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *SequenceOrTrackNameEvent) String() string {
	return fmt.Sprintf("&SequenceOrTrackNameEvent{text: \"%v\"}", string(e.text))
}

type InstrumentNameEvent struct {
	deltaTime *DeltaTime
	text      []byte
}

func (e *InstrumentNameEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

type LyricsEvent struct {
	deltaTime *DeltaTime
	text      []byte
}

func (e *LyricsEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

type MarkerEvent struct {
	deltaTime *DeltaTime
	text      []byte
}

func (e *MarkerEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *MarkerEvent) String() string {
	return fmt.Sprintf("&MarkerEvent{text: \"%v\"}", string(e.text))
}

type CuePointEvent struct {
	deltaTime *DeltaTime
	text      []byte
}

func (e *CuePointEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

type SetTempoEvent struct {
	deltaTime *DeltaTime
	tempo     []byte
}

func (e *SetTempoEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *SetTempoEvent) String() string {
	return fmt.Sprintf("&SetTempoEvent{}")
}

type SMPTEOffsetEvent struct {
	deltaTime *DeltaTime
	tempo     []byte
}

func (e *SMPTEOffsetEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *SMPTEOffsetEvent) String() string {
	return fmt.Sprintf("&SMPTEOffsetEvent{}")
}

type TimeSignatureEvent struct {
	deltaTime *DeltaTime
	tempo     []byte
}

func (e *TimeSignatureEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *TimeSignatureEvent) String() string {
	return fmt.Sprintf("&TimeSignatureEvent{}")
}

type KeySignatureEvent struct {
	deltaTime *DeltaTime
	tempo     []byte
}

func (e *KeySignatureEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *KeySignatureEvent) String() string {
	return fmt.Sprintf("&KeySignatureEvent{}")
}

type EndOfTrackEvent struct {
	deltaTime *DeltaTime
}

func (e *EndOfTrackEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *EndOfTrackEvent) String() string {
	return "&EndOfTrackEvent{}"
}

type Event interface {
	DeltaTime() *DeltaTime
}

// parseEvent reads stream and returns event and its size.
func parseEvent(stream []byte) (Event, int, error) {
	deltaTime, err := parseDeltaTime(stream)
	if err != nil {
		return nil, 0, err
	}

	var eventType EventType

	data := bytes.NewReader(stream)
	sizeOfDeltaTime := int64(len(deltaTime.value))
	binary.Read(io.NewSectionReader(data, sizeOfDeltaTime, 1), binary.BigEndian, &eventType)

	if eventType == Meta {
		return parseMetaEvent(stream[sizeOfDeltaTime:], deltaTime)
	}

	return parseMIDIControlEvent(stream, deltaTime, eventType)
}

func parseMetaEvent(stream []byte, deltaTime *DeltaTime) (Event, int, error) {
	var event Event

	metaEventType := MetaEventType(stream[1])
	sizeOfMetaEventData := int64(stream[2])
	metaEventData := stream[3 : sizeOfMetaEventData+3]

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
	case SequenceOrTrackName:
		event = &SequenceOrTrackNameEvent{
			deltaTime: deltaTime,
			text:      metaEventData,
		}
	case InstrumentName:
		event = &InstrumentNameEvent{
			deltaTime: deltaTime,
			text:      metaEventData,
		}
	case Lyrics:
		event = &LyricsEvent{
			deltaTime: deltaTime,
			text:      metaEventData,
		}
	case Marker:
		event = &MarkerEvent{
			deltaTime: deltaTime,
			text:      metaEventData,
		}
	case CuePoint:
		event = &CuePointEvent{
			deltaTime: deltaTime,
			text:      metaEventData,
		}
	case SetTempo:
		event = &SetTempoEvent{
			deltaTime: deltaTime,
			tempo:     metaEventData,
		}
	case SMPTEOffset:
		event = &SMPTEOffsetEvent{
			deltaTime: deltaTime,
		}
	case TimeSignature:
		event = &TimeSignatureEvent{
			deltaTime: deltaTime,
		}
	case KeySignature:
		event = &KeySignatureEvent{
			deltaTime: deltaTime,
		}
	case EndOfTrack:
		event = &EndOfTrackEvent{
			deltaTime: deltaTime,
		}
	default:
		event = &AlienEvent{
			deltaTime:     deltaTime,
			metaEventType: metaEventType,
			data:          metaEventData,
		}
	}

	sizeOfEvent := len(deltaTime.value) + 3 + int(sizeOfMetaEventData)

	return event, sizeOfEvent, nil
}

func parseMIDIControlEvent(stream []byte, deltaTime *DeltaTime, eventType EventType) (Event, int, error) {
	var event Event

	parameter := stream[1:2]
	channel := uint8(eventType) & 0x0f
	eventType = eventType & 0xf0
	sizeOfMIDIControlEvent := 3

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
		sizeOfMIDIControlEvent = 2
	case ChannelAfterTouch:
		event = &ChannelAfterTouchEvent{
			deltaTime: deltaTime,
			channel:   channel,
			velocity:  uint8(parameter[0]),
		}
		sizeOfMIDIControlEvent = 2
	case PitchBend:
		event = &NoteOffEvent{
			deltaTime: deltaTime,
			channel:   channel,
			note:      uint8(parameter[0]),
			velocity:  uint8(parameter[1]),
		}
	}

	sizeOfEvent := len(deltaTime.value) + sizeOfMIDIControlEvent

	return event, sizeOfEvent, nil
}
