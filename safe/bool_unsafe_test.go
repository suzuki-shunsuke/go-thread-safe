package safe

import (
	"testing"
)

func TestBool_GetUnsafe(t *testing.T) {
	v := true
	age := NewBool(v)
	a := age.GetUnsafe()
	if a != v {
		t.Fatalf("Bool.GetUnsafe() = %t, wanted %t", a, v)
	}
}

func TestBool_SetUnsafe(t *testing.T) {
	v := true
	age := NewBool(false)
	age.SetUnsafe(v)
	if age.value != v {
		t.Fatalf("Bool.GetUnsafe() = %t, wanted %t", age.value, v)
	}
}
