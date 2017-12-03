package midi

import (
	"fmt"
	"io/ioutil"
	"log"
)

type Parser struct {
	data              []byte
	position          int
	previousEventType uint8
	logger            *log.Logger
}

func (p *Parser) debugf(format string, v ...interface{}) {
	format = fmt.Sprintf("midi: [%v] %v", p.position, format)
	p.logger.Printf(format, v...)
}

func (p *Parser) debugln(v ...interface{}) {
	a := make([]interface{}, len(v)+1)
	a[0] = fmt.Sprintf("midi: [%v]", p.position)

	for i := 0; i < len(v); i++ {
		a[i+1] = v[i]
	}

	p.logger.Println(a...)
}

// Parse parses standard MIDI (*.mid) data.
func (p *Parser) Parse(stream []byte) (*MIDI, error) {
	p.logger.Printf("start parsing %v bytes\n", len(stream))

	formatType, numberOfTracks, timeDivision, err := p.parseHeader(stream)
	if err != nil {
		return nil, err
	}

	tracks, err := p.parseTracks(stream[p.position:], int(numberOfTracks))
	if err != nil {
		return nil, err
	}

	midi := &MIDI{
		formatType:   formatType,
		timeDivision: &TimeDivision{value: timeDivision},
		Tracks:       tracks,
	}

	p.logger.Println("successfully done")

	return midi, nil
}

// parseHeader parses stream begins with MThd.
func (p *Parser) parseHeader(stream []byte) (formatType, numberOfTracks, timeDivision uint16, err error) {
	p.debugf("start parsing MThd")

	if string(stream[:4]) != "MThd" {
		return formatType, numberOfTracks, timeDivision, fmt.Errorf("midi: invalid chunk ID %v", stream[:4])
	}

	p.position += 4
	p.debugln("parsing MThd completed")

	p.position += 4 // skip read header size
	p.debugln("skip parsing size of header chunk")

	p.debugln("start parsing format type")

	formatType = uint16(stream[p.position+1])

	p.position += 2
	p.debugf("parsing format type completed (formatType=%v)", formatType)

	p.debugln("start parsing number of tracks")

	numberOfTracks = uint16(stream[p.position])
	numberOfTracks = numberOfTracks << 8
	numberOfTracks += uint16(stream[p.position+1])

	p.position += 2
	p.debugf("parsing number of tracks completed (%v)", numberOfTracks)

	p.debugln("start parsing time division")

	timeDivision = uint16(stream[p.position])
	timeDivision = timeDivision << 8
	timeDivision += uint16(stream[p.position+1])

	p.position += 2
	p.debugf("parsing time division completed (timeDivision = %v)", timeDivision)

	return formatType, numberOfTracks, timeDivision, nil
}

// parseTracks parses stream begins with MTrk.
func (p *Parser) parseTracks(stream []byte, numberOfTracks int) ([]*Track, error) {
	start := 0
	tracks := make([]*Track, numberOfTracks)

	for n := 0; n < numberOfTracks; n++ {
		p.debugln("start parsing MTrk")
		if string(stream[start:start+4]) != "MTrk" {
			return nil, fmt.Errorf("midi: invalid track ID %v", stream[start:start+4])
		}

		start += 4
		p.debugln("parsing MTrk completed")

		p.debugln("start parsing size of track")

		chunkSize := uint32(stream[start])
		chunkSize = chunkSize << 8
		chunkSize += uint32(stream[start+1])
		chunkSize = chunkSize << 8
		chunkSize += uint32(stream[start+2])
		chunkSize = chunkSize << 8
		chunkSize += uint32(stream[start+3])

		start += 4
		p.debugf("parsing size of track completed (chunkSize=%v)", chunkSize)

		track, err := p.parseTrack(stream[start:])
		if err != nil {
			return nil, err
		}

		tracks[n] = track
		start += int(chunkSize)
	}

	return tracks, nil
}

// parseTrack parses stream begins with delta time and ends with end of track event.
func (p *Parser) parseTrack(stream []byte) (*Track, error) {
	start := 0
	sizeOfStream := len(stream)
	track := &Track{
		Events: []Event{},
	}
	for {
		if start >= sizeOfStream {
			break
		}

		event, sizeOfEvent, err := p.parseEvent(stream[start:])
		if err != nil {
			return nil, err
		}
		track.Events = append(track.Events, event)
		start += sizeOfEvent

		switch event.(type) {
		case *EndOfTrackEvent:
			return track, nil
		}
	}

	return nil, fmt.Errorf("midi: missing end of track event")
}

