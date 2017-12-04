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

	p.logger.Println("successfully done")

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
		return p.parseMetaEvent(deltaTime, eventType, runningStatus)
	case SystemExclusive, DividedSystemExclusive:
		return p.parseSystemExclusiveEvent(deltaTime, eventType, runningStatus)
	default:
		return p.parseMIDIControlEvent(deltaTime, eventType, runningStatus)
	}
}

// parseMetaEvent parses
func (p *Parser) parseMetaEvent(deltaTime *DeltaTime, eventType uint8, runningStatus bool) (event Event, err error) {
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
			deltaTime: deltaTime,
			text:      data,
		}
	case CopyrightNotice:
		event = &CopyrightNoticeEvent{
			deltaTime: deltaTime,
			text:      data,
		}
	case SequenceOrTrackName:
		event = &SequenceOrTrackNameEvent{
			deltaTime: deltaTime,
			text:      data,
		}
	case InstrumentName:
		event = &InstrumentNameEvent{
			deltaTime: deltaTime,
			text:      data,
		}
	case Lyrics:
		event = &LyricsEvent{
			deltaTime: deltaTime,
			text:      data,
		}
	case Marker:
		event = &MarkerEvent{
			deltaTime: deltaTime,
			text:      data,
		}
	case CuePoint:
		event = &CuePointEvent{
			deltaTime: deltaTime,
			text:      data,
		}
	case MIDIPortPrefix:
		event = &MIDIPortPrefixEvent{
			deltaTime: deltaTime,
			port:      uint8(data[0]),
		}
	case MIDIChannelPrefix:
		event = &MIDIChannelPrefixEvent{
			deltaTime: deltaTime,
			channel:   uint8(data[0]),
		}
	case SetTempo:
		tempo := uint32(data[0])
		tempo = (tempo << 8) + uint32(data[1])
		tempo = (tempo << 8) + uint32(data[2])
		event = &SetTempoEvent{
			deltaTime: deltaTime,
			tempo:     tempo,
		}
	case SMPTEOffset:
		event = &SMPTEOffsetEvent{
			deltaTime: deltaTime,
			hour:      data[0],
			minute:    data[1],
			second:    data[2],
			frame:     data[3],
			subFrame:  data[4],
		}
	case TimeSignature:
		event = &TimeSignatureEvent{
			deltaTime:      deltaTime,
			numerator:      uint8(data[0]),
			denominator:    uint8(data[1]),
			metronomePulse: uint8(data[2]),
			quarterNote:    uint8(data[3]),
		}
	case KeySignature:
		event = &KeySignatureEvent{
			deltaTime: deltaTime,
			key:       int8(data[0]),
			scale:     uint8(data[1]),
		}
	case SequencerSpecific:
		event = &SequencerSpecificEvent{
			deltaTime: deltaTime,
			data:      data,
		}
	case EndOfTrack:
		event = &EndOfTrackEvent{
			deltaTime: deltaTime,
		}
	default:
		event = &AlienEvent{
			deltaTime:     deltaTime,
			metaEventType: metaEventType,
			data:          data,
		}
	}

	p.position += sizeOfData
	p.debugf("parsing event completed (event = %v)", event)

	return event, nil
}

// parseSystemExclusiveEvent parses
func (p *Parser) parseSystemExclusiveEvent(deltaTime *DeltaTime, eventType uint8, runningStatus bool) (event Event, err error) {
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
			deltaTime: deltaTime,
			data:      data,
		}
	case DividedSystemExclusive:
		event = &DividedSystemExclusiveEvent{
			deltaTime: deltaTime,
			data:      data,
		}
	}

	p.position += sizeOfData
	p.debugf("parsing event completed (event = %v)", event)

	return event, nil
}

// parseMIDIControlEvent parses
func (p *Parser) parseMIDIControlEvent(deltaTime *DeltaTime, eventType uint8, runningStatus bool) (event Event, err error) {
	p.debugln("start parsing MIDI control event")

	channel := uint8(eventType) & 0x0f
	sizeOfData := 2
	data := p.data[p.position : p.position+sizeOfData]

	switch eventType & 0xf0 {
	case NoteOff:
		event = &NoteOffEvent{
			deltaTime: deltaTime,
			channel:   channel,
			note:      Note(data[0]),
			velocity:  data[1],
		}
	case NoteOn:
		event = &NoteOnEvent{
			deltaTime: deltaTime,
			channel:   channel,
			note:      Note(data[0]),
			velocity:  data[1],
		}
	case NoteAfterTouch:
		event = &NoteAfterTouchEvent{
			deltaTime: deltaTime,
			channel:   channel,
			note:      Note(data[0]),
			velocity:  data[1],
		}
	case Controller:
		event = &ControllerEvent{
			deltaTime: deltaTime,
			channel:   channel,
			control:   Control(data[0]),
			value:     data[1],
		}
	case ProgramChange:
		sizeOfData = 1
		event = &ProgramChangeEvent{
			deltaTime: deltaTime,
			channel:   channel,
			program:   GM(data[0]),
		}
	case ChannelAfterTouch:
		sizeOfData = 1
		event = &ChannelAfterTouchEvent{
			deltaTime: deltaTime,
			channel:   channel,
			velocity:  data[0],
		}
	case PitchBend:
		pitch := uint16(data[0]&0x7f) << 7
		pitch += uint16(data[1] & 0x7f)
		event = &PitchBendEvent{
			deltaTime: deltaTime,
			channel:   channel,
			pitch:     pitch,
		}
	default:
		return nil, fmt.Errorf("undefined")
		//sizeOfMIDIControlEvent = 2
		event = &ContinuousControllerEvent{
			deltaTime: deltaTime,
		}
	}

	p.position += sizeOfData
	p.debugf("parsing event completed (event = %v)", event)

	return event, nil
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
