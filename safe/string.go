package safe

import (
	"encoding/json"
	"sync"
)

// String wraps bool.
// String must be used as the pointer because String has sync.RWMutex as a private field.
// A RWMutex must not be copied after first use.
// https://golang.org/pkg/sync/#RWMutex
type String struct {
	value string
	mutex sync.RWMutex
}

func (s *String) MarshalJSON() ([]byte, error) {
	s.mutex.RLock()
	v := s.value
	s.mutex.RUnlock()
	return json.Marshal(v)
}

func (s *String) UnmarshalJSON(b []byte) error {
	s.mutex.Lock()
	err := json.Unmarshal(b, &s.value)
	s.mutex.Unlock()
	return err
}

func (s *String) String() string {
	s.mutex.RLock()
	v := s.value
	s.mutex.RUnlock()
	return "String{" + v + "}"
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

func (s *String) Add(v string) {
	s.mutex.Lock()
	s.value += v
	s.mutex.Unlock()
}

func (s *String) AddR(v string) string {
	s.mutex.Lock()
	a := s.value + v // escapes to heap
	s.value = a
	s.mutex.Unlock()
	return a
}
