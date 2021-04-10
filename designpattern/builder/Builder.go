package builder

type Builder interface {
	SetCPU(cpu string) Builder
	SetMemory(memory string) Builder
	SetHardDisk(hardDisk string) Builder
	Build() *Computer
}

type ComputerBuilder struct {
	computer *Computer
}

func (c *ComputerBuilder) SetCPU(cpu string) Builder {
	if c.computer == nil {
		c.computer = new(Computer)
	}
	c.computer.SetCpu(cpu)
	return c
}

func (c *ComputerBuilder) SetMemory(memory string) Builder {
	if c.computer == nil {
		c.computer = new(Computer)
	}
	c.computer.SetMemory(memory)
	return c
}

func (c *ComputerBuilder) SetHardDisk(hardDisk string) Builder {
	if c.computer == nil {
		c.computer = new(Computer)
	}
	c.computer.SetHardDisk(hardDisk)
	return c
}

func (c *ComputerBuilder) Build() *Computer {
	return c.computer
}
