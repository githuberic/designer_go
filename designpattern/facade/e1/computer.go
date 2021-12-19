package e1

import "fmt"

// CPU
type CPU struct {
}

func (p *CPU) start() {
	fmt.Println("Start CPU")
}

// 内存
type Memory struct {
}

func (p *Memory) start() {
	fmt.Println("Start memory management")
}

// 硬盘
type Disk struct {
}

func (p *Disk) start() {
	fmt.Println("Start disk")
}

// 开机键
type StartBtn struct {
}

func (StartBtn) start() {
	cpu := &CPU{}
	cpu.start()

	memory := &Memory{}
	memory.start()

	disk := &Disk{}
	disk.start()
}
