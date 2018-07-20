package process

import (
	"github.com/dustin/go-humanize"
	cf "github.com/euclid1990/gomworker/configs"
	"github.com/struCoder/pidusage"
	"time"
)

// Define process status
const (
	PROCESS_STATUS_STARTED       = "started"
	PROCESS_STATUS_RUNNING       = "running"
	PROCESS_STATUS_STOPPED       = "stopped"
	PROCESS_STATUS_AKSED_TO_STOP = "asked_to_stop"
)

// ProcessStatus is a wrapper with the process current status.
type ProcessStatus struct {
	Status    string
	Restarts  int
	StartTime int64
	Uptime    string
	Sys       *pidusage.SysInfo
}

// SetStatus will set the process string status.
func (pS *ProcessStatus) SetStatus(status string) {
	pS.Status = status
}

// AddRestart will add one restart to the process status.
func (pS *ProcessStatus) AddRestart() {
	pS.Restarts++
}

// InitUptime will record process start time
func (pS *ProcessStatus) InitUptime() {
	pS.StartTime = time.Now().Unix()
}

// SetUptime will figure out process uptime
func (pS *ProcessStatus) SetUptime() {
	pS.Uptime = humanize.Time(time.Unix(pS.StartTime))
}

// ResetUptime will reset uptime
func (pS *ProcessStatus) ResetUptime() {
	pS.Uptime = "0s"
}

// SetSysInfo will get current process cpu and memory usage
func (pS *ProcessStatus) SetSysInfo(pid int) {
	var err error
	pS.Sys, err = pidusage.GetStat(pid)
	if err != nil {
		Logf(cf.LOG_WARNING, "", err)
	}
}
