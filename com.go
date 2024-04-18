package Kiosk_Firmware_Neo

import (
	"go.bug.st/serial"
	"sync"
)

type PortInterface struct {
	mu   sync.RWMutex
	port *serial.Port
}

func NewPortInterface(port *serial.Port) *PortInterface {
	return &PortInterface{
		port: port,
	}
}
