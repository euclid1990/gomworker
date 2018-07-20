package process

import (
	"errors"
	"ioutil"
	"os"
	"strconv"
	"syscall"
)

// Process is a os.Process wrapper with Status and more info that will be used on Master to maintain the process health.
type Process struct {
	Name      string
	Cmd       string
	Args      []string
	Path      string
	PidFile   string
	OutFile   string
	ErrFile   string
	KeepAlive bool
	Pid       int
	Status    *ProcessStatus
	process   *os.Process
}

// Start will execute the command Cmd that should run the process.
func (p *Process) Start() error {
	p.Status.SetStatus(PROCESS_STATUS_STARTED)
	outFile, err := os.OpenFile(p.OutFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		return err
	}
	errFile, err := os.OpenFile(p.ErrFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		return err
	}
	// Get current directory working path
	wd, _ := os.Getwd()
	procAtr := &os.ProcAttr{
		Dir: wd,
		Env: os.Environ(),
		Files: []*os.File{
			os.Stdin,
			outFile,
			errFile,
		},
	}
	args := append([]string{p.Name}, p.Args...)
	process, err := os.StartProcess(p.Cmd, args, procAtr)
	if err != nil {
		return err
	}
	p.process = process
	p.Pid = p.process.Pid
	err = ioutil.WriteFile(p.PidFile, []byte(strconv.Itoa(proc.process.Pid)), 0660)
	if err != nil {
		return err
	}
	p.Status.InitUptime()
	p.Status.SetStatus(PROCESS_STATUS_RUNNING)
	return nil
}

// ForceStop will forcefully send a SIGKILL signal to process killing it instantly.
func (p *Process) ForceStop() error {
	if p.process != nil {
		err := p.process.Signal(syscall.SIGKILL)
		p.Status.SetStatus(PROCESS_STATUS_STOPPED)
		p.Release()
		return err
	}
	return errors.New("Process does not exist.")
}

// GracefullyStop will send a SIGTERM signal asking the process to terminate.
// The process may choose to die gracefully or ignore this signal completely.
// In that case the process will keep running unless you call ForceStop()
// Returns an error in case there's any.
func (p *Process) GracefullyStop() error {
	if p.process != nil {
		err := p.process.Signal(syscall.SIGTERM)
		p.Status.SetStatus(PROCESS_STATUS_AKSED_TO_STOP)
		return err
	}
	return errors.New("Process does not exist.")
}

// Restart will try to gracefully stop the process and then start it again.
func (p *Process) Restart() error {
	if p.IsAlive() {
		err := p.GracefullyStop()
		if err != nil {
			return err
		}
	}
	return p.Start()
}

// Delete will delete everything created by this process, including the out, err and pid file.
func (p *Process) Delete() error {
	p.Release()
	err := os.Remove(p.Outfile)
	if err != nil {
		return err
	}
	err = os.Remove(p.Errfile)
	if err != nil {
		return err
	}
	return os.RemoveAll(p.Path)
}

// IsAlive will check if the process is alive or not.
func (p *Process) IsAlive() bool {
	p, err := os.FindProcess(p.Pid)
	if err != nil {
		return false
	}
	return p.Signal(syscall.Signal(0)) == nil
}

// Identifier is that will be used by watcher to keep track of its processes
func (p *Process) Identifier() string {
	return p.Name
}

// ShouldKeepAlive will returns true if the process should be kept alive or not
func (p *Process) ShouldKeepAlive() bool {
	return p.KeepAlive
}

// AddRestart is add one restart to process status
func (p *Process) AddRestart() {
	p.Status.AddRestart()
}

// NotifyStopped that process was stopped so we can set its PID to -1
func (p *Process) NotifyStopped() {
	p.Pid = -1
}

// SetStatus will set process status
func (p *Process) SetStatus(status string) {
	p.Status.SetStatus(status)
}

// SetUptime will set Uptime
func (p *Process) SetUptime() {
	p.Status.SetUptime()
}

// SetSysInfo will get current process cpu and memory usage
func (p *Process) SetSysInfo() {
	p.Status.SetSysInfo(p.process.Pid)
}

// ResetUpTime will set Uptime
func (p *Process) ResetUpTime() {
	p.Status.ResetUptime()
}

// GetPid will return process current PID
func (p *Process) GetPid() int {
	return p.Pid
}

// GetStatus will return process current status
func (p *Process) GetStatus() *ProcessStatus {
	if !p.IsAlive() {
		p.ResetUpTime()
	} else {
		// Update uptime
		p.SetUptime()
	}
	// Update cpu and memory
	p.SetSysInfo()

	return p.Status
}

// GetPath will return process path
func (p *Process) GetPath() string {
	return p.Path
}

// GetOutFile will return process out file
func (p *Process) GetOutFile() string {
	return p.OutFile
}

// GetPidFile will return process pid file
func (p *Process) GetPidFile() string {
	return p.PidFile
}

// GetErrFile will return process error file
func (p *Process) GetErrFile() string {
	return p.ErrFile
}

// GetName will return current process name
func (p *Process) GetName() string {
	return p.Name
}

// Watch will stop execution and wait until the process change its state. Usually changing state, means that the process died.
// Returns a tuple with the new process state and an error in case there's any.
func (p *Process) Watch() (*os.ProcessState, error) {
	return p.process.Wait()
}

// Will release the process and remove its PID file
func (p *Process) Release() {
	if p.process != nil {
		p.process.Release()
	}
	os.Remove(p.Pidfile)
}
