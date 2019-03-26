package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"runtime"
	"strconv"
)

func dealwithErr(err error) {
	if err != nil {
		fmt.Println(err)
		//os.Exit(-1)
	}
}

func GetHardwareData() {
	runtimeOS := runtime.GOOS
	// memory
	vmStat, err := mem.VirtualMemory()
	dealwithErr(err)

	// disk - start from "/" mount point for Linux
	// might have to change for Windows!!
	// don't have a Window to test this out, if detect OS == windows
	// then use "\" instead of "/"

	diskStat, err := disk.Usage("/")
	dealwithErr(err)

	// cpu - get CPU number of cores and speed
	cpuStat, err := cpu.Info()
	dealwithErr(err)
	percentage, err := cpu.Percent(0, true)
	dealwithErr(err)

	// host or machine kernel, uptime, platform Info
	hostStat, err := host.Info()
	dealwithErr(err)

	// get interfaces MAC/hardware address
	interfStat, err := net.Interfaces()
	dealwithErr(err)

	html := "OS : " + runtimeOS + "\n"
	html += "Total memory: " + strconv.FormatUint(vmStat.Total, 10) + " bytes \n"
	html += "Free memory: " + strconv.FormatUint(vmStat.Free, 10) + " bytes\n"
	html += "Percentage used memory: " + strconv.FormatFloat(vmStat.UsedPercent, 'f', 2, 64) + "%\n"

	// get disk serial number.... strange... not available from disk package at compile time
	// undefined: disk.GetDiskSerialNumber
	//serial := disk.GetDiskSerialNumber("/dev/sda")

	//html += "Disk serial number: " + serial + "\n"

	html += "Total disk space: " + strconv.FormatUint(diskStat.Total, 10) + " bytes \n"
	html += "Used disk space: " + strconv.FormatUint(diskStat.Used, 10) + " bytes\n"
	html += "Free disk space: " + strconv.FormatUint(diskStat.Free, 10) + " bytes\n"
	html += "Percentage disk space usage: " + strconv.FormatFloat(diskStat.UsedPercent, 'f', 2, 64) + "%\n"

	// since my machine has one CPU, I'll use the 0 index
	// if your machine has more than 1 CPU, use the correct index
	// to get the proper data
	html += "CPU index number: " + strconv.FormatInt(int64(cpuStat[0].CPU), 10) + "\n"
	html += "VendorID: " + cpuStat[0].VendorID + "\n"
	html += "Family: " + cpuStat[0].Family + "\n"
	html += "Number of cores: " + strconv.FormatInt(int64(cpuStat[0].Cores), 10) + "\n"
	html += "Model Name: " + cpuStat[0].ModelName + "\n"
	html += "Speed: " + strconv.FormatFloat(cpuStat[0].Mhz, 'f', 2, 64) + " MHz \n"

	for idx, cpupercent := range percentage {
		html += "Current CPU utilization: [" + strconv.Itoa(idx) + "] " + strconv.FormatFloat(cpupercent, 'f', 2, 64) + "%\n"
	}

	html += "Hostname: " + hostStat.Hostname + "\n"
	html += "Uptime: " + strconv.FormatUint(hostStat.Uptime, 10) + "\n"
	html += "Number of processes running: " + strconv.FormatUint(hostStat.Procs, 10) + "\n"

	// another way to get the operating system name
	// both darwin for Mac OSX, For Linux, can be ubuntu as platform
	// and linux for OS

	html += "OS: " + hostStat.OS + "\n"
	html += "Platform: " + hostStat.Platform + "\n"

	// the unique hardware id for this machine
	html += "Host ID(uuid): " + hostStat.HostID + "\n"

	for _, interf := range interfStat {
		html += "------------------------------------------------------\n"
		html += "Interface Name: " + interf.Name + "\n"

		if interf.HardwareAddr != "" {
			html += "Hardware(MAC) Address: " + interf.HardwareAddr + "\n"
		}

		for _, flag := range interf.Flags {
			html += "Interface behavior or flags: " + flag + "\n"
		}

		for _, addr := range interf.Addrs {
			html += "IPv6 or IPv4 addresses: " + addr.String() + "\n"
		}
	}

	fmt.Println(html)
}

func main() {
	GetHardwareData()
}
