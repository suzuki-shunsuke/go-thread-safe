package safe

// GetUnsafe gets a value without lock.
func (b *Bool) GetUnsafe() bool {
	return b.value
}

// SetUnsafe sets a value without lock.
func (b *Bool) SetUnsafe(v bool) {
	b.value = v
}
