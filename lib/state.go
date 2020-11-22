package lib

import (
	"encoding/json"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"time"
)

type (
	State struct {
		Mem Memory
		Swap Swap
		Cpu []CPU
		Host Host
		Disk Disk
		Nets []Net
		Partitions []Partition
		CollectTimestamp int64
		CollectTime string
		Token string
	}
	Memory struct {
		//{"total":17179869184,"available":5309943808,"used":11869925376,"usedPercent":69.09205913543701,"free":407195648,"active":5144727552,"inactive":4902748160,"wired":3687264256}
		Total uint64
		Available uint64
		Used uint64
		UsedPercent float64
		Free uint64
		Active uint64
		Inactive uint64
		Wired uint64
	}
	Swap struct {
		//{"total":1073741824,"used":524288,"free":1073217536,"usedPercent":0.048828125}
		Total uint64
		Used uint64
		UsedPercent float64
		Free uint64
	}
	CPU struct {
		//[]
		//{"cpu":0,"vendorId":"GenuineIntel","family":"6","model":"158","stepping":10,"physicalId":"","coreId":"","cores":6,"modelName":"Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz","mhz":2600,"cacheSize":256,"flags":["fpu","vme","de","pse","tsc","msr","pae","mce","cx8","apic","sep","mtrr","pge","mca","cmov","pat","pse36","clfsh","ds","acpi","mmx","fxsr","sse","sse2","ss","htt","tm","pbe","sse3","pclmulqdq","dtes64","mon","dscpl","vmx","smx","est","tm2","ssse3","fma","cx16","tpr","pdcm","sse4.1","sse4.2","x2apic","movbe","popcnt","aes","pcid","xsave","osxsave","seglim64","tsctmr","avx1.0","rdrand","f16c","rdwrfsgs","tsc_thread_offset","sgx","bmi1","hle","avx2","smep","bmi2","erms","invpcid","rtm","fpu_csds","mpx","rdseed","adx","smap","clfsopt","ipt","sgxlc","mdclear","tsxfa","ibrs","stibp","l1df","ssbd","syscall","xd","1gbpage","em64t","lahf","lzcnt","prefetchw","rdtscp","tsci"],"microcode":""}
		Id        int32
		VendorID   string
		Family     string
		Model      string
		Stepping   int32
		PhysicalID string
		CoreID     string
		Cores      int32
		ModelName  string
		Mhz        float64
		CacheSize  int32
		Flags      []string
		Microcode  string
		Percent float64
		Percents []float64
	}
	Host struct {
		//{"hostname":"hq-mac00451","uptime":20368,"bootTime":1605575435,"procs":618,"os":"darwin","platform":"darwin","platformFamily":"Standalone Workstation","platformVersion":"11.0.1","kernelVersion":"20.1.0","kernelArch":"x86_64","virtualizationSystem":"","virtualizationRole":"","hostid":"84c6dc45-6b02-335f-9439-5d2a9bc385a4"}
		Hostname             string
		Uptime               uint64
		BootTime             uint64
		Procs                uint64
		OS                   string
		Platform             string
		PlatformFamily       string
		PlatformVersion      string
		KernelVersion        string
		KernelArch           string
		VirtualizationSystem string
		VirtualizationRole   string
		HostID               string
	}
	Disk struct {
		Path              string
		Fstype            string
		Total             uint64
		Free              uint64
		Used              uint64
		UsedPercent       float64
		InodesTotal       uint64
		InodesUsed        uint64
		InodesFree        uint64
		InodesUsedPercent float64
	}
	Net struct {
		//[]
		//{"name":"all","bytesSent":2208708953,"bytesRecv":4096856793,"packetsSent":3184766,"packetsRecv":4421366}
		Name        string
		BytesSent   uint64
		BytesRecv   uint64
		PacketsSent uint64
		PacketsRecv uint64
		Errin       uint64
		Errout      uint64
		Dropin      uint64
		Dropout     uint64
		Fifoin      uint64
		Fifoout     uint64
	}
	Process struct {
	}
	Partition struct {
		//[]
		//{"device":"/dev/disk1s5s1","mountpoint":"/","fstype":"apfs","opts":"ro,journaled,multilabel"}
		Device     string
		Mountpoint string
		Fstype     string
		Opts       string
	}
)

func GetState() string {
	bs, _ := json.Marshal(collet())
	return string(bs)
}

