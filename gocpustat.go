package gocpustat

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// CPUStat struct type for CPU Stat
type CPUStat struct {
	NumberOfLogicalCPU int
	TotalCore          int
}

// CPUInfo struct type for CPU Information
type CPUInfo struct {
	// Id processor ID
	ID int
	// CPUCores holds number of cores
	CPUCores int
}

// CPUInfoAll struct type of CPUInfo
type CPUInfoAll struct {
	CPUInfoAll []CPUInfo
}

// GetCPUInfo will return infformation about
// all processors installed on a Linux machine
func GetCPUInfo(path string) ([]CPUInfo, error) {
	b, err := ioutil.ReadFile(path)
	CPUInfoAll := []CPUInfo{}
	if err != nil {
		return CPUInfoAll, err
	}
	content := string(b)
	lines := strings.Split(content, "\n\n")

	for i := 0; i < len(lines)-1; i++ {
		line := strings.Split(lines[i], "\n")
		CPUInfo := CPUInfo{}
		for j := 0; j < len(line); j++ {
			values := strings.Split(line[j], ":")

			switch key := strings.TrimSpace(values[0]); key {
			case "processor":
				val, _ := strconv.Atoi(strings.TrimSpace(values[1]))
				CPUInfo.ID = val
			case "cpu cores":
				val, _ := strconv.Atoi(strings.TrimSpace(values[1]))
				CPUInfo.CPUCores = val
			}

		}
		CPUInfoAll = append(CPUInfoAll, CPUInfo)
	}
	return CPUInfoAll, nil
}

// GetCPUStat returns number of Processors and number of core per Processor
// installed in the given Linux machine
func GetCPUStat() CPUStat {
	cpuInfo, _ := GetCPUInfo("/proc/cpuinfo")
	cpuStat := CPUStat{}
	cpuStat.NumberOfLogicalCPU = len(cpuInfo)
	if len(cpuInfo) > 0 {
		cpuStat.TotalCore = cpuInfo[0].CPUCores
		cpuStat.ThreadPerCore = cpuStat.NumberOfLogicalCPU / cpuStat.TotalCore
	}

	return cpuStat
}
