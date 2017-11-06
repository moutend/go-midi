package midi

import (
	"bytes"
	"encoding/binary"
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
