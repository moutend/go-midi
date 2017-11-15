//go:generate stringer -type=Note -output=note_string.go
//go:generate stringer -type=GM -output=gm_string.go
//go:generate stringer -type=Rhythm -output=rhythm_string.go

package midi

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

// Note
type Note byte

const (
	Cminus2 Note = iota
	Dbminus2
	Dminus2
	Ebminus2
	Eminus2
	Fminus2
	Gbminus2
	Gminus2
	Abminus2
	Aminus2
	Bbminus2
	Bminus2
	Cminus1
	Dbminus1
	Dminus1
	Ebminus1
	Eminus1
	Fminus1
	Gbminus1
	Gminus1
	Abminus1
	Aminus1
	Bbminus1
	Bminus1
	C0
	Db0
	D0
	Eb0
	E0
	F0
	Gb0
	G0
	Ab0
	A0
	Bb0
	B0
	C1
	Db1
	D1
	Eb1
	E1
	F1
	Gb1
	G1
	Ab1
	A1
	Bb1
	B1
	C2
	Db2
	D2
	Eb2
	E2
	F2
	Gb2
	G2
	Ab2
	A2
	Bb2
	B2
	C3
	Db3
	D3
	Eb3
	E3
	F3
	Gb3
	G3
	Ab3
	A3
	Bb3
	B3
	C4
	Db4
	D4
	Eb4
	E4
	F4
	Gb4
	G4
	Ab4
	A4
	Bb4
	B4
	C5
	Db5
	D5
	Eb5
	E5
	F5
	Gb5
	G5
	Ab5
	A5
	Bb5
	B5
	C6
	Db6
	D6
	Eb6
	E6
	F6
	Gb6
	G6
	Ab6
	A6
	Bb6
	B6
	C7
	Db7
	D7
	Eb7
	E7
	F7
	Gb7
	G7
	Ab7
	A7
	Bb7
	B7
	C8
	Db8
	D8
	Eb8
	E8
	F8
	Gb8
	G8
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
	AcousticGrandPiano GM = iota
	BrightAcousticPiano
	ElectricGrandPiano
	HonkyTonkPiano
	ElectricPiano1
	ElectricPiano2
	Harpsichord
	Clavi
	Celesta
	Glockenspiel
	MusicBox
	Vibraphone
	Marimba
	Xylophone
	TubularBells
	Dulcimer
	DrawbarOrgan
	PercussiveOrgan
	RockOrgan
	ChurchOrgan
	ReedOrgan
	Accordion
	Harmonica
	TangoAccordion
	AcousticNylonGuitar
	AcousticSteelGuitar
	ElectricJazzGuitar
	ElectricCleanGuitar
	ElectricMutedGuitar
	OverdrivenGuitar
	DistortionGuitar
	GuitarHarmonics
	AcousticBass
	ElectricFingerBass
	ElectricPickBass
	FretlessBass
	SlapBass1
	SlapBass2
	SynthBass1
	SynthBass2
	Violin
	Viola
	Cello
	Contrabass
	TremoloStrings
	PizzicatoStrings
	OrchestralHarp
	Timpani
	StringEnsemble1
	StringEnsemble2
	SynthStrings1
	SynthStrings2
	ChoirAahs
	VoiceOohs
	SynthVoice
	OrchestraHit
	Trumpet
	Trombone
	Tuba
	MutedTrumpet
	FrenchHorn
	BrassSection
	SynthBrass1
	SynthBrass2
	SopranoSax
	AltoSax
	TenorSax
	BaritoneSax
	Oboe
	EnglishHorn
	Bassoon
	Clarinet
	Piccolo
	Flute
	Recorder
	PanFlute
	BlownBottle
	Shakuhachi
	Whistle
	Ocarina
	Lead1Square
	Lead2Sawtooth
	Lead3Calliope
	Lead4Chiff
	Lead5Charang
	Lead6Voice
	Lead7Fifths
	Lead8BassLead
	Pad1NewAge
	Pad2Warm
	Pad3Polysynth
	Pad4Choir
	Pad5Bowed
	Pad6Metallic
	Pad7Halo
	Pad8Sweep
	FX1Rain
	FX2Soundtrack
	FX3Crystal
	FX4Atmosphere
	FX5Brightness
	FX6Goblins
	FX7Echoes
	FX8SciFi
	Sitar
	Banjo
	Shamisen
	Koto
	Kalimba
	Bagpipe
	Fiddle
	Shanai
	TinkleBell
	Agogo
	SteelDrums
	Woodblock
	TaikoDrum
	MelodicTom
	SynthDrum
	ReverseCymbal
	GuitarFretNoise
	BreathNoise
	Seashore
	BirdTweet
	TelephoneRing
	Helicopter
	Applause
	Gunshot
)

const (
	BankSelect                      = 0x00
	Modulation                      = 0x01
	BreathController                = 0x02
	FootController                  = 0x04
	PortamentoTime                  = 0x05
	DataEntry                       = 0x06
	MainVolume                      = 0x07
	Balance                         = 0x08
	Pan                             = 0x0a
	Expression                      = 0x0b
	EffectControl1                  = 0x0c
	EffectControl2                  = 0x0d
	GeneralPurposeController1       = 0x10
	GeneralPurposeController2       = 0x11
	GeneralPurposeController3       = 0x12
	GeneralPurposeController4       = 0x13
	BankSelectLSB                   = 0x20
	ModulationLSB                   = 0x21
	BreathControllerLSB             = 0x22
	FootControllerLSB               = 0x24
	PortamentoTimeLSB               = 0x25
	DataEntryLSB                    = 0x26
	MainVolumeLSB                   = 0x27
	BalanceLSB                      = 0x28
	PanLSB                          = 0x2a
	ExpressionLSB                   = 0x2b
	EffectControl1LSB               = 0x2c
	EffectControl2LSB               = 0x2d
	GeneralPurposeController1LSB    = 0x30
	GeneralPurposeController2LSB    = 0x31
	GeneralPurposeController3LSB    = 0x32
	GeneralPurposeController4LSB    = 0x33
	Hold1                           = 0x40
	PortamentoOnOff                 = 0x41
	Sostenuto                       = 0x42
	SoftPedal                       = 0x43
	LegatoFootswitch                = 0x44
	Hold2                           = 0x45
	SoundVariation                  = 0x46
	HarmonicIntensity               = 0x47
	ReleaseTime                     = 0x48
	AttackTime                      = 0x49
	Brightness                      = 0x4a
	DecayTime                       = 0x4b
	VibratoRate                     = 0x4c
	VibratoDepth                    = 0x4d
	VibratoDelay                    = 0x4e
	UndefinedSoundController        = 0x4f
	GeneralPurposeController5       = 0x50
	GeneralPurposeController6       = 0x51
	GeneralPurposeController7       = 0x52
	GeneralPurposeController8       = 0x53
	PortamentoControl               = 0x54
	ReverbSendLevel                 = 0x5b
	TremoloDepth                    = 0x5c
	ChorusSendLevel                 = 0x5d
	CelesteDepth                    = 0x5e
	PhaserDepth                     = 0x5f
	DataIncrement                   = 0x60
	DataDecrement                   = 0x61
	NonRegisteredParameterNumberLSB = 0x62
	NonRegisteredParameterNumberMSB = 0x63
	RegisteredParameterNumberLSB    = 0x64
	RegisteredParameterNumberMSB    = 0x65
)
