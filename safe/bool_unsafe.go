package safe

// GetUnsafe gets a value without lock.
func (s *Bool) GetUnsafe() bool {
	return s.value
}

// SetUnsafe sets a value without lock.
func (s *Bool) SetUnsafe(v bool) {
	s.value = v
}
