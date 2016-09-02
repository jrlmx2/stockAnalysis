package term

import (
	"os"
	"os/signal"
	"syscall"
)

var sigc chan os.Signal

var die int

// NewTerm Spawns a process to signal any program interupts should shut down.
func NewTerm() {
	if sigc == nil {
		sigc = make(chan os.Signal, 1)
		signal.Notify(sigc,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGQUIT)

		go func() {
			<-sigc
			die = 1
		}()
	}
}

func Channel() *chan os.Signal {
	return &sigc
}

func Kill() {
	die = 1
}

func WasTerminated() bool {
	if sigc == nil {
		NewTerm()
	}
	if die == 1 {
		return true
	} else {
		return false
	}
}
