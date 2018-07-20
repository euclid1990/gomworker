package process

import (
	"path"
)

type Prepare struct {
	Name       string
	SourcePath string
	Cmd        string
	SysFolder  string
	Language   string
	KeepAlive  bool
	Args       []string
}

// Identifier is a function that get process name
func (pre *Prepare) Identifier() string {
	return pre.Name
}

// Get process base path
func (pre *Prepare) getPath() string {
	return path.Join(pre.SysFolder, pre.Name)
}

// Get process id path
func (pre *Prepare) getPidPath() string {
	return pre.getPath() + ".pid"
}

// Get process output path
func (pre *Prepare) getOutPath() string {
	return pre.getPath() + ".out"
}

// Get process error path
func (pre *Prepare) getErrPath() string {
	return pre.getPath() + ".err"
}

func (pre *Prepare) Start() (*ProcessContainer, error) {
	p := &Process{
		Name:      pre.Name,
		Cmd:       pre.Cmd,
		Args:      pre.Args,
		Path:      pre.getPath(),
		Path:      pre.getPath(),
		PidFile:   pre.getPidPath(),
		OutFile:   pre.getOutPath(),
		ErrFile:   pre.getErrPath(),
		KeepAlive: pre.KeepAlive,
		Status:    &ProcessStatus{},
	}
	err := p.Start()
	return p, err
}
