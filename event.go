package midi

// Event represents any MIDI events, including meta and system exclusive.
type Event interface {
	DeltaTime() *DeltaTime
	Serialize() []byte

	SetRunningStatus(bool)
	RunningStatus() bool
}
