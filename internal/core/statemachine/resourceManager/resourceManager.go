package resourcemanager

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/KiraCore/ryokai/pkg/ryokaicommon/types"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func GetTotalSystemResources() (*types.SystemResources, error) {
	// Get virtual memory stats
	systemResources := &types.SystemResources{}

	vMemStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	systemResources.RAM = vMemStat.VMallocTotal
	logicCpuCores, err := cpu.Counts(true)
	if err != nil {
		return nil, err
	}
	systemResources.Cpu = uint16(logicCpuCores)

	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	cwd, err = filepath.Abs(cwd)
	if err != nil {
		return nil, err
	}

	partitions, err := disk.Partitions(true)
	if err != nil {
		return nil, err
	}

	for _, p := range partitions {
		if len(cwd) >= len(p.Mountpoint) && cwd[:len(p.Mountpoint)] == p.Mountpoint {
			log.Printf("Current directory %s is on partition %s mounted at %s\n", cwd, p.Device, p.Mountpoint)
			usage, err := disk.Usage(p.Mountpoint)
			if err != nil {
				fmt.Println("Error getting partition usage:", err)
				return nil, err
			}
			systemResources.Disk = uint(usage.Free)
			log.Printf("Disk Usage of Partition %s: Total: %d, Free: %d, Used: %d\n", p.Device, usage.Total, usage.Free, usage.Used)
			break
		}
	}

	return systemResources, nil
}
