package safe

import (
	"testing"
)

func TestBool_GetUnsafe(t *testing.T) {
	v := true
	age := &Bool{value: v}
	a := age.GetUnsafe()
	if a != v {
		t.Fatalf("Bool.GetUnsafe() = %t, wanted %t", a, v)
	}
}

func TestBool_SetUnsafe(t *testing.T) {
	v := true
	age := &Bool{}
	age.SetUnsafe(v)
	if age.value != v {
		t.Fatalf("Bool.GetUnsafe() = %t, wanted %t", age.value, v)
	}
}
