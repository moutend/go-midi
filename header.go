package midi

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

// Header represents MIDI header. It contains format type, number of tracks and time division.
type Header struct {
	formatType   uint16
	tracks       uint16
	timeDivision *TimeDivision
}

// FormatType returns MIDI format type.
func (h *Header) FormatType() int {
	return int(h.formatType)
}

// Tracks returns number of tracks.
func (h *Header) Tracks() int {
	return int(h.tracks)
}

// TimeDivision returns time division.
func (h *Header) TimeDivision() *TimeDivision {
	if h.timeDivision == nil {
		h.timeDivision = &TimeDivision{}
	}
	return h.timeDivision
}

// String returns string representation of the Header.
// The returned string is meant for debugging.
func (h *Header) String() string {
	return fmt.Sprintf("FormatType: %d Tracks: %d TimeDivision: %v", h.formatType, h.tracks, h.TimeDivision().String())
}

func (h *Header) serialize() []byte {
	data := bytes.NewBuffer([]byte{})

	binary.Write(data, binary.BigEndian, []byte("MThd"))
	binary.Write(data, binary.BigEndian, uint32(6))
	binary.Write(data, binary.BigEndian, h.formatType)
	binary.Write(data, binary.BigEndian, h.tracks)
	binary.Write(data, binary.BigEndian, h.timeDivision.value)

	return data.Bytes()
}

func parseHeader(stream []byte) (*Header, error) {
	const MThd uint32 = 0x4d546864
	var chunkId uint32
	var timeDivision uint16

	header := &Header{}
	data := bytes.NewReader(stream)
	binary.Read(io.NewSectionReader(data, 0, 4), binary.BigEndian, &chunkId)
	if chunkId != MThd {
		return nil, fmt.Errorf("midi: invalid chunk ID for header: %x", chunkId)
	}

	binary.Read(io.NewSectionReader(data, 8, 2), binary.BigEndian, &header.formatType)
	binary.Read(io.NewSectionReader(data, 10, 2), binary.BigEndian, &header.tracks)
	binary.Read(io.NewSectionReader(data, 12, 2), binary.BigEndian, &timeDivision)
	header.TimeDivision().Set(timeDivision)

	return header, nil
}
