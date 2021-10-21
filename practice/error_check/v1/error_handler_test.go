package error_check_v0

import (
	"designer_go/practice/error_check"
	"encoding/binary"
	"io"
)

func parse(r io.Reader) (*error_check.Point, error) {
	var p error_check.Point
	var err error

	read := func(data interface{}) {
		if err != nil {
			return
		}
		err = binary.Read(r, binary.BigEndian, data)
	}

	read(&p.Longitude)
	read(&p.Latitude)
	read(&p.Distance)
	read(&p.ElevationGain)
	read(&p.ElevationLoss)

	if err != nil {
		return &p, err
	}

	return &p, nil
}

