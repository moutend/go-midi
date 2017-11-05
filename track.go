package midi

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type Track struct {
	Events []*Event
}

func parseTrack(stream []byte, start int64) (*Track, error) {
	var chunkId ChunkId
	var chunkSize uint32

	track := &Track{}
	data := bytes.NewReader(stream)

	binary.Read(io.NewSectionReader(data, start, 4), binary.BigEndian, &chunkId)
	if chunkId != MTrk {
		return nil, fmt.Errorf("midi: invalid track ID %v", chunkId)
	}

	binary.Read(io.NewSectionReader(data, 4, 4), binary.BigEndian, &chunkSize)

	return track, nil
}

func parseTracks(data []byte, size uint16) ([]*Track, error) {
	tracks := make([]*Track, size)
	track, err := parseTrack(data, int64(14))
	if err != nil {
		return nil, err
	}
	tracks[0] = track

	return tracks, nil
}
