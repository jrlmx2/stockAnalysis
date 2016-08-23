package term

import (
	"os"
	"os/signal"
	"syscall"
)

// NewTerm Spawns a process to signal any program interupts should shut down.
func NewTerm() *int {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	temp := 0
	die := &temp
	go func() {
		<-sigc
		temp = 1
	}()
	return die
}
