package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

/*
 The PSU is connected to the RPi.
 We use dead simple interaction with the sys/class/gpio
 interface to switch it on and off.
*/

func _initializeGpio(pin int) {
	exec.Command("echo", fmt.Sprintf("%d", pin), ">", "/sys/class/gpio/export")
	exec.Command("echo", "out", ">",
		fmt.Sprintf("/sys/class/gpio%d/direction", pin))
	time.Sleep(500 * time.Millisecond)
}

func _isGpioInitialized(pin int) bool {
	iopath := fmt.Sprintf("/sys/class/gpio%d", pin)
	if _, err := os.Stat(iopath); !os.IsNotExist(err) {
		return false
	}
	return true
}

func _assertGpioInitialized(pin int) {
	if _isGpioInitialized(pin) {
		return
	}
	_initializeGpio(pin)
}

func readPowerState(pin int) bool {
	return false
}
