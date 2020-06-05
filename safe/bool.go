package safe

import (
	"encoding/json"
	"strconv"
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

func (b *Bool) String() string {
	b.mutex.RLock()
	v := b.value
	b.mutex.RUnlock()
	return "Bool{" + strconv.FormatBool(v) + "}"
}

func (b *Bool) MarshalJSON() ([]byte, error) {
	b.mutex.RLock()
	v := b.value
	b.mutex.RUnlock()
	return json.Marshal(v)
}

func (b *Bool) UnmarshalJSON(buf []byte) error {
	b.mutex.Lock()
	err := json.Unmarshal(buf, &b.value)
	b.mutex.Unlock()
	return err
}

func (b *Bool) Get() bool {
	b.mutex.RLock()
	v := b.value
	b.mutex.RUnlock()
	return v
}

func (b *Bool) Set(v bool) {
	b.mutex.Lock()
	b.value = v
	b.mutex.Unlock()
}

func (b *Bool) SetFunc(f func(v bool) bool) {
	b.mutex.Lock()
	b.value = f(b.value)
	b.mutex.Unlock()
}

func (b *Bool) Invert() {
	b.mutex.Lock()
	b.value = !b.value
	b.mutex.Unlock()
}

func (b *Bool) InvertR() bool {
	b.mutex.Lock()
	a := !b.value
	b.value = a
	b.mutex.Unlock()
	return a
}
