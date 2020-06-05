package safe

import (
	"strconv"
	"sync"
)

// Int wraps a int.
// Int must be used as the pointer because Int has sync.RWMutex as a private field.
// A RWMutex must not be copied after first use.
// https://golang.org/pkg/sync/#RWMutex
type Int struct {
	value int
	mutex sync.RWMutex
}

func (i *Int) String() string {
	i.mutex.RLock()
	v := i.value
	i.mutex.RUnlock()
	return "Int{" + strconv.Itoa(v) + "}"
}

// Get gets a value with lock.
func (i *Int) Get() int {
	i.mutex.RLock()
	v := i.value
	i.mutex.RUnlock()
	return v
}

// Set sets a value with lock.
func (i *Int) Set(v int) {
	i.mutex.Lock()
	i.value = v
	i.mutex.Unlock()
}

// SetFunc gets a value and calls the function and sets the returned value with lock.
// This is used to update the value based on the original value atomicaly.
func (i *Int) SetFunc(f func(v int) int) {
	i.mutex.Lock()
	i.value = f(i.value)
	i.mutex.Unlock()
}

// Add adds a value with lock.
func (i *Int) Add(v int) {
	i.mutex.Lock()
	i.value += v
	i.mutex.Unlock()
}

// AddR adds a value with lock.
func (i *Int) AddR(v int) int {
	i.mutex.Lock()
	a := i.value + v
	i.value = a
	i.mutex.Unlock()
	return a
}

// Sub substitutes a value with lock.
func (i *Int) Sub(v int) {
	i.mutex.Lock()
	i.value -= v
	i.mutex.Unlock()
}

// SubR substitutes a value with lock.
func (i *Int) SubR(v int) int {
	i.mutex.Lock()
	a := i.value - v
	i.value = a
	i.mutex.Unlock()
	return a
}

// Mul multiplies a value with lock.
func (i *Int) Mul(v int) {
	i.mutex.Lock()
	i.value *= v
	i.mutex.Unlock()
}

// MulR multiplies a value with lock.
func (i *Int) MulR(v int) int {
	i.mutex.Lock()
	a := i.value * v
	i.value = a
	i.mutex.Unlock()
	return a
}

// Div divides a value with lock.
func (i *Int) Div(v int) {
	i.mutex.Lock()
	i.value /= v
	i.mutex.Unlock()
}

// DivR divides a value with lock.
func (i *Int) DivR(v int) int {
	i.mutex.Lock()
	a := i.value / v
	i.value = a
	i.mutex.Unlock()
	return a
}
