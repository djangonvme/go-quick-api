package app

import (
	"sync"
	"time"
)

// var DLock

type LockerIf interface {
	Lock(key string, value string, exp time.Duration) (bool, error)
	Unlock(key string, value string) (bool, error)
}

type ProcessLock struct {
	status map[string]string
	m      sync.RWMutex
}

func NewProcessLocker() *ProcessLock {
	return &ProcessLock{
		status: make(map[string]string),
	}
}
func (c *ProcessLock) Lock(key string, value string, exp time.Duration) (bool, error) {
	c.m.Lock()
	defer c.m.Unlock()
	if _, ok := c.status[key]; ok {
		return false, nil
	}
	c.status[key] = value
	return true, nil
}

func (c *ProcessLock) Unlock(key string, value string) (bool, error) {
	c.m.Lock()
	defer c.m.Unlock()
	if ov, ok := c.status[key]; ok {
		if value != "" && ov != value {
			return false, nil
		}
		delete(c.status, key)
	}
	return true, nil
}
