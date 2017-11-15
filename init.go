package midi

var logger *midiLogger

func init() {
	logger = newMIDILogger()
}
