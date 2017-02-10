package model

import "github.com/jrlmx2/stockAnalysis/utils/logger"

var log *logger.Logger

func SetLogger(l *logger.Logger) {
	log = l
	log.Info("Model start...")
}
