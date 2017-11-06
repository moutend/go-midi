package midi

type Note uint8

const (
	C3  Note = 0x3c
	Db3 Note = 0x3d
	D3  Note = 0x3e
	Eb3 Note = 0x3f
	E3  Note = 0x40
	F3  Note = 0x41
	Gb3 Note = 0x42
	G3  Note = 0x43
	Ab3 Note = 0x44
	A3  Note = 0x45
	Bb3 Note = 0x46
	B3  Note = 0x47
	C4  Note = 0x60
)

type ChunkId uint32

const (
	MThd ChunkId = 0x4d546864
	MTrk ChunkId = 0x4d54726B
)

type EventType uint8

const (
	NoteOff           EventType = 0x80
	NoteOn            EventType = 0x90
	NoteAfterTouch    EventType = 0xa0
	Controller        EventType = 0xb0
	ProgramChange     EventType = 0xc0
	ChannelAfterTouch EventType = 0xd0
	PitchBend         EventType = 0xe0
	NormalSysEx       EventType = 0xf0
	EndOfNormalSysEx  EventType = 0xf7
	Meta              EventType = 0xff
)

type MetaEventType uint8

const (
	Text                MetaEventType = 0x01
	CopyrightNotice     MetaEventType = 0x02
	SequenceOrTrackName MetaEventType = 0x03
	InstrumentName      MetaEventType = 0x04
	Lyrics              MetaEventType = 0x05
	Marker              MetaEventType = 0x06
	CuePoint            MetaEventType = 0x07
	MIDIPortPrefix      MetaEventType = 0x20
	MIDIChannelPrefix   MetaEventType = 0x21
	SetTempo            MetaEventType = 0x51
	SMPTEOffset         MetaEventType = 0x54
	TimeSignature       MetaEventType = 0x58
	KeySignature        MetaEventType = 0x59
	SequencerSpecific   MetaEventType = 0x7f
	EndOfTrack          MetaEventType = 0x2f
)

type Rhythm uint16

const (
	AcousticBassDrum Rhythm = 0x35
	BassDrum1        Rhythm = 0x36
	SideStick        Rhythm = 0x37
	AcousticSnare    Rhythm = 0x38
	HandClap         Rhythm = 0x39
	ElectricSnare    Rhythm = 0x40
	LowFloorTom      Rhythm = 0x41
	ClosedHiHat      Rhythm = 0x42
	HighFloorTom     Rhythm = 0x43
	PedalHiHat       Rhythm = 0x44
	LowTom           Rhythm = 0x45
	OpenHiHat        Rhythm = 0x46
	LowMidTom        Rhythm = 0x47
	HiMidTom         Rhythm = 0x48
	CrashCymbal1     Rhythm = 0x49
	HighTom          Rhythm = 0x50
	RideCymbal1      Rhythm = 0x51
	ChineseCymbal    Rhythm = 0x52
	RideBell         Rhythm = 0x53
	Tambourine       Rhythm = 0x54
	SplashCymbal     Rhythm = 0x55
	Cowbell          Rhythm = 0x56
	CrashCymbal2     Rhythm = 0x57
	Vibraslap        Rhythm = 0x58
	RideCymbal2      Rhythm = 0x59
	HiBongo          Rhythm = 0x60
	LowBongo         Rhythm = 0x61
	MuteHiConga      Rhythm = 0x62
	OpenHiConga      Rhythm = 0x63
	LowConga         Rhythm = 0x64
	HighTimbale      Rhythm = 0x65
	LowTimbale       Rhythm = 0x66
	HighAgogo        Rhythm = 0x67
	LowAgogo         Rhythm = 0x68
	Cabasa           Rhythm = 0x69
	Maracas          Rhythm = 0x70
	ShortWhistle     Rhythm = 0x71
	LongWhistle      Rhythm = 0x72
	ShortGuiro       Rhythm = 0x73
	LongGuiro        Rhythm = 0x74
	Claves           Rhythm = 0x75
	HiWoodBlock      Rhythm = 0x76
	LowWoodBlock     Rhythm = 0x77
	MuteCuica        Rhythm = 0x78
	OpenCuica        Rhythm = 0x79
	MuteTriangle     Rhythm = 0x80
	OpenTriangle     Rhythm = 0x81
)

type GM uint16

