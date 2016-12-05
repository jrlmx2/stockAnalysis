package listener

import (
	"errors"
	"net"
	"os"
	"time"

	"github.com/jrlmx2/stockAnalysis/utils/term"
)

type Listener struct {
	*net.TCPListener                 //Wrapped listener
	stop             *chan os.Signal //Channel used only to indicate listener should shutdown
}

func New(l net.Listener) (*Listener, error) {
	tcpL, ok := l.(*net.TCPListener)

	if !ok {
		return nil, errors.New("Cannot wrap listener")
	}

	retval := &Listener{}
	retval.TCPListener = tcpL
	retval.stop = term.Channel()

	return retval, nil
}

var StoppedError = errors.New("Listener stopped")

func (sl *Listener) Accept() (net.Conn, error) {

	for {
		//Wait up to one second for a new connection
		sl.SetDeadline(time.Now().Add(time.Second))

		newConn, err := sl.TCPListener.Accept()

		//Check for the channel being closed
		select {
		case <-*sl.stop:
			return nil, StoppedError
		default:
			//If the channel is still open, continue as normal
		}

		if err != nil {
			netErr, ok := err.(net.Error)

			//If this is a timeout, then continue to wait for
			//new connections
			if ok && netErr.Timeout() && netErr.Temporary() {
				continue
			}
		}

		return newConn, err
	}
}

func (sl *Listener) Stop() {
	close(*sl.stop)
}
