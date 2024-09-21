package midi

import (
	"fmt"
	"log"

	"github.com/moutend/go-midi/constant"
	"github.com/moutend/go-midi/deltatime"
	"github.com/moutend/go-midi/event"
	"github.com/moutend/go-midi/quantity"
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
	format = fmt.Sprintf("midi: [%v] %v", p.position, format)
	p.logger.Printf(format, v...)
}

func (p *Parser) debugln(v ...interface{}) {
	if p.logger == nil {
		return
	}
	prefix := fmt.Sprintf("midi: [%v]", p.position)
	a := []interface{}{prefix}
	a = append(a, v...)
	p.logger.Println(a...)
}

// Parse parses standard MIDI (*.mid) data.
func (p *Parser) Parse() (*MIDI, error) {
	p.debugf("start parsing %v bytes\n", len(p.data))

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
		Events: []event.Event{},
	}
	for {
		if p.position >= sizeOfStream {
			break
		}

		e, err := p.parseEvent()
		if err != nil {
			return nil, err
		}

		track.Events = append(track.Events, e)

		switch e.(type) {
		case *event.EndOfTrackEvent:
			return track, nil
		}
	}

	return nil, fmt.Errorf("midi: missing end of track event")
}

// parseEvent parses stream begins with delta time.
func (p *Parser) parseEvent() (event event.Event, err error) {
	
	p.debugln("start parsing delta time")

	deltaTime, err := deltatime.Parse(p.data[p.position:])
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
	case constant.Meta:
		event, err = p.parseMetaEvent(eventType)
	case constant.SystemExclusive, constant.DividedSystemExclusive:
		event, err = p.parseSystemExclusiveEvent(eventType)
	default:
		event, err = p.parseMIDIControlEvent(eventType)
	}

	if event == nil {
		return nil, fmt.Errorf("Event is null")
	}

	event.DeltaTime().Quantity().SetValue(deltaTime.Quantity().Value())
	event.SetRunningStatus(runningStatus)

	return event, err
}

// parseMetaEvent parses
func (p *Parser) parseMetaEvent(eventType uint8) (e event.Event, err error) {
	p.debugln("start parsing meta event type")

	metaEventType := p.data[p.position]

	p.position += 1
	p.debugf("parsing meta event type completed (0x%x)", metaEventType)

	p.debugln("start parsing size of meta event")

	q, err := quantity.Parse(p.data[p.position:])
	if err != nil {
		return nil, err
	}

	p.position += len(q.Value())
	p.debugf("parsing size of meta event completed (%v)", q.Uint32())

	sizeOfData := int(q.Uint32())
	data := p.data[p.position : p.position+sizeOfData]

	switch metaEventType {
	case constant.Text:
		v := &event.TextEvent{}
		v.SetText(data)
		e = v
	case constant.CopyrightNotice:
		v := &event.CopyrightNoticeEvent{}
		v.SetText(data)
		e = v
	case constant.SequenceOrTrackName:
		v := &event.SequenceOrTrackNameEvent{}
		v.SetText(data)
		e = v
	case constant.InstrumentName:
		v := &event.InstrumentNameEvent{}
		v.SetText(data)
		e = v
	case constant.Lyrics:
		v := &event.LyricsEvent{}
		v.SetText(data)
		e = v
	case constant.Marker:
		v := &event.MarkerEvent{}
		v.SetText(data)
		e = v
	case constant.CuePoint:
		v := &event.CuePointEvent{}
		v.SetText(data)
		e = v
	case constant.MIDIPortPrefix:
		v := &event.MIDIPortPrefixEvent{}
		v.SetPort(data[0])
		e = v
	case constant.MIDIChannelPrefix:
		v := &event.MIDIChannelPrefixEvent{}
		v.SetChannel(data[0])
		e = v
	case constant.SetTempo:
		tempo := uint32(data[0])
		tempo = (tempo << 8) + uint32(data[1])
		tempo = (tempo << 8) + uint32(data[2])
		v := &event.SetTempoEvent{}
		v.SetTempo(tempo)
		e = v
	case constant.SMPTEOffset:
		v := &event.SMPTEOffsetEvent{}
		v.SetHour(data[0])
		v.SetMinute(data[1])
		v.SetSecond(data[2])
		v.SetFrame(data[3])
		v.SetSubFrame(data[4])
		e = v
	case constant.TimeSignature:
		v := &event.TimeSignatureEvent{}
		v.SetNumerator(data[0])
		v.SetDenominator(data[1])
		v.SetMetronomePulse(data[2])
		v.SetQuarterNote(data[3])
		e = v
	case constant.KeySignature:
		v := &event.KeySignatureEvent{}
		v.SetKey(int8(data[0]))
		v.SetScale(data[1])
	case constant.SequencerSpecific:
		v := &event.SequencerSpecificEvent{}
		v.SetData(data)
		e = v
	case constant.EndOfTrack:
		e = &event.EndOfTrackEvent{}
	default:
		v := &event.AlienEvent{}
		v.SetMetaEventType(metaEventType)
		v.SetData(data)
		e = v
	}

	p.position += sizeOfData
	p.debugf("parsing event completed (event = %v)", e)

	return e, nil
}

