package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)

	cpuCount, _ := cpu.Counts(false)
	fmt.Println(cpuCount)

	duration := time.Duration(10) * time.Microsecond
	duration = 0
	cpuPercentage, _ := cpu.Percent(duration, true)
	fmt.Println(cpuPercentage)
	// cpuPercentage, _ = cpu.Percent(duration, true)
	// fmt.Println(cpuPercentage)
}
