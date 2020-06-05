package safe

import (
	"testing"
)

func TestString_GetUnsafe(t *testing.T) {
	v := "hello"
	age := &String{value: v}
	a := age.GetUnsafe()
	if a != v {
		t.Fatalf("String.GetUnsafe() = %s, wanted %s", a, v)
	}
}

func TestString_SetUnsafe(t *testing.T) {
	v := "foo"
	age := &String{}
	age.SetUnsafe(v)
	if age.value != v {
		t.Fatalf("String.GetUnsafe() = %s, wanted %s", age.value, v)
	}
}

func TestString_AddUnsafe(t *testing.T) {
	v := " foo"
	age := &String{value: "hello"}
	age.AddUnsafe(v)
	exp := "hello foo"
	if age.value != exp {
		t.Fatalf("String.GetUnsafe() = %s, wanted %s", age.value, exp)
	}
}
