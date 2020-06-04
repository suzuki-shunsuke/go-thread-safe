package safe

import (
	"sync"
)

type Bool struct {
	value bool
	mutex sync.RWMutex
}

func NewBool(v bool) *Bool {
	return &Bool{
		value: v,
	}
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
