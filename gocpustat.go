package gocpustat

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// CPUStat struct type for CPU Stat
type CPUStat struct {
	NumberOfCPU int
	CorePerCPU  int
	TotalCore   int
}

// CPUInfo struct type for CPU Information
type CPUInfo struct {
	// Id processor ID
	ID int `json:"processor"`
	// CPUCores holds number of cores
	CPUCores int `json:"CPU cores"`
}

// CPUInfoAll struct type of CPUInfo
type CPUInfoAll struct {
	CPUInfoAll []CPUInfo `json:"CPUInfoAll"`
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
			//fmt.Println(values[0])

			switch key := strings.TrimSpace(values[0]); key {
			case "processor":
				val, _ := strconv.Atoi(strings.TrimSpace(values[1]))
				CPUInfo.ID = val
			case "cpu cores":
				val, _ := strconv.Atoi(strings.TrimSpace(values[1]))
				CPUInfo.CPUCores = val
			}

		}
		fmt.Println(CPUInfo.CPUCores)
		CPUInfoAll = append(CPUInfoAll, CPUInfo)
	}
	return CPUInfoAll, nil
}

// GetCPUStat returns number of Processors and number of core per Processor
// installed in the given Linux machine
func GetCPUStat() CPUStat {
	cpuInfo, _ := GetCPUInfo("test_CPUinfo")
	cpuStat := CPUStat{}
	coreCount := 0
	for i := 0; i < len(cpuInfo); i++ {
		coreCount += cpuInfo[i].CPUCores
		// change logic later
		cpuStat.CorePerCPU = cpuInfo[i].CPUCores
	}
	cpuStat.NumberOfCPU = len(cpuInfo)
	cpuStat.TotalCore = coreCount
	return cpuStat
}
