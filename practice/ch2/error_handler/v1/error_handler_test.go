package error_handler_test

import (
	"encoding/binary"
	"io"
)

func parse(r io.Reader) (*Point, error) {
	var p Point

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
}


// 函数式编程的方式
func parse(r io.Reader) (*Point, error) {
	var p Point
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


type Reader struct {
	r   io.Reader
	err error
}
func (r *Reader) read(data interface{}) {
	if r.err == nil {
		r.err = binary.Read(r.r, binary.BigEndian, data)
	}
}
func parse(input io.Reader) (*Point, error) {
	var p Point
	r := Reader{r: input}

	r.read(&p.Longitude)
	r.read(&p.Latitude)
	r.read(&p.Distance)
	r.read(&p.ElevationGain)
	r.read(&p.ElevationLoss)

	if r.err != nil {
		return nil, r.err
	}

	return &p, nil
}