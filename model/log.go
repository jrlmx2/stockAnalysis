package model

import logging "github.com/op/go-logging"

var logger *logging.Logger

func SetLogger(l *logging.Logger) {
	logger = l
}
