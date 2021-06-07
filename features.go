package fflags

import (
	"strings"
	"sync"
)

// Feature is the flag representing a given feature tree
type Feature string

type features struct {
	l      sync.RWMutex
	status map[Feature]bool
}

var f *features

func init() {
	f = &features{
		status: make(map[Feature]bool),
	}
}

// Define defines the feature with the given name and marks it disabled
func Define(name string) Feature {
	f.l.Lock()
	defer f.l.Unlock()

	canonical := Feature(strings.ToLower(name))
	if _, exists := f.status[canonical]; !exists {
		f.status[canonical] = false
	}
	return canonical
}

// Enable will mark the given feature and all its sub-features as enabled
func (feature Feature) Enable() {
	f.Enable(feature)
}

func (feature Feature) Disable() {
	f.Disable(feature)
}

func IsEnabled(feature Feature) bool {
	return f.isActive(feature)
}

func IfEnabled(feature Feature, do func() error) error {
	if f.isActive(feature) {
		return do()
	}
	return nil
}

func (s *features) Enable(feature Feature) {
	s.l.Lock()
	defer s.l.Unlock()

	s.status[feature] = true
}

func (s *features) Disable(feature Feature) {
	s.l.Lock()
	defer s.l.Unlock()

	s.status[feature] = false
	// TODO: take care of sub-features features
}

func (s *features) isActive(feature Feature) bool {
	// TODO: Make this more efficient
	parts := strings.Split(string(feature), ".")
	path := make([]Feature, len(parts))
	for i := len(parts); i > 0; i-- {
		name := strings.Join(parts[0:i], ".")
		path[len(path)-i] = Feature(name)
	}
	s.l.RLock()
	defer s.l.RUnlock()
	for _, p := range path {
		if active := s.status[p]; active {
			return true
		}
	}
	return false
}
