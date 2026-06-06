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

type pcBuilder struct {
	pc *PC
}

func NewPcBuilder() *pcBuilder {
	return &pcBuilder{
		pc: &PC{},
	}
}

func (pb *pcBuilder) SetCPU(cpuName string) *pcBuilder {
	pb.pc.CPU = cpuName
	return pb
}

func (pb *pcBuilder) SetMemory(memorySize byte) *pcBuilder {
	pb.pc.Memory = memorySize
	return pb
}

func (pb *pcBuilder) SetStorage(storageSize byte) *pcBuilder {
	pb.pc.Storage = storageSize
	return pb
}

func (pb *pcBuilder) SetWifi(isWifi bool) *pcBuilder {
	pb.pc.Wifi = isWifi
	return pb
}

func (pb *pcBuilder) SetOS(osType string) *pcBuilder {
	pb.pc.OS = osType
	return pb
}

func (pb *pcBuilder) Build() (*PC, error) {
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