func collet() State {
	state := State{}

	state.CollectTime = time.Now().Format("2006-01-02 15:04:05")
	state.CollectTimestamp = time.Now().Unix()

	_mem  := Memory{}
	_swap  := Swap{}
	_cpus := make([]CPU, 0)
	_host:= Host{}
	_disk := Disk{}
	_nets := make([]Net, 0)
	_parts := make([]Partition, 0)


	vm, _:= mem.VirtualMemory()
	_mem.Active = vm.Active
	_mem.Available = vm.Available
	_mem.Free = vm.Free
	_mem.Inactive = vm.Inactive
	_mem.Total = vm.Total
	_mem.Used = vm.Used
	_mem.Wired = vm.Wired
	_mem.UsedPercent = vm.UsedPercent
	// 内存
	state.Mem = _mem
	sm, _ := mem.SwapMemory()
	_swap.UsedPercent = sm.UsedPercent
	_swap.Used = sm.Used
	_swap.Total = sm.Total
	_swap.Free = sm.Free
	// Swap
	state.Swap = _swap
	c, _ := cpu.Info()
	for _,_c := range c {
		_cpu := CPU{}
		_cpu.CacheSize = _c.CacheSize
		_cpu.Cores = _c.Cores
		_cpu.Id = _c.CPU
		_cpu.Stepping = _c.Stepping
		_cpu.CoreID = _c.CoreID
		_cpu.Family = _c.Family
		_cpu.Flags = _c.Flags
		_cpu.Mhz = _c.Mhz
		_cpu.Microcode = _c.Microcode
		_cpu.Model = _c.Model
		_cpu.ModelName = _c.ModelName
		_cpu.PhysicalID = _c.PhysicalID
		_cpu.VendorID = _c.VendorID


		_p_each, _ := cpu.Percent(time.Second, true)
		_p_all, _ := cpu.Percent(time.Second, false)

		_cpu.Percent = _p_all[0]
		_cpu.Percents = _p_each

		_cpus = append(_cpus, _cpu)

	}
	//  CPUs
	state.Cpu = _cpus
	n, _ := host.Info()
	_host.BootTime =  n.BootTime
	_host.Procs =  n.Procs
	_host.Uptime =  n.Uptime
	_host.HostID =  n.HostID
	_host.Hostname =  n.Hostname
	_host.KernelArch =  n.KernelArch
	_host.KernelVersion =  n.KernelVersion
	_host.OS =  n.OS
	_host.Platform =  n.Platform
	_host.PlatformFamily =  n.PlatformFamily
	_host.PlatformVersion =  n.PlatformVersion
	_host.VirtualizationRole =  n.VirtualizationRole
	_host.VirtualizationSystem =  n.VirtualizationSystem
	// Host
	state.Host =  _host
	d, _ := disk.Usage("/")
	_disk.Free = d.Free
	_disk.Total = d.Total
	_disk.Used = d.Used
	_disk.InodesFree = d.InodesFree
	_disk.InodesTotal = d.InodesTotal
	_disk.UsedPercent = d.UsedPercent
	_disk.Fstype = d.Fstype
	_disk.InodesUsedPercent = d.InodesUsedPercent
	_disk.Path = d.Path
	_disk.InodesUsed = d.InodesUsed
	// Disk
	state.Disk = _disk

	net_all, _ := net.IOCounters(false)
	net_each, _ := net.IOCounters(true)
	_net_all  := Net{}
	_net_all.Name = net_all[0].Name
	_net_all.PacketsSent = net_all[0].PacketsSent
	_net_all.PacketsRecv = net_all[0].PacketsRecv
	_net_all.BytesRecv = net_all[0].BytesRecv
	_net_all.BytesSent = net_all[0].BytesSent
	_net_all.Dropin = net_all[0].Dropin
	_net_all.Dropout = net_all[0].Dropout
	_net_all.Errin = net_all[0].Errin
	_net_all.Errout = net_all[0].Errout
	_net_all.Fifoin = net_all[0].Fifoin
	_net_all.Fifoout = net_all[0].Fifoout
	_nets = append(_nets, _net_all)
	for _, nvone := range net_each {
		_net_one  := Net{}
		_net_one.Name = nvone.Name
		_net_one.PacketsSent = nvone.PacketsSent
		_net_one.PacketsRecv = nvone.PacketsRecv
		_net_one.BytesRecv = nvone.BytesRecv
		_net_one.BytesSent = nvone.BytesSent
		_net_one.Dropin = nvone.Dropin
		_net_one.Dropout = nvone.Dropout
		_net_one.Errin = nvone.Errin
		_net_one.Errout = nvone.Errout
		_net_one.Fifoin = nvone.Fifoin
		_net_one.Fifoout = nvone.Fifoout
		_nets = append(_nets, _net_one)
	}
	// Nets
	state.Nets = _nets
	ps, _ := disk.Partitions(true)
	for _, p := range ps {
		_part := Partition{}
		_part.Fstype = p.Fstype
		_part.Device = p.Device
		_part.Mountpoint = p.Mountpoint
		_part.Opts = p.Opts
		_parts = append(_parts, _part)
	}
	// Partitions
	state.Partitions = _parts

	return state

	//fmt.Println(cpu.Percent(time.Second, true))
	//fmt.Println(cpu.Percent(time.Second, false))
}
