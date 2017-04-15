package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	fmt.Println("*********************************************")
	cpuCount, _ := cpu.Counts(false)
	fmt.Println("CPU counts", cpuCount)

	duration := time.Duration(10) * time.Microsecond
	cpuPercentage, _ := cpu.Percent(duration, true)
	fmt.Println("CPU usage percentage", cpuPercentage)
	// cpuPercentage, _ = cpu.Percent(duration, true)
	// fmt.Println(cpuPercentage)

	times, _ := cpu.Times(false)
	fmt.Println("CPU Times", times)

	// info, _ := cpu.Info()
	// fmt.Println("CPU Info", info)

	// testCPUPercent(true)

}

func testCPUPercent(percpu bool) {
	numcpu := runtime.NumCPU()
	testCount := 1

	if runtime.GOOS != "windows" {
		testCount = 1
		v, err := cpu.Percent(time.Millisecond, percpu)
		if err != nil {
			fmt.Errorf("error %v", err)
		}
		// Skip CircleCI which CPU num is different
		if os.Getenv("CIRCLECI") != "true" {
			fmt.Println("in os.getenv")
			if (percpu && len(v) != numcpu) || (!percpu && len(v) != 1) {
				fmt.Errorf("wrong number of entries from CPUPercent: %v", v)
			}
		}
	}
	for i := 0; i < testCount; i++ {
		fmt.Println("in loop")
		duration := time.Duration(10) * time.Microsecond
		v, err := cpu.Percent(duration, percpu)
		if err != nil {
			fmt.Errorf("error %v", err)
		}
		for _, percent := range v {
			fmt.Println(percent)
			// Check for slightly greater then 100% to account for any rounding issues.
			if percent < 0.0 || percent > 100.0001*float64(numcpu) {
				fmt.Errorf("CPUPercent value is invalid: %f", percent)
			}
		}
	}
}
