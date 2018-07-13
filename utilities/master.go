package utilities

import (
	cf "github.com/euclid1990/gomworker/configs"
	"github.com/sevlyar/go-daemon"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"sync"
	"syscall"
)

type Master struct {
	ctx     *daemon.Context
	wSignal os.Signal
}

func NewMaster() *Master {
	m := &Master{}
	m.GetCtx()
	return m
}

func (m *Master) GetCtx() *daemon.Context {
	ctx := &daemon.Context{
		PidFileName: path.Join(filepath.Dir(cf.CONTEXT_PID_PATH), cf.CONTEXT_PID_MAIN_FILENAME),
		PidFilePerm: 0644,
		LogFileName: path.Join(filepath.Dir(cf.CONTEXT_LOG_PATH), cf.CONTEXT_LOG_MAIN_FILENAME),
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
	}
	m.ctx = ctx
	return ctx
}

func (m *Master) IsRunning() (bool, *os.Process, error) {
	running := true
	d, err := m.ctx.Search()
	if err != nil {
		running = false
	} else if err := d.Signal(syscall.Signal(0)); err != nil {
		running = false
	}
	return running, d, err
}

func (m *Master) Kill(pid int, signal os.Signal) error {
	p, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	defer p.Release()
	Logf(cf.LOG_INFO, "Kill %v with %v", pid, signal)
	return p.Signal(signal)
}

func (m *Master) WaitChildSignal(wg *sync.WaitGroup) {
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	c := make(chan os.Signal, 1)
	// Only SIGUSR1/SIGUSR2 will be sent to the channel.
	signal.Notify(c, syscall.SIGUSR1, syscall.SIGUSR2)
	wg.Add(1)
	go func() {
		m.wSignal = <-c
		defer wg.Done()
	}()
}

func (m *Master) Start() {
	var wg sync.WaitGroup
	m.WaitChildSignal(&wg)
	if r, _, _ := m.IsRunning(); r {
		Log(cf.LOG_INFO, "Gomworker is already running.")
		return
	}

	d, err := m.ctx.Reborn()
	if err != nil {
		Logf(cf.LOG_ERROR, "Unable to start Gomworker: %v", err)
	}

	pid := os.Getpid()

	if d != nil {
		Logf(cf.LOG_INFO, "Parent's process ID: %v", pid)
		wg.Wait()
		if m.wSignal == syscall.SIGUSR1 {
			Log(cf.LOG_INFO, "Gomworker successfully started.")
			return
		}
	} else {
		Log(cf.LOG_INFO, "Runs its own copy on child's process.")
		Logf(cf.LOG_INFO, "Child's process ID: %v", pid)
		m.Kill(os.Getpid(), syscall.SIGUSR1)
		wg.Wait()
		defer m.ctx.Release()
	}

	Log(cf.LOG_INFO, "Gomworker starting up ...")
	// TO BE DEFINED

	// Send signal to parent's process to kill goroutine
	m.Kill(os.Getppid(), syscall.SIGUSR1)

	// Listen signals sent to child's process to kill daemon
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// Block until a signal is received.
	<-c

	Log(cf.LOG_INFO, "Gomworker got signal to stop.")
	// TO BE DEFINED

	os.Exit(0)
}

func (m *Master) Stop() {
	Log(cf.LOG_INFO, "Gomworker stopping ...")
	if r, d, _ := m.IsRunning(); r {
		if err := d.Signal(syscall.Signal(syscall.SIGQUIT)); err != nil {
			Logf(cf.LOG_ERROR, "Gomworker failed to stop daemon: %v", err)
		}
	} else {
		m.ctx.Release()
		Log(cf.LOG_ERROR, "Gomworker is not running to stop.")
	}
	Log(cf.LOG_INFO, "Gomworker daemon is terminated.")
}
