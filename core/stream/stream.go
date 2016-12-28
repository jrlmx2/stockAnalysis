package stream

import "bufio"

type Stream interface {
	Connection() *bufio.Reader
	Reopen()
}
