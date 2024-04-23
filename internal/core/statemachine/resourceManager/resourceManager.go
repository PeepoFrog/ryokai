package resourcemanager

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/KiraCore/ryokai/pkg/ryokaicommon/types"
	osutils "github.com/KiraCore/ryokai/pkg/ryokaicommon/utils/os"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

const SysRes string = "/tmp/testConfig.toml"

func GetTotalSystemResources() (*types.SystemResources, error) {
	systemResources := &types.SystemResources{}

	vMemStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	systemResources.RAM.Set(vMemStat.Total)

	logicCpuCores, err := cpu.Counts(true)
	if err != nil {
		return nil, err
	}
	systemResources.Cpu.Set(uint(logicCpuCores))

	disk, err := getCurrentDisk()
	if err != nil {
		return nil, err
	}
	systemResources.Disk.Set(disk.Total)

	return systemResources, nil
}

func GetAvailableResources() (*types.SystemResources, error) {
	systemResources := &types.SystemResources{}

	vMemStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	systemResources.RAM.Set(vMemStat.Available)

	logicCpuCores, err := cpu.Counts(true)
	if err != nil {
		return nil, err
	}
	systemResources.Cpu.Set(uint(logicCpuCores))

	disk, err := getCurrentDisk()
	if err != nil {
		return nil, err
	}
	systemResources.Disk.Set(disk.Free)

	return systemResources, nil
}

func WriteSystemRecourses(systemResourcesToWrite *types.SystemResources, configFilePath string) error {
	isDir := osutils.IsDir(SysRes)
	if isDir {
		return fmt.Errorf("error while writing %v file, file is a dir", SysRes)
	}

	cfgDir := filepath.Dir(configFilePath)
	exist := osutils.PathExists(cfgDir)
	if !exist {
		err := os.MkdirAll(cfgDir, 0o755)
		if err != nil {
			return err
		}
	}

	f, err := os.Create(configFilePath)
	if err != nil {
		return err
	}

	encoder := toml.NewEncoder(f)
	err = encoder.Encode(systemResourcesToWrite)
	if err != nil {
		return err
	}

	return nil
}

func ReadSystemRecourses(configFilePath string) (*types.SystemResources, error) {
	exist := osutils.PathExists(configFilePath)
	isDir := osutils.IsDir(configFilePath)
	if !exist && isDir {
		return nil, fmt.Errorf("unable to read system resources state file, exist=%v, isDir=%v", exist, isDir)
	}

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	sysRes := &types.SystemResources{}
	err = toml.Unmarshal(data, sysRes)
	if err != nil {
		return nil, err
	}
	return sysRes, nil
}

func getCurrentDisk() (*disk.UsageStat, error) {
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
			// log.Printf("Current directory %s is on partition %s mounted at %s\n", cwd, p.Device, p.Mountpoint)
			usage, err := disk.Usage(p.Mountpoint)
			if err != nil {
				fmt.Println("Error getting partition usage:", err)
				return nil, err
			}

			// log.Printf("Disk Usage of Partition %s: Total: %d, Free: %d, Used: %d\n", p.Device, usage.Total, usage.Free, usage.Used)
			return usage, nil
		}
	}
	return nil, fmt.Errorf("unable to get current disk, disk not found")
}