// parseEvent parses stream begins with delta time.
func (p *Parser) parseEvent(stream []byte) (event Event, sizeOfEvent int, err error) {
	p.debugln("start parsing event")

	deltaTime, err := parseDeltaTime(stream)
	if err != nil {
		return nil, 0, err
	}

	sizeOfDeltaTime := len(deltaTime.Quantity().value)
	eventType := stream[sizeOfDeltaTime]

	if eventType < 0x80 && p.previousEventType >= 0x80 {
		eventType = p.previousEventType
	}
	switch eventType {
	case Meta:
		return p.parseMetaEvent(stream[sizeOfDeltaTime:], deltaTime)
	case SystemExclusive, DividedSystemExclusive:
		return p.parseSystemExclusiveEvent(stream[sizeOfDeltaTime:], deltaTime)
	default:
		return p.parseMIDIControlEvent(stream[sizeOfDeltaTime:], deltaTime, eventType)
	}
}

// parseMetaEvent parses stream begins with 0xff.
func (p *Parser) parseMetaEvent(stream []byte, deltaTime *DeltaTime) (event Event, sizeOfEvent int, err error) {
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
		tempo := uint32(metaEventData[0])
		tempo = (tempo << 8) + uint32(metaEventData[1])
		tempo = (tempo << 8) + uint32(metaEventData[2])
		event = &SetTempoEvent{
			deltaTime: deltaTime,
			tempo:     tempo,
		}
	case SMPTEOffset:
		event = &SMPTEOffsetEvent{
			deltaTime: deltaTime,
			hour:      metaEventData[0],
			minute:    metaEventData[1],
			second:    metaEventData[2],
			frame:     metaEventData[3],
			subFrame:  metaEventData[4],
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
	case SequencerSpecific:
		event = &SequencerSpecificEvent{
			deltaTime: deltaTime,
			data:      metaEventData,
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

	p.previousEventType = Meta
	p.debugf("parsing event completed (event = %v)", event)

	return event, sizeOfEvent, nil
}

// parseSystemExclusiveEvent parses stream begins with 0xf0 or 0xf7.
func (p *Parser) parseSystemExclusiveEvent(stream []byte, deltaTime *DeltaTime) (event Event, sizeOfEvent int, err error) {
	q, err := parseQuantity(stream[1:])
	if err != nil {
		return nil, 0, err
	}

	offset := 1 + len(q.value)
	sizeOfData := int(q.Uint32())
	sizeOfEvent = len(deltaTime.Quantity().value) + offset + sizeOfData
	eventType := stream[0]

	switch eventType {
	case SystemExclusive:
		event = &SystemExclusiveEvent{
			deltaTime: deltaTime,
			data:      stream[offset : offset+sizeOfData],
		}
	case DividedSystemExclusive:
		event = &DividedSystemExclusiveEvent{
			deltaTime: deltaTime,
			data:      stream[offset : offset+sizeOfData],
		}
	}

	p.previousEventType = eventType
	p.debugf("parsing event completed (event = %v)", event)

	return event, sizeOfEvent, nil
}

// parseMIDIControlEvent parses stream begins with 0x8_...0xe_.
func (p *Parser) parseMIDIControlEvent(stream []byte, deltaTime *DeltaTime, eventType byte) (event Event, sizeOfEvent int, err error) {
	parameter := stream[1:3]
	channel := uint8(eventType) & 0x0f
	sizeOfMIDIControlEvent := 3

	switch eventType & 0xf0 {
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
			control:   Control(parameter[0]),
			value:     uint8(parameter[1]),
		}
	case ProgramChange:
		sizeOfMIDIControlEvent = 2
		event = &ProgramChangeEvent{
			deltaTime: deltaTime,
			channel:   channel,
			program:   GM(parameter[0]),
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
		sizeOfMIDIControlEvent = 2
		event = &ContinuousControllerEvent{
			deltaTime: deltaTime,
			control:   uint8(stream[0]),
			value:     uint8(stream[1]),
		}
	}

	sizeOfEvent = len(deltaTime.Quantity().Value()) + sizeOfMIDIControlEvent

	p.previousEventType = eventType
	p.debugf("parsing event completed (event = %v)", event)

	return event, sizeOfEvent, nil
}

func parseDeltaTime(stream []byte) (*DeltaTime, error) {
	q, err := parseQuantity(stream)
	if err != nil {
		return nil, err
	}

	deltaTime := &DeltaTime{q}

	return deltaTime, nil
}

func parseQuantity(stream []byte) (*Quantity, error) {
	if len(stream) == 0 {
		return nil, fmt.Errorf("midi: stream is empty")
	}

	var i int
	q := &Quantity{}

	for {
		if i > 3 {
			return nil, fmt.Errorf("midi: maximum size of variable quantity is 4 bytes")
		}
		if len(stream) < (i + 1) {
			return nil, fmt.Errorf("midi: missing next byte")
		}
		if stream[i] < 0x80 {
			break
		}
		i++
	}

	q.value = make([]byte, i+1)
	copy(q.value, stream)

	return q, nil
}

// SetLogger sets logger.
func (p *Parser) SetLogger(logger *log.Logger) *Parser {
	if logger != nil {
		p.logger = logger
	}

	return p
}

// NewParser returns Parser.
func NewParser(data []byte) *Parser {
	return &Parser{
		data:   data,
		logger: log.New(ioutil.Discard, "discard logging messages", 0),
	}
}
