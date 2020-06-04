package safe

// GetUnsafe gets a value without lock.
func (i *Int) GetUnsafe() int {
	return i.value
}

// SetUnsafe sets a value without lock.
func (i *Int) SetUnsafe(v int) {
	i.value = v
}

// AddUnsafe adds a value without lock.
func (i *Int) AddUnsafe(v int) {
	i.value += v
}

// SubUnsafe substitutes a value without lock.
func (i *Int) SubUnsafe(v int) {
	i.value -= v
}

// MulUnsafe multiplies a value without lock.
func (i *Int) MulUnsafe(v int) {
	i.value *= v
}

// DivUnsafe divides a value without lock.
func (i *Int) DivUnsafe(v int) {
	i.value /= v
}
