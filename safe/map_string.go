package safe

import (
	"encoding/json"
	"fmt"
	"sync"
)

// MapString wraps map[string]string.
// MapString must be created by NewMapString.
type MapString struct {
	value map[string]string
	mutex sync.RWMutex
}

// NewMapString creates a MapString.
// The argument `value` must not be nil.
// Note that the argument `value` is holden in MapString, so don't read and write `value` out of the MapString.
func NewMapString(value map[string]string) *MapString {
	// To avoid the heap allocation, don't copy `value` and create a new map.
	return &MapString{ // escapes to heap
		value: value,
	}
}

func (m *MapString) String() string {
	m.mutex.RLock()
	v := "MapString{" + fmt.Sprintf("%v", m.value) + "}"
	m.mutex.RUnlock()
	return v
}

func (m *MapString) MarshalJSON() ([]byte, error) {
	m.mutex.RLock()
	b, err := json.Marshal(m.value)
	m.mutex.RUnlock()
	return b, err
}

func (m *MapString) UnmarshalJSON(buf []byte) error {
	m.mutex.Lock()
	err := json.Unmarshal(buf, &m.value)
	m.mutex.Unlock()
	return err
}

// Get gets a value from the map with lock.
func (m *MapString) Get(k string) string {
	m.mutex.RLock()
	v := m.value[k]
	m.mutex.RUnlock()
	return v
}

// Get gets a value from the map with lock.
func (m *MapString) GetOk(k string) (string, bool) {
	m.mutex.RLock()
	v, ok := m.value[k]
	m.mutex.RUnlock()
	return v, ok
}

// Has checks whether the map has the key with lock.
func (m *MapString) Has(k string) bool {
	m.mutex.RLock()
	_, ok := m.value[k]
	m.mutex.RUnlock()
	return ok
}

// Len gets the length of the map with lock.
func (m *MapString) Len() int {
	m.mutex.RLock()
	v := len(m.value)
	m.mutex.RUnlock()
	return v
}

// Delete deletes the key from the map with lock.
func (m *MapString) Delete(k string) {
	m.mutex.Lock()
	delete(m.value, k)
	m.mutex.Unlock()
}

// DeleteR deletes the key from the map and returns the value with lock.
func (m *MapString) DeleteR(k string) string {
	m.mutex.Lock()
	v := m.value[k]
	delete(m.value, k)
	m.mutex.Unlock()
	return v
}

// DeleteROk deletes the key from the map and returns the value with lock.
func (m *MapString) DeleteROk(k string) (string, bool) {
	m.mutex.Lock()
	v, ok := m.value[k]
	if ok {
		delete(m.value, k)
	}
	m.mutex.Unlock()
	return v, ok
}

// Set sets the key and value to the map with lock.
func (m *MapString) Set(k, v string) {
	m.mutex.Lock()
	m.value[k] = v
	m.mutex.Unlock()
}

// SetDefault sets the key and value to the map if the map doesn't have the key with lock.
func (m *MapString) SetDefault(k, v string) {
	m.mutex.Lock()
	if _, ok := m.value[k]; !ok {
		m.value[k] = v
	}
	m.mutex.Unlock()
}

// SetDefaultR sets the key and value to the map if the map doesn't have the key and returns the value with lock.
// true is returned if the map has already haven the key and the value isn't updated.
func (m *MapString) SetDefaultR(k, v string) (string, bool) {
	m.mutex.Lock()
	a, ok := m.value[k]
	if !ok {
		m.value[k] = v
		a = v
	}
	m.mutex.Unlock()
	return a, ok
}

// SetFunc gets a value of the key from the map and calls the function and sets the returned value to the map with lock.
// This is used to update the value based on the original value atomicaly.
func (m *MapString) SetFunc(k string, f func(string, bool) string) {
	m.mutex.Lock()
	v, ok := m.value[k]
	m.value[k] = f(v, ok)
	m.mutex.Unlock()
}

// Range gets all pairs of the key and value from the map with lock and calls the function.
func (m *MapString) Range(f func(k, v string)) {
	m.mutex.RLock()
	copiedM := make(map[string]string, len(m.value))
	for k, v := range m.value {
		copiedM[k] = v
	}
	m.mutex.RUnlock()
	for k, v := range copiedM {
		f(k, v)
	}
}

// RangeB gets all pairs of the key and value from the map with lock and calls the function.
// If the function returns false, the loop ends.
func (m *MapString) RangeB(f func(k, v string) bool) {
	m.mutex.RLock()
	copiedM := make(map[string]string, len(m.value))
	for k, v := range m.value {
		copiedM[k] = v
	}
	m.mutex.RUnlock()

	for k, v := range copiedM {
		if !f(k, v) {
			break
		}
	}
}

// Copy copies and creates a new MapString.
func (m *MapString) Copy(target *MapString) {
	m.mutex.RLock()
	for k, v := range m.value {
		target.value[k] = v
	}
	m.mutex.RUnlock()
}

// CopyData copies an internal map[string]string to target.
func (m *MapString) CopyData(target map[string]string) {
	m.mutex.RLock()
	for k, v := range m.value {
		target[k] = v
	}
	m.mutex.RUnlock()
}
