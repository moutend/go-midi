package midi

import (
	"fmt"
	"log"
)

type Parser struct {
	data              []byte
	position          int
	previousEventType uint8
	logger            *log.Logger
}

func (p *Parser) debugf(format string, v ...interface{}) {
	if p.logger == nil {
		return
	}
	p.logger.Print(fmt.Sprintf("midi: [%v] ", p.position))
	p.logger.Printf(format, v...)
}

func (p *Parser) debugln(v ...interface{}) {
	if p.logger == nil {
		return
	}
	p.logger.Print(fmt.Sprintf("midi: [%v] ", p.position))
	p.logger.Println(v...)
}

// Parse parses standard MIDI (*.mid) data.
func (p *Parser) Parse(stream []byte) (*MIDI, error) {
	p.debugf("start parsing %v bytes\n", len(stream))

	formatType, numberOfTracks, timeDivision, err := p.parseHeader()
	if err != nil {
		return nil, err
	}

	tracks, err := p.parseTracks(numberOfTracks)
	if err != nil {
		return nil, err
	}

	midi := &MIDI{
		formatType:   formatType,
		timeDivision: &TimeDivision{value: timeDivision},
		Tracks:       tracks,
	}

	p.debugln("successfully done")

	return midi, nil
}

// parseHeader parses stream begins with MThd.
func (p *Parser) parseHeader() (formatType, numberOfTracks, timeDivision uint16, err error) {
	p.debugf("start parsing MThd")

	mthd := string(p.data[p.position:4])
	if mthd != "MThd" {
		return formatType, numberOfTracks, timeDivision, fmt.Errorf("midi: invalid chunk ID %v", mthd)
	}

	p.position += 4
	p.debugln("parsing MThd completed")

	p.debugln("start parsing header size")

	headerSize := p.data[p.position+3]
	if headerSize != 6 {
		return formatType, numberOfTracks, timeDivision, fmt.Errorf("midi: header size must be always 6 bytes (%v)", headerSize)
	}

	p.position += 4
	p.debugln("parsing header size completed")

	p.debugln("start parsing format type")

	formatType = uint16(p.data[p.position+1])
	if formatType > 3 {
		return formatType, numberOfTracks, timeDivision, fmt.Errorf("midi: format type should be 1, 2 or 3")
	}

	p.position += 2
	p.debugf("parsing format type completed (formatType=%v)", formatType)

	p.debugln("start parsing number of tracks")

	numberOfTracks = uint16(p.data[p.position])
	numberOfTracks = numberOfTracks << 8
	numberOfTracks += uint16(p.data[p.position+1])

	p.position += 2
	p.debugf("parsing number of tracks completed (%v)", numberOfTracks)

	p.debugln("start parsing time division")

	timeDivision = uint16(p.data[p.position])
	timeDivision = timeDivision << 8
	timeDivision += uint16(p.data[p.position+1])

	p.position += 2
	p.debugf("parsing time division completed (timeDivision = %v)", timeDivision)

	return formatType, numberOfTracks, timeDivision, nil
}

// parseTracks parses stream begins with MTrk.
func (p *Parser) parseTracks(numberOfTracks uint16) ([]*Track, error) {
	tracks := make([]*Track, numberOfTracks)

	for n := 0; n < int(numberOfTracks); n++ {
		p.debugln("start parsing MTrk")

		mtrk := string(p.data[p.position : p.position+4])
		if mtrk != "MTrk" {
			return nil, fmt.Errorf("midi: invalid track ID %v", mtrk)
		}

		p.position += 4
		p.debugln("parsing MTrk completed")

		p.debugln("start parsing size of track")

		chunkSize := uint32(p.data[p.position])
		chunkSize = chunkSize << 8
		chunkSize += uint32(p.data[p.position+1])
		chunkSize = chunkSize << 8
		chunkSize += uint32(p.data[p.position+2])
		chunkSize = chunkSize << 8
		chunkSize += uint32(p.data[p.position+3])

		p.position += 4
		p.debugf("parsing size of track completed (chunkSize=%v)", chunkSize)

		track, err := p.parseTrack()
		if err != nil {
			return nil, err
		}

		tracks[n] = track
	}

	return tracks, nil
}

// parseTrack parses stream begins with delta time and ends with end of track event.
func (p *Parser) parseTrack() (*Track, error) {
	sizeOfStream := len(p.data)
	track := &Track{
		Events: []Event{},
	}
	for {
		if p.position >= sizeOfStream {
			break
		}

		event, err := p.parseEvent()
		if err != nil {
			return nil, err
		}

		track.Events = append(track.Events, event)

		switch event.(type) {
		case *EndOfTrackEvent:
			return track, nil
		}
	}

	return nil, fmt.Errorf("midi: missing end of track event")
}

// parseEvent parses stream begins with delta time.
func (p *Parser) parseEvent() (event Event, err error) {
	p.debugln("start parsing delta time")

	deltaTime, err := parseDeltaTime(p.data[p.position:])
	if err != nil {
		return nil, err
	}

	p.position += len(deltaTime.Quantity().Value())
	p.debugf("parsing delta time completed (%v)", deltaTime.Quantity().Uint32())

	p.debugln("start parsing event type")

	runningStatus := false
	eventType := p.data[p.position]
	p.position += 1

	if eventType < 0x80 && p.previousEventType >= 0x80 {
		p.debugln("running status enabled for this event")

		runningStatus = true
		eventType = p.previousEventType
		p.position -= 1
	}

	p.previousEventType = eventType
	p.debugf("parsing event type completed (0x%x)", eventType)

	switch eventType {
	case Meta:
		event, err = p.parseMetaEvent(eventType)
	case SystemExclusive, DividedSystemExclusive:
		event, err = p.parseSystemExclusiveEvent(eventType)
	default:
		event, err = p.parseMIDIControlEvent(eventType)
	}

	event.DeltaTime().Quantity().SetValue(deltaTime.Quantity().Value())
	event.SetRunningStatus(runningStatus)

	return event, err
}

