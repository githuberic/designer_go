package builder

type Director struct {
	builder Builder
}

func (d *Director) Create(cpu string, memory string, hardDisk string) *Computer {
	if d.builder == nil {
		computerBuilder := new(ComputerBuilder)
		d.builder = computerBuilder
	}
	return d.builder.SetCPU(cpu).SetMemory(memory).SetHardDisk(hardDisk).Build()
}
