package safe

func (s *String) GetUnsafe() string {
	return s.value
}

func (s *String) SetUnsafe(v string) {
	s.value = v
}

func (s *String) AddUnsafe(v string) {
	s.value += v
}
