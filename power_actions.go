package main

import (
	"github.com/cameliot/alpaca"
)

const POWER_ON = "@power/ON"
const POWER_OFF = "@power/OFF"
const POWER_STATE = "@power/STATE"

func PowerState(on bool) alpaca.Action {
	return alpaca.Action{
		Type:    POWER_STATE,
		Payload: on,
	}
}
