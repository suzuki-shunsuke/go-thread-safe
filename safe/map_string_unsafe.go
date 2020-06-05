package safe

// GetUnsafe gets a value from the map without lock.
func (m *MapString) GetUnsafe(k string) string {
	return m.value[k]
}

// GetOkUnsafe gets a value from the map without lock.
func (m *MapString) GetOkUnsafe(k string) (string, bool) {
	v, ok := m.value[k]
	return v, ok
}

// HasUnsafe checks whether the map has the key without lock.
func (m *MapString) HasUnsafe(k string) bool {
	_, ok := m.value[k]
	return ok
}

// LenUnsafe gets the length of the map without lock.
func (m *MapString) LenUnsafe() int {
	return len(m.value)
}

// DeleteUnsafe deletes the key from the map without lock.
func (m *MapString) DeleteUnsafe(k string) {
	delete(m.value, k)
}

// SetUnsafe sets the key and value to the map without lock.
func (m *MapString) SetUnsafe(k, v string) {
	m.value[k] = v
}

// SetDefaultUnsafe sets the key and value to the map if the map doesn't have the key without lock.
func (m *MapString) SetDefaultUnsafe(k, v string) {
	if _, ok := m.value[k]; !ok {
		m.value[k] = v
	}
}

// RangeUnsafe gets all pairs of the key and value from the map and call the function without lock.
func (m *MapString) RangeUnsafe(f func(k, v string)) {
	for k, v := range m.value {
		f(k, v)
	}
}

// RangeBUnsafe gets pairs of the key and value from the map and call the function without lock.
// If the function returns false, the loop ends.
func (m *MapString) RangeBUnsafe(f func(k, v string) bool) {
	for k, v := range m.value {
		if !f(k, v) {
			break
		}
	}
}
