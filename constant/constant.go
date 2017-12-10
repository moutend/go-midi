package constant

const (
	NoteOff                = 0x80
	NoteOn                 = 0x90
	NoteAfterTouch         = 0xa0
	Controller             = 0xb0
	ProgramChange          = 0xc0
	ChannelAfterTouch      = 0xd0
	PitchBend              = 0xe0
	SystemExclusive        = 0xf0
	DividedSystemExclusive = 0xf7
	Meta                   = 0xff
)

const (
	Text                = 0x01
	CopyrightNotice     = 0x02
	SequenceOrTrackName = 0x03
	InstrumentName      = 0x04
	Lyrics              = 0x05
	Marker              = 0x06
	CuePoint            = 0x07
	MIDIPortPrefix      = 0x20
	MIDIChannelPrefix   = 0x21
	SetTempo            = 0x51
	SMPTEOffset         = 0x54
	TimeSignature       = 0x58
	KeySignature        = 0x59
	SequencerSpecific   = 0x7f
	EndOfTrack          = 0x2f
)
