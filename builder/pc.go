package builder

import (
	"errors"
)

type PC struct {
	CPU     string
	Memory  byte
	Storage byte
	Wifi    bool
	OS      string
}

type PcBuilder struct {
	pc *PC
}

func NewPcBuilder() *PcBuilder {
	return &PcBuilder{
		pc: &PC{},
	}
}

func (pb *PcBuilder) SetCPU(cpuName string) *PcBuilder {
	pb.pc.CPU = cpuName
	return pb
}

func (pb *PcBuilder) SetMemory(memorySize byte) *PcBuilder {
	pb.pc.Memory = memorySize
	return pb
}

func (pb *PcBuilder) SetStorage(storageSize byte) *PcBuilder {
	pb.pc.Storage = storageSize
	return pb
}

func (pb *PcBuilder) SetWifi(isWifi bool) *PcBuilder {
	pb.pc.Wifi = isWifi
	return pb
}

func (pb *PcBuilder) SetOS(osType string) *PcBuilder {
	pb.pc.OS = osType
	return pb
}

func (pb *PcBuilder) Build() (*PC, error) {
	// validation
	if pb.pc.CPU == "" {
		return nil, errors.New("CPU is required.")
	}
	if pb.pc.Memory == 0 {
		return nil, errors.New("Memory is required.")
	}

	pc := &PC{
		CPU:     pb.pc.CPU,
		Memory:  pb.pc.Memory,
		Storage: pb.pc.Storage,
		Wifi:    pb.pc.Wifi,
		OS:      pb.pc.OS,
	}

	// set defacult value
	if pc.Storage == 0 {
		pc.Storage = 64
	}
	if pc.OS == "" {
		pc.OS = "Windows 11"
	}

	return pc, nil
}
