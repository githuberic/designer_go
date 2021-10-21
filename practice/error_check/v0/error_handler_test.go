package error_check_v0

import (
	"designer_go/practice/error_check"
	"encoding/binary"
	"io"
)

func parse(r io.Reader) (*error_check.Point, error) {
	var p error_check.Point

	if err := binary.Read(r, binary.BigEndian, &p.Longitude); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.Latitude); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.Distance); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.ElevationGain); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.ElevationLoss); err != nil {
		return nil, err
	}
	return nil, nil
}