// parseSystemExclusiveEvent parses
func (p *Parser) parseSystemExclusiveEvent(eventType uint8) (e event.Event, err error) {
	p.debugln("start parsing size of system exclusive event")

	q, err := quantity.Parse(p.data[p.position:])
	if err != nil {
		return nil, err
	}

	p.position += len(q.Value())
	p.debugf("parsing size of system exclusive event completed (%v)", q.Uint32())

	sizeOfData := int(q.Uint32())
	data := p.data[p.position : p.position+sizeOfData]

	switch eventType {
	case constant.SystemExclusive:
		v := &event.SystemExclusiveEvent{}
		v.SetData(data)
		e = v
	case constant.DividedSystemExclusive:
		v := &event.DividedSystemExclusiveEvent{}
		v.SetData(data)
	}

	p.position += sizeOfData
	p.debugf("parsing event completed (event = %v)", e)

	return e, nil
}

// parseMIDIControlEvent parses
func (p *Parser) parseMIDIControlEvent(eventType uint8) (e event.Event, err error) {
	p.debugln("start parsing MIDI control event")

	channel := uint8(eventType) & 0x0f
	sizeOfData := 2
	data := p.data[p.position : p.position+sizeOfData]

	switch eventType & 0xf0 {
	case constant.NoteOff:
		v := &event.NoteOffEvent{}
		v.SetChannel(channel)
		v.SetNote(constant.Note(data[0]))
		v.SetVelocity(data[1])
		e = v
	case constant.NoteOn:
		v := &event.NoteOnEvent{}
		v.SetChannel(channel)
		v.SetNote(constant.Note(data[0]))
		v.SetVelocity(data[1])
		e = v
	case constant.NoteAfterTouch:
		v := &event.NoteAfterTouchEvent{}
		v.SetChannel(channel)
		v.SetNote(constant.Note(data[0]))
		v.SetVelocity(data[1])
		e = v
	case constant.Controller:
		v := &event.ControllerEvent{}
		v.SetChannel(channel)
		v.SetControl(constant.Control(data[0]))
		v.SetValue(data[1])
		e = v
	case constant.ProgramChange:
		sizeOfData = 1
		v := &event.ProgramChangeEvent{}
		v.SetChannel(channel)
		v.SetProgram(constant.GM(data[0]))
		e = v
	case constant.ChannelAfterTouch:
		sizeOfData = 1
		v := &event.ChannelAfterTouchEvent{}
		v.SetChannel(channel)
		v.SetVelocity(data[0])
		e = v
	case constant.PitchBend:
		pitch := uint16(data[0]&0x7f) << 7
		pitch += uint16(data[1] & 0x7f)
		v := &event.PitchBendEvent{}
		v.SetChannel(channel)
		v.SetPitch(pitch)
		e = v
	}

	p.position += sizeOfData
	p.debugf("parsing event completed (event = %v)", e)

	return e, nil
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
