package safe

import (
	"sync"
)

type String struct {
	value string
	mutex sync.RWMutex
}

func NewString(v string) *String {
	return &String{
		value: v,
	}
}

func (s *String) Get() string {
	s.mutex.RLock()
	v := s.value
	s.mutex.RUnlock()
	return v
}

func (s *String) Set(v string) {
	s.mutex.Lock()
	s.value = v
	s.mutex.Unlock()
}

func (s *String) SetFunc(f func(v string) string) {
	s.mutex.Lock()
	s.value = f(s.value)
	s.mutex.Unlock()
}

func (s *String) SetFuncR(f func(v string) string) string {
	s.mutex.Lock()
	v := f(s.value)
	s.value = v
	s.mutex.Unlock()
	return v
}

func (s *String) Add(v string) {
	s.mutex.Lock()
	s.value += v
	s.mutex.Unlock()
}

func (s *String) AddR(v string) string {
	s.mutex.Lock()
	a := s.value + v
	s.value = a
	s.mutex.Unlock()
	return a
}
