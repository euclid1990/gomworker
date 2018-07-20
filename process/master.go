package process

import (
	"errors"
	cf "github.com/euclid1990/gomworker/configs"
	util "github.com/euclid1990/gomworker/utilities"
	"os"
	"path"
	"sync"
	"time"
)

const (
	UPDATE_INTERVAL = 30 // Unit: second
)

type RunProcessCb func() bool

type Master struct {
	sync.Mutex
	Dir       string                      // Dir is the main directory
	Watcher   *Watcher                    // Watcher is a watcher instance
	Processes map[string]ProcessContainer // Processes is a map containing all started processes
}

// NewMaster will create a Master instance.
func NewMaster(dir string) *Master {
	m := &Master{}
	// Get current directory working path
	wd, _ := os.Getwd()
	if dir == "" {
		dir = wd
	}
	os.MkdirAll(path.Dir(dir), 0777)
	m.Dir = path.Dir(wd)
	m.Processes = make(map[string]*ProcessContainer)
	util.Log(cf.LOG_INFO, "[Master] Init Watcher.")
	m.Watcher = NewWatcher()
	util.Log(cf.LOG_INFO, "[Master] Revive all processes listed on ListProcesses.")
	m.Revive()
	util.Log(cf.LOG_INFO, "[Master] Keep processes running forever.")
	go m.Watch()
	util.Log(cf.LOG_INFO, "[Master] Run update processes status in Goroutines.")
	go m.UpdateStatus()
}

// Get one process by name and return it
func (m *Master) GetProcess(name string) ProcessContainer {
	if p, ok := m.Processes[name]; ok {
		return p
	}
	return nil
}

// ListProcesses will return a list of all processes.
func (m *Master) ListProcesses() []ProcessContainer {
	pList := []ProcessContainer{}
	for _, v := range m.Processes {
		pList = append(pList, v)
	}
	return pList
}

// Revive will start-again all processes on Master.ListProcesses.
func (m *Master) Revive() error {
	m.Lock()
	defer m.Unlock()
	pList := m.ListProcesses()
	util.Logf(cf.LOG_INFO, "[Master] Revive Total: %v Processes.", len(pList))
	for _, p := range pList {
		if !p.ShouldKeepAlive() {
			continue
		}
		identifier := pre.Identifier()
		util.Logf(cf.LOG_INFO, "[Master] Reviving Process: %v.", identifier)
		if err := m.start(p); err != nil {
			util.Logf(cf.LOG_INFO, "[Master] Failed revive process: %v.", identifier)
			return fmt.Errorf("Failed revive process: %v.", identifier)
		}
	}
	return nil
}

// UpdateStatus will update a process status every 30s.
func (m *Master) UpdateStatus() {
	for {
		m.Lock()
		pList := m.ListProcesses()
		for _, p := range pList {
			m.updateStatusProcess(p)
		}
		m.Unlock()
		time.Sleep(time.Second * UPDATE_INTERVAL)
	}
}

// Update status of an process
func (m *Master) updateStatusProcess(p ProcessContainer) {
	if p.IsAlive() {
		p.SetStatus(PROCESS_STATUS_RUNNING)
	} else {
		p.NotifyStopped()
		p.SetStatus(PROCESS_STATUS_STOPPED)
	}
}

// RunProcess will run *Prepare(process) and add it to the watcher.
func (m *Master) RunProcess(pre *Prepare, cb RunProcessCb) error {
	m.Lock()
	defer m.Unlock()
	identifier := pre.Identifier()
	if _, ok := m.Processes[identifier]; ok {
		util.Logf(cf.LOG_WARNING, "[Master] Process %s is already exist.", identifier)
		return errors.New("Unable to run a process that already exist.")
	}
	// Try to run prepare process
	p, err := pre.Start()
	if err != nill {
		return err
	}
	m.Processes[identifier] = p
	m.Watcher.Add(p)
	cb()
	return nil
}

// StartProcess will a start a process with the given name.
func (m *Master) StartProcess(name string) error {
	m.Lock()
	defer m.Unlock()
	if p, ok := m.Processes[name]; ok {
		return m.start(p)
	}
	return errors.New("Unable to start an unknown process.")
}

// NOT thread safe method.
func (m *Master) start(p ProcessContainer) error {
	if p.IsAlive() {
		return errors.New("Unable to start an running process.")
	}
	err := p.Start()
	if err != nil {
		return err
	}
	p.SetUptime()
	m.Watcher.Add(p)
	return nil
}

// StopProcess will stop a process with the given name.
func (m *Master) StopProcess(name string) error {
	m.Lock()
	defer m.Unlock()
	if p, ok := m.Processes[name]; ok {
		return m.stop(p)
	}
	return errors.New("Unable to start an unknown process.")
}

// NOT thread safe method.
func (m *Master) stop(p ProcessContainer) {
	if p.IsAlive() {
		identifier := p.Identifier()
		waitStop := m.Watcher.StopWatcher(identifier)
		err := p.GracefullyStop()
		if err != nil {
			return err
		}
		if waitStop != nil {
			<-waitStop
			p.NotifyStopped()
			p.SetStatus(PROCESS_STATUS_STOPPED)
		}
	}

}

// RestartProcess will restart a process with the given name.
func (m *Master) RestartProcess(name string) {
	m.Lock()
	defer m.Unlock()
	if p, ok := m.Processes[name]; ok {
		return m.restart(p)
	}
	return errors.New("Unable to restart an unknown process.")
}

// NOT thread safe method.
func (m *Master) restart(p ProcessContainer) error {
	p.AddRestart()
	if err := m.stop(p); err != nil {
		return err
	}
	err := m.start(p)
	return err
}

// DeleteProcess will delete a process with the given name.
func (m *Master) DeleteProcess(name string) {
	m.Lock()
	defer m.Unlock()
	if p, ok := m.Processes[name]; ok {
		return m.delete(p)
	}
	return errors.New("Unable to delete an unknown process.")
}

// NOT thread safe method.
func (m *Master) delete(p ProcessContainer) error {
	if err := m.stop(p); err != nil {
		return err
	}
	identifier := p.Identifier()
	delete(m.Processes, identifier)
	err := p.Delete()
	return err
}
