package process

// ProcessContainer is a interface that about process
type ProcessContainer interface {
	Start() error
	ForceStop() error
	GracefullyStop() error
	Restart() error
	Delete() error
	IsAlive() bool
	Identifier() string
	ShouldKeepAlive() bool
	AddRestart()
	NotifyStopped()
	SetStatus(status string)
	SetUptime()
	SetSysInfo()
	GetPid() int
	GetStatus() *ProcessStatus
	GetPath() string
	GetOutFile() string
	GetPidFile() string
	GetErrFile() string
	GetName() string
	Watch() (*os.ProcessState, error)
	Release()
}
