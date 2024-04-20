package types

type SystemResources struct {
	RAM  uint64
	Cpu  uint16 // logic cores
	Disk uint   // diskSpace
}
