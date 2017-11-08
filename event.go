package midi

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type Event interface {
	DeltaTime() *DeltaTime
}

// parseEvent reads stream and returns event and its size.
func parseEvent(stream []byte) (Event, int, error) {
	deltaTime, err := parseDeltaTime(stream)
	if err != nil {
		return nil, 0, err
	}
	const SystemExclusive byte = 0xf0
	const Meta = 0xff
	var eventType byte

	data := bytes.NewReader(stream)
	sizeOfDeltaTime := int64(len(deltaTime.Quantity().Value()))
	binary.Read(io.NewSectionReader(data, sizeOfDeltaTime, 1), binary.BigEndian, &eventType)

	if eventType == Meta {
		return parseMetaEvent(stream[sizeOfDeltaTime:], deltaTime)
	}
	if eventType == SystemExclusive {
		return parseSystemExclusiveEvent(stream[sizeOfDeltaTime:], deltaTime)
	}

	return parseMIDIControlEvent(stream, deltaTime, eventType)
}

func parseSystemExclusiveEvent(stream []byte, deltaTime *DeltaTime) (event Event, sizeOfEvent int, err error) {
	q, err := parseQuantity(stream[1:])
	if err != nil {
		return nil, 0, err
	}

	offset := len(deltaTime.Quantity().Value()) + 1 + len(q.value)
	sizeOfSystemExclusiveEventData := int(q.Uint32())
	sizeOfEvent = offset + sizeOfSystemExclusiveEventData

	event = &SystemExclusiveEvent{
		deltaTime: deltaTime,
		data:      stream[offset : offset+sizeOfSystemExclusiveEventData],
	}

	return event, sizeOfEvent, nil
}

func parseMetaEvent(stream []byte, deltaTime *DeltaTime) (event Event, sizeOfEvent int, err error) {
	metaEventType := stream[1]
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
	case MIDIPortPrefix:
		event = &MIDIPortPrefixEvent{
			deltaTime: deltaTime,
			port:      uint8(metaEventData[0]),
		}
	case MIDIChannelPrefix:
		event = &MIDIChannelPrefixEvent{
			deltaTime: deltaTime,
			channel:   uint8(metaEventData[0]),
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

	sizeOfEvent = len(deltaTime.Quantity().Value()) + 3 + int(sizeOfMetaEventData)

	return event, sizeOfEvent, nil
}

func parseMIDIControlEvent(stream []byte, deltaTime *DeltaTime, eventType byte) (event Event, sizeOfEvent int, err error) {
	parameter := stream[1:3]
	channel := uint8(eventType) & 0x0f
	eventType = eventType & 0xf0
	sizeOfMIDIControlEvent := 3

	switch eventType {
	case NoteOff:
		event = &NoteOffEvent{
			deltaTime: deltaTime,
			channel:   channel,
			note:      Note(parameter[0]),
			velocity:  parameter[1],
		}
	case NoteOn:
		event = &NoteOnEvent{
			deltaTime: deltaTime,
			channel:   channel,
			note:      Note(parameter[0]),
			velocity:  parameter[1],
		}
	case NoteAfterTouch:
		event = &NoteAfterTouchEvent{
			deltaTime: deltaTime,
			channel:   channel,
			note:      Note(parameter[0]),
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
		sizeOfMIDIControlEvent = 2
		event = &ProgramChangeEvent{
			deltaTime: deltaTime,
			channel:   channel,
			program:   uint8(parameter[0]),
		}
	case ChannelAfterTouch:
		sizeOfMIDIControlEvent = 2
		event = &ChannelAfterTouchEvent{
			deltaTime: deltaTime,
			channel:   channel,
			velocity:  uint8(parameter[0]),
		}
	case PitchBend:
		event = &NoteOffEvent{
			deltaTime: deltaTime,
			channel:   channel,
			note:      Note(parameter[0]),
			velocity:  parameter[1],
		}
	default:
		return nil, 0, fmt.Errorf("midi: invalid MIDI control event")
	}

	sizeOfEvent = len(deltaTime.Quantity().Value()) + sizeOfMIDIControlEvent

	return event, sizeOfEvent, nil
}
