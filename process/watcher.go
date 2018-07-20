package process

import (
	cf "github.com/euclid1990/gomworker/configs"
	util "github.com/euclid1990/gomworker/utilities"
	"os"
	"sync"
)

// ProcessState is a wrapper with the process state and an error in case there's any.
type ProcessState struct {
	state *os.ProcessState
	err   error
}

// ProcessWatcher is a wrapper that act as a object that watches a process.
type ProcessWatcher struct {
	pState      chan *ProcessState
	pContainer  ProcessContainer
	stopWatcher chan bool
}

// Watcher is responsible for watching a list of processes and report to Master in case the process dies at some point.
type Watcher struct {
	sync.Mutex
	restartProc chan ProcessContainer      // A channel with the dead processes that need to be restarted.
	watchProcs  map[string]*ProcessWatcher // A map of list process identifier
}

// NewWatcher will create a Watcher instance.
func NewWatcher() *Watcher {
	watcher := &Watcher{
		restartProc: make(chan ProcessContainer),
		watchProcs:  make(map[string]*ProcessWatcher),
	}
	return watcher
}

// GetRestartProc is a wrapper to export the channel restartProc.
func (w *Watcher) GetRestartProc() chan ProcessContainer {
	return w.restartProc
}

// GetWatcherProcs is a wrapper to export the map watchProcs.
func (w *Watcher) GetWatcherProcs() map[string]*ProcessWatcher {
	return w.watchProcs
}

// Add a process to watcher list
func (w *Watcher) Add(p ProcessContainer) {
	w.Lock()
	defer w.Unlock()

	identifier := p.Identifier()

	if _, ok := w.watchProcs[id]; ok {
		util.Logf(cf.LOG_WARNING, "A watcher for this process [%s] already exists.", identifier)
		return
	}

	processWatcher := &ProcessWatcher{
		pState:      make(chan *ProcessState, 1),
		pContainer:  p,
		stopWatcher: make(chan bool, 1),
	}
	w.watchProcs[identifier] = processWatcher

	// Goroutine: Watching process state
	go func() {
		util.Logf(cf.LOG_INFO, "Starting watcher on process [%s].", identifier)
		// Waits for the Process to exit, and then returns a ProcessState
		state, err := p.Watch()
		// Send state cho channel
		processWatcher.pState <- &ProcessState{
			state: state,
			err:   err,
		}
	}()

	// Goroutine: Remove process in watchProcs to stop or restart process
	go func() {
		defer delete(w.watchProcs, identifier)
		select {
		case processState := <-processWatcher.pState:
			util.Logf(cf.LOG_INFO, "Process [%s] is dead, advising master...", identifier)
			util.Logf(cf.LOG_INFO, "State of Process [%s] is [%s] with error [%v].", identifier, processState.state.String(), pState.err)
			w.restartProc <- processWatcher.pContainer
			break
		case <-processWatcher.stopWatcher:
			break
		}
	}()
}

// StopWatcher will stop a running watcher on a process with identifier 'identifier'
// Returns a channel that will be populated when the watcher is finally done.
func (w *Watcher) StopWatcher(identifier string) chan bool {
	if processWatcher, ok := w.watchProcs[identifier]; ok {
		util.Logf(cf.LOG_INFO, "Stopping watcher on process [%s].", identifier)
		// Delete process from Watcher.watchProcs
		processWatcher.stopWatcher <- true
		waitStop := make(chan bool, 1)
		go func() {
			// Wait to process change state to dead because it is killed at anywhere
			<-processWatcher.pState
			// Send notify to waitStop to report process release successfully
			waitStop <- true
		}()
		return waitStop
	}
	return nil
}