const (
	AcousticGrandPiano  GM = 0x00
	BrightAcousticPiano GM = 0x01
	ElectricGrandPiano  GM = 0x02
	HonkyTonkPiano      GM = 0x03
	ElectricPiano1      GM = 0x04
	ElectricPiano2      GM = 0x05
	Harpsichord         GM = 0x06
	Clavi               GM = 0x07
	Celesta             GM = 0x08
	Glockenspiel        GM = 0x09
	MusicBox            GM = 0x0a
	Vibraphone          GM = 0x0b
	Marimba             GM = 0x0c
	Xylophone           GM = 0x0d
	TubularBells        GM = 0x0e
	Dulcimer            GM = 0x0f
	DrawbarOrgan        GM = 0x10
	PercussiveOrgan     GM = 0x11
	RockOrgan           GM = 0x12
	ChurchOrgan         GM = 0x13
	ReedOrgan           GM = 0x14
	Accordion           GM = 0x15
	Harmonica           GM = 0x16
	TangoAccordion      GM = 0x17
	AcousticNylonGuitar GM = 0x18
	AcousticSteelGuitar GM = 0x19
	ElectricJazzGuitar  GM = 0x1a
	ElectricCleanGuitar GM = 0x1b
	ElectricMutedGuitar GM = 0x1c
	OverdrivenGuitar    GM = 0x1d
	DistortionGuitar    GM = 0x1e
	GuitarHarmonics     GM = 0x1f
	AcousticBass        GM = 0x20
	ElectricFingerBass  GM = 0x21
	ElectricPickBass    GM = 0x22
	FretlessBass        GM = 0x23
	SlapBass1           GM = 0x24
	SlapBass2           GM = 0x25
	SynthBass1          GM = 0x26
	SynthBass2          GM = 0x27
	Violin              GM = 0x28
	Viola               GM = 0x29
	Cello               GM = 0x2a
	Contrabass          GM = 0x2b
	TremoloStrings      GM = 0x2c
	PizzicatoStrings    GM = 0x2d
	OrchestralHarp      GM = 0x2e
	Timpani             GM = 0x2f
	StringEnsemble1     GM = 0x30
	StringEnsemble2     GM = 0x31
	SynthStrings1       GM = 0x32
	SynthStrings2       GM = 0x33
	ChoirAahs           GM = 0x34
	VoiceOohs           GM = 0x35
	SynthVoice          GM = 0x36
	OrchestraHit        GM = 0x37
	Trumpet             GM = 0x38
	Trombone            GM = 0x39
	Tuba                GM = 0x3a
	MutedTrumpet        GM = 0x3b
	FrenchHorn          GM = 0x3c
	BrassSection        GM = 0x3d
	SynthBrass1         GM = 0x3e
	SynthBrass2         GM = 0x3f
	SopranoSax          GM = 0x40
	AltoSax             GM = 0x41
	TenorSax            GM = 0x42
	BaritoneSax         GM = 0x43
	Oboe                GM = 0x44
	EnglishHorn         GM = 0x45
	Bassoon             GM = 0x46
	Clarinet            GM = 0x47
	Piccolo             GM = 0x48
	Flute               GM = 0x49
	Recorder            GM = 0x4a
	PanFlute            GM = 0x4b
	BlownBottle         GM = 0x4c
	Shakuhachi          GM = 0x4d
	Whistle             GM = 0x4e
	Ocarina             GM = 0x4f
	Lead1Square         GM = 0x50
	Lead2Sawtooth       GM = 0x51
	Lead3Calliope       GM = 0x52
	Lead4Chiff          GM = 0x53
	Lead5Charang        GM = 0x54
	Lead6Voice          GM = 0x55
	Lead7Fifths         GM = 0x56
	Lead8BassLead       GM = 0x57
	Pad1NewAge          GM = 0x58
	Pad2Warm            GM = 0x59
	Pad3Polysynth       GM = 0x5a
	Pad4Choir           GM = 0x5b
	Pad5Bowed           GM = 0x5c
	Pad6Metallic        GM = 0x5d
	Pad7Halo            GM = 0x5e
	Pad8Sweep           GM = 0x5f
	FX1Rain             GM = 0x60
	FX2Soundtrack       GM = 0x61
	FX3Crystal          GM = 0x62
	FX4Atmosphere       GM = 0x63
	FX5Brightness       GM = 0x64
	FX6Goblins          GM = 0x65
	FX7Echoes           GM = 0x66
	FX8SciFi            GM = 0x67
	Sitar               GM = 0x68
	Banjo               GM = 0x69
	Shamisen            GM = 0x6a
	Koto                GM = 0x6b
	Kalimba             GM = 0x6c
	Bagpipe             GM = 0x6d
	Fiddle              GM = 0x6e
	Shanai              GM = 0x6f
	TinkleBell          GM = 0x70
	Agogo               GM = 0x71
	SteelDrums          GM = 0x72
	Woodblock           GM = 0x73
	TaikoDrum           GM = 0x74
	MelodicTom          GM = 0x75
	SynthDrum           GM = 0x76
	ReverseCymbal       GM = 0x77
	GuitarFretNoise     GM = 0x78
	BreathNoise         GM = 0x79
	Seashore            GM = 0x7a
	BirdTweet           GM = 0x7b
	TelephoneRing       GM = 0x7c
	Helicopter          GM = 0x7d
	Applause            GM = 0x7e
	Gunshot             GM = 0x7f
)
