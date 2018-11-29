package main

import (
	"github.com/cameliot/alpaca"
)

type PowerSvc struct {
	on bool
}

func NewPowerSvc() *PowerSvc {
	return &PowerSvc{
		on: false,
	}
}

func (self *PowerSvc) Handle(actions alpaca.Actions, dispatch alpaca.Dispatch) {

	// Initialize state
}
