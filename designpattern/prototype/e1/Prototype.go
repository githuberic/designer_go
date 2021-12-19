package e1

import (
	"bytes"
	"encoding/gob"
)

//速度速值
type Speed int

type FanSpeed struct {
	Speed Speed
}

type Money struct {
	Length float64
}

type Memory struct {
	Count      int
	MemorySize []int
}

//电脑信息
type Computer struct {
	SystemName string              //系统名字
	UseNumber  int                 //使用次数
	Memory     Memory              //存储
	Fan        map[string]FanSpeed //风扇
	Money      Money               //售价
}

func (c *Computer) Clone() *Computer {
	copy := *c
	return &copy
}

func (c *Computer) BackUp() *Computer {
	pc := new(Computer)

	if err := deepCopy(pc, c); err != nil {
		panic(err.Error())
	}
	return pc
}

func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer

	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}
