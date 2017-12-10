package event

import "github.com/moutend/go-midi/deltatime"

// Event represents any MIDI events, including meta and system exclusive.
type Event interface {
	DeltaTime() *deltatime.DeltaTime
	Serialize() []byte

	SetRunningStatus(bool)
	RunningStatus() bool
}
