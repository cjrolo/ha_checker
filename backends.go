package main

import (
	"fmt"
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
	Ip   int64
	Port int16
	Down bool
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
