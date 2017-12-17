//go:generate stringer -type=Note -output=note_string.go

package constant

import (
	"strconv"
	"strings"
)

// Note represents a note. The note number 60 corresponds to C3.
type Note uint8

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

var noteMap map[string]int = map[string]int{
	"c":  0,
	"c#": 1,
	"db": 1,
	"d":  2,
	"d#": 3,
	"eb": 3,
	"e":  4,
	"f":  5,
	"f#": 6,
	"gb": 6,
	"g":  7,
	"g#": 8,
	"ab": 8,
	"a":  9,
	"a#": 10,
	"bb": 10,
	"b":  11,
}

func ParseNote(s string) (Note, error) {
	s = strings.ToLower(s)
	i, err := strconv.Atoi(s)
	if err == nil {
		return Note(i), nil
	}
	octaveStr := s[len(s)-1 : len(s)-0]
	octave, err := strconv.Atoi(octaveStr)
	if err != nil {
		return Note(0), err
	}
	signStr := s[len(s)-2 : len(s)-1]
	if signStr == "-" {
		octave *= -1
	}
	noteStr := s[0:2]
	if noteStr[1:2] != "b" {
		noteStr = noteStr[0:1]
	}
	note := noteMap[noteStr] + (octave+2)*12
	return Note(uint8(note)), nil
}
