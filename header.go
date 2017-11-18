package midi

import "fmt"

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

func (h *Header) SetTracks(tracks uint16) {
	h.tracks = tracks
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

// Serialize serializes MIDI header.
func (h *Header) Serialize() []byte {
	bs := []byte("MThd")
	bs = append(bs, 0x00, 0x00, 0x00, 0x06)
	bs = append(bs, 0x00, byte(h.formatType&0xff))
	bs = append(bs, byte((h.tracks>>8)&0xff))
	bs = append(bs, byte(h.tracks&0xff))
	bs = append(bs, h.TimeDivision().Serialize()...)

	return bs
}

// parseHeader parses stream begins with MThd.
func parseHeader(stream []byte) (*Header, error) {
	var start int

	logger.Println("start parsing MThd")

	if string(stream[:4]) != "MThd" {
		return nil, fmt.Errorf("midi: invalid chunk ID %v", stream[:4])
	}

	start += 4
	logger.parsedBytes += 4
	logger.Println("parsing MThd completed")

	start += 4 // skip read header size
	logger.parsedBytes += 4
	logger.Println("skip parsing size of header chunk")

	logger.Println("start parsing format type")

	formatType := uint16(stream[start+1])

	start += 2
	logger.parsedBytes += 2
	logger.Printf("parsing format type completed (formatType=%v)", formatType)

	logger.Println("start parsing number of tracks")

	numberOfTracks := uint16(stream[start])
	numberOfTracks = numberOfTracks << 8
	numberOfTracks += uint16(stream[start+1])

	start += 2
	logger.parsedBytes += 2
	logger.Printf("parsing number of tracks completed (%v)", numberOfTracks)

	logger.Println("start parsing time division")

	timeDivision := uint16(stream[start])
	timeDivision = timeDivision << 8
	timeDivision += uint16(stream[start+1])

	start += 2
	logger.parsedBytes += 2
	logger.Printf("parsing time division completed (%v)", timeDivision)

	header := &Header{
		formatType: formatType,
		tracks:     numberOfTracks,
	}
	header.TimeDivision().Set(timeDivision)

	return header, nil
}
