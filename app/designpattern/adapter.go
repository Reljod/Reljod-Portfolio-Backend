package designpattern

import (
	log "github.com/Reljod/Reljod-Portfolio-Backend/app/logger"
)

var logger = log.Logger

type IProcess interface {
	Process()
}

type Adapter struct {
	adaptee Adaptee
}

func (adapter Adapter) Process() {
	logger.Info("Adapter Process")
	adapter.adaptee.convert()
}

type Adaptee struct {
	adapterType int
}

func (adaptee Adaptee) convert() {
	logger.Info("Adaptee Convert Method")
}
