package meson_msg

type MachineStateBaseMsg struct {
	OS           string  //`json:"os"`
	CPU          string  //`json:"cpu"` // cpu model name
	MemTotal     int64   //`json:"mem_total"` // uint: byte
	MemAvailable int64   //`json:"mem_avail"`
	Version      string  //`json:"version"`
	CpuUsage     float64 //`json:"cpu_usage"`
}

type TerminalStatesMsg struct {
	Port             int        //`json:"port"`
	NetInMbs         [5]float64 //`json:"net_in_mbs"`
	NetOutMbs        [5]float64 //`json:"net_out_mbs"`
	CdnDiskTotal     int64      //`json:"cdn_disk_total"`
	CdnDiskUsed      int64      //`json:"cdn_disk_avail"`
	MachineSetupTime string     //`json:"machine_setup_time"`
	SequenceId       int        //`json:"sequence_id"`
	MachineStateBaseMsg
}
