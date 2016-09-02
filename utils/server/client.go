package server

import (
	"net/http"
	"time"
)

var Client *http.Client = &http.Client{Timeout: time.Duration(0)}
