package midi

import "fmt"

// Event represents any MIDI events, including meta and system exclusive.
type Event interface {
	DeltaTime() *DeltaTime
}

// parseEvent parses stream begins with delta time.
func parseEvent(stream []byte) (Event, int, error) {
	deltaTime, err := parseDeltaTime(stream)
	if err != nil {
		return nil, 0, err
	}

	sizeOfDeltaTime := len(deltaTime.Quantity().value)
	eventType := stream[sizeOfDeltaTime]

	switch eventType {
	case Meta:
		return parseMetaEvent(stream[sizeOfDeltaTime:], deltaTime)
	case SystemExclusive, DividedSystemExclusive:
		return parseSystemExclusiveEvent(stream[sizeOfDeltaTime:], deltaTime)
	default:
		return parseMIDIControlEvent(stream, deltaTime, eventType)
	}
}

// parseMetaEvent parses stream begins with 0xff.
func parseMetaEvent(stream []byte, deltaTime *DeltaTime) (event Event, sizeOfEvent int, err error) {
	q, err := parseQuantity(stream[2:])
	if err != nil {
		return nil, 0, err
	}

	offset := 2 + len(q.value)
	sizeOfData := int(q.Uint32())
	sizeOfEvent = len(deltaTime.Quantity().Value()) + offset + sizeOfData

	metaEventType := stream[1]
	metaEventData := stream[offset : offset+sizeOfData]

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
			deltaTime:      deltaTime,
			numerator:      uint8(metaEventData[0]),
			denominator:    uint8(metaEventData[1]),
			metronomePulse: uint8(metaEventData[2]),
			quarterNote:    uint8(metaEventData[3]),
		}
	case KeySignature:
		event = &KeySignatureEvent{
			deltaTime: deltaTime,
			key:       int8(metaEventData[0]),
			scale:     uint8(metaEventData[1]),
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

	return event, sizeOfEvent, nil
}

// parseSystemExclusiveEvent parses stream begins with 0xf0 or 0xf7.
func parseSystemExclusiveEvent(stream []byte, deltaTime *DeltaTime) (event Event, sizeOfEvent int, err error) {
	q, err := parseQuantity(stream[1:])
	if err != nil {
		return nil, 0, err
	}

	offset := 1 + len(q.value)
	sizeOfData := int(q.Uint32())
	sizeOfEvent = len(deltaTime.Quantity().value) + offset + sizeOfData

	event = &SystemExclusiveEvent{
		deltaTime: deltaTime,
		data:      stream[offset : offset+sizeOfData],
	}

	return event, sizeOfEvent, nil
}

// parseMetaEvent parses stream begins with 0x8_...0xe_.
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
		pitch := uint16(parameter[0]&0x7f) << 7
		pitch += uint16(parameter[1] & 0x7f)
		event = &PitchBendEvent{
			deltaTime: deltaTime,
			channel:   channel,
			pitch:     pitch,
		}
	default:
		return nil, 0, fmt.Errorf("midi: invalid MIDI control event")
	}

	sizeOfEvent = len(deltaTime.Quantity().Value()) + sizeOfMIDIControlEvent

	return event, sizeOfEvent, nil
}
