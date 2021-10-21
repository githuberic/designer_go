package error_check_v0

import (
	"designer_go/practice/error_check"
	"encoding/binary"
	"io"
)

func parse(input io.Reader) (*error_check.Point, error) {
	var p error_check.Point
	r := Reader{r: input}

	r.read(&p.Longitude)
	r.read(&p.Latitude)
	r.read(&p.Distance)
	r.read(&p.ElevationGain)
	r.read(&p.ElevationLoss)

	if r.err != nil {
		return &p, r.err
	}

	return &p, nil
}

type Reader struct {
	r   io.Reader
	err error
}

func (r *Reader) read(data interface{}) {
	if r.err == nil {
		r.err = binary.Read(r.r, binary.BigEndian, data)
	}
}
