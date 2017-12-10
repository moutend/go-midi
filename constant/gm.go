//go:generate stringer -type=GM -output=gm_string.go

package constant

// GM represents GM tones.
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
