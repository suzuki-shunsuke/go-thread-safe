package safe

import (
	"sync"
)

// Bool wraps bool.
// Bool must be used as the pointer because Bool has sync.RWMutex as a private field.
// A RWMutex must not be copied after first use.
// https://golang.org/pkg/sync/#RWMutex
type Bool struct {
	value bool
	mutex sync.RWMutex
}

func (s *Bool) Get() bool {
	s.mutex.RLock()
	v := s.value
	s.mutex.RUnlock()
	return v
}

func (s *Bool) Set(v bool) {
	s.mutex.Lock()
	s.value = v
	s.mutex.Unlock()
}

func (s *Bool) SetFunc(f func(v bool) bool) {
	s.mutex.Lock()
	s.value = f(s.value)
	s.mutex.Unlock()
}

func (s *Bool) Invert() {
	s.mutex.Lock()
	s.value = !s.value
	s.mutex.Unlock()
}

func (s *Bool) InvertR() bool {
	s.mutex.Lock()
	a := !s.value
	s.value = a
	s.mutex.Unlock()
	return a
}