// parseMetaEvent parses
func (p *Parser) parseMetaEvent(eventType uint8) (event Event, err error) {
	p.debugln("start parsing meta event type")

	metaEventType := p.data[p.position]

	p.position += 1
	p.debugf("parsing meta event type completed (0x%x)", metaEventType)

	p.debugln("start parsing size of meta event")

	q, err := parseQuantity(p.data[p.position:])
	if err != nil {
		return nil, err
	}

	p.position += len(q.value)
	p.debugf("parsing size of meta event completed (%v)", q.Uint32())

	sizeOfData := int(q.Uint32())
	data := p.data[p.position : p.position+sizeOfData]

	switch metaEventType {
	case Text:
		event = &TextEvent{
			text: data,
		}
	case CopyrightNotice:
		event = &CopyrightNoticeEvent{
			text: data,
		}
	case SequenceOrTrackName:
		event = &SequenceOrTrackNameEvent{
			text: data,
		}
	case InstrumentName:
		event = &InstrumentNameEvent{
			text: data,
		}
	case Lyrics:
		event = &LyricsEvent{
			text: data,
		}
	case Marker:
		event = &MarkerEvent{
			text: data,
		}
	case CuePoint:
		event = &CuePointEvent{
			text: data,
		}
	case MIDIPortPrefix:
		event = &MIDIPortPrefixEvent{
			port: uint8(data[0]),
		}
	case MIDIChannelPrefix:
		event = &MIDIChannelPrefixEvent{
			channel: uint8(data[0]),
		}
	case SetTempo:
		tempo := uint32(data[0])
		tempo = (tempo << 8) + uint32(data[1])
		tempo = (tempo << 8) + uint32(data[2])
		event = &SetTempoEvent{
			tempo: tempo,
		}
	case SMPTEOffset:
		event = &SMPTEOffsetEvent{
			hour:     data[0],
			minute:   data[1],
			second:   data[2],
			frame:    data[3],
			subFrame: data[4],
		}
	case TimeSignature:
		event = &TimeSignatureEvent{
			numerator:      uint8(data[0]),
			denominator:    uint8(data[1]),
			metronomePulse: uint8(data[2]),
			quarterNote:    uint8(data[3]),
		}
	case KeySignature:
		event = &KeySignatureEvent{
			key:   int8(data[0]),
			scale: uint8(data[1]),
		}
	case SequencerSpecific:
		event = &SequencerSpecificEvent{
			data: data,
		}
	case EndOfTrack:
		event = &EndOfTrackEvent{}
	default:
		event = &AlienEvent{
			metaEventType: metaEventType,
			data:          data,
		}
	}

	p.position += sizeOfData
	p.debugf("parsing event completed (event = %v)", event)

	return event, nil
}

// parseSystemExclusiveEvent parses
func (p *Parser) parseSystemExclusiveEvent(eventType uint8) (event Event, err error) {
	p.debugln("start parsing size of system exclusive event")

	q, err := parseQuantity(p.data[p.position:])
	if err != nil {
		return nil, err
	}

	p.position += len(q.value)
	p.debugf("parsing size of system exclusive event completed (%v)", q.Uint32())

	sizeOfData := int(q.Uint32())
	data := p.data[p.position : p.position+sizeOfData]

	switch eventType {
	case SystemExclusive:
		event = &SystemExclusiveEvent{
			data: data,
		}
	case DividedSystemExclusive:
		event = &DividedSystemExclusiveEvent{
			data: data,
		}
	}

	p.position += sizeOfData
	p.debugf("parsing event completed (event = %v)", event)

	return event, nil
}

// parseMIDIControlEvent parses
func (p *Parser) parseMIDIControlEvent(eventType uint8) (event Event, err error) {
	p.debugln("start parsing MIDI control event")

	channel := uint8(eventType) & 0x0f
	sizeOfData := 2
	data := p.data[p.position : p.position+sizeOfData]

	switch eventType & 0xf0 {
	case NoteOff:
		event = &NoteOffEvent{
			channel:  channel,
			note:     Note(data[0]),
			velocity: data[1],
		}
	case NoteOn:
		event = &NoteOnEvent{
			channel:  channel,
			note:     Note(data[0]),
			velocity: data[1],
		}
	case NoteAfterTouch:
		event = &NoteAfterTouchEvent{
			channel:  channel,
			note:     Note(data[0]),
			velocity: data[1],
		}
	case Controller:
		event = &ControllerEvent{
			channel: channel,
			control: Control(data[0]),
			value:   data[1],
		}
	case ProgramChange:
		sizeOfData = 1
		event = &ProgramChangeEvent{
			channel: channel,
			program: GM(data[0]),
		}
	case ChannelAfterTouch:
		sizeOfData = 1
		event = &ChannelAfterTouchEvent{
			channel:  channel,
			velocity: data[0],
		}
	case PitchBend:
		pitch := uint16(data[0]&0x7f) << 7
		pitch += uint16(data[1] & 0x7f)
		event = &PitchBendEvent{
			channel: channel,
			pitch:   pitch,
		}
	}

	p.position += sizeOfData
	p.debugf("parsing event completed (event = %v)", event)

	return event, nil
}

// SetLogger sets logger.
func (p *Parser) SetLogger(logger *log.Logger) *Parser {
	p.logger = logger

	return p
}

// NewParser returns Parser.
func NewParser(data []byte) *Parser {
	return &Parser{
		data: data,
	}
}
