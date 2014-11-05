package main

import (
	"fmt"
	"net"
	"time"
)

type Backend interface {
	IsDown() bool
	Verify()
}

type BackendGroup interface {
	SelectBackend() Backend
	VerifyBackends()
}

type BaseBackend struct {
	Ip             int64
	Port           int16
	Down           bool
	LastValidCheck time.Time
}

type BaseBackendGroup struct {
	Group []BaseBackend
}

func NewBaseBackend(ip int64, port int16) (*BaseBackend, error) {
	if port < 0 || ip < 0 {
		return nil, fmt.Errorf("IP and Port can't have negative values!")
	}
	return &BaseBackend{Ip: ip, Port: port}, nil
}

func (bb *BaseBackend) IsDown() bool {
	return bb.Down
}

func (bb *BaseBackend) Verify() {
	_, err := net.DialTimeout("tcp", fmt.Sprintf("%d:%d", bb.Ip, bb.Port), time.Second*10)
	// Timeout... Set it as down!
	if err != nil {
		if bb.Down == false {
			bb.LastValidCheck = time.Now()
		}
		bb.Down = true
		return
	}
	// No error
	bb.LastValidCheck = time.Now()
	return
}
