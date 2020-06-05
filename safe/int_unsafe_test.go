package safe

import (
	"testing"
)

func TestInt_GetUnsafe(t *testing.T) {
	v := 5
	age := &Int{value: v}
	a := age.GetUnsafe()
	if a != v {
		t.Fatalf("Int.GetUnsafe() = %d, wanted %d", a, v)
	}
}

func TestInt_SetUnsafe(t *testing.T) {
	v := 3
	age := &Int{}
	age.SetUnsafe(v)
	if age.value != v {
		t.Fatalf("Int.GetUnsafe() = %d, wanted %d", age.value, v)
	}
}

func TestInt_AddUnsafe(t *testing.T) {
	age := &Int{value: 1}
	age.AddUnsafe(3)
	exp := 4
	if age.value != exp {
		t.Fatalf("Int.GetUnsafe() = %d, wanted %d", age.value, exp)
	}
}

func TestInt_SubUnsafe(t *testing.T) {
	age := &Int{value: 4}
	age.SubUnsafe(1)
	exp := 3
	if age.value != exp {
		t.Fatalf("Int.GetUnsafe() = %d, wanted %d", age.value, exp)
	}
}

func TestInt_MulUnsafe(t *testing.T) {
	age := &Int{value: 4}
	age.MulUnsafe(2)
	exp := 8
	if age.value != exp {
		t.Fatalf("Int.GetUnsafe() = %d, wanted %d", age.value, exp)
	}
}

func TestInt_DivUnsafe(t *testing.T) {
	age := &Int{value: 10}
	age.DivUnsafe(2)
	exp := 5
	if age.value != exp {
		t.Fatalf("Int.GetUnsafe() = %d, wanted %d", age.value, exp)
	}
}
