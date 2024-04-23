package types

type SystemResources struct {
	RAM  RAM  // get GB 1024 / 1024 / 1024
	Cpu  CPU  // logic cores
	Disk Disk // diskSpace, get GB 1024 / 1024 / 1024
}

type (
	RAM  uint64
	Disk uint64
	CPU  uint
)

func (r *RAM) GetGB() float64 {
	return float64(*r) / 1024 / 1024 / 1024
}

func (d *Disk) GetGB() float64 {
	return float64(*d) / 1024 / 1024 / 1024
}

func (r *RAM) Set(bytes uint64) {
	*r = RAM(bytes)
}

func (d *Disk) Set(bytes uint64) {
	*d = Disk(bytes)
}

func (c *CPU) Set(cores uint) {
	*c = CPU(cores)
}
