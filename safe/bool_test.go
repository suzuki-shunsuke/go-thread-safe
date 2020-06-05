package safe

import (
	"encoding/json"
	"sync"
	"testing"
)

func TestBool_String(t *testing.T) {
	age := &Bool{}
	var wg sync.WaitGroup
	wg.Add(2)
	a := ""
	go func() {
		age.Set(true)
		wg.Done()
	}()
	go func() {
		a = age.String()
		wg.Done()
	}()
	wg.Wait()
	a = age.String()
	exp := "Bool{true}"
	if a != exp {
		t.Fatalf("Bool.String() = %s, wanted %s", a, exp)
	}
}

func TestBool_MarshalJSON(t *testing.T) {
	age := &Bool{}
	var wg sync.WaitGroup
	wg.Add(2)
	var err error
	go func() {
		age.Set(true)
		wg.Done()
	}()
	go func() {
		_, err = json.Marshal(age)
		wg.Done()
	}()
	wg.Wait()

	if err != nil {
		t.Fatal(err)
	}
	b, err := json.Marshal(age)
	if err != nil {
		t.Fatal(err)
	}
	exp := "true"

	if string(b) != exp {
		t.Fatalf("Bool.MarshalJSON() = %s, wanted %s", string(b), exp)
	}
}

func TestBool_UnmarshalJSON(t *testing.T) {
	age := &Bool{}
	var wg sync.WaitGroup
	buf := []byte("true")
	wg.Add(2)
	var err error
	go func() {
		age.Set(false)
		wg.Done()
	}()
	go func() {
		err = json.Unmarshal(buf, age)
		wg.Done()
	}()
	wg.Wait()

	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(buf, age); err != nil {
		t.Fatal(err)
	}
	exp := true

	if age.value != exp {
		t.Fatalf("Bool.UnmarshalJSON() = %t, wanted %t", age.value, exp)
	}
}

func TestBool_Get(t *testing.T) {
	v := true
	age := &Bool{value: v}
	var wg sync.WaitGroup
	wg.Add(2)
	a := false
	b := false
	go func() {
		a = age.Get()
		wg.Done()
	}()
	go func() {
		b = age.Get()
		wg.Done()
	}()
	wg.Wait()
	if a != v {
		t.Fatalf("Bool.Get() = %t, wanted %t", a, v)
	}
	if b != v {
		t.Fatalf("Bool.Get() = %t, wanted %t", b, v)
	}
}

func TestBool_Set(t *testing.T) {
	age := &Bool{}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		age.Set(true)
		wg.Done()
	}()
	go func() {
		age.Set(true)
		wg.Done()
	}()
	wg.Wait()
	a := age.Get()
	exp := true
	if a != exp {
		t.Fatalf("Bool.Get() = %t, wanted %t", a, exp)
	}
}

func TestBool_SetFunc(t *testing.T) {
	age := &Bool{value: true}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		age.SetFunc(func(v bool) bool {
			return !v
		})
		wg.Done()
	}()
	go func() {
		age.SetFunc(func(v bool) bool {
			return !v
		})
		wg.Done()
	}()
	wg.Wait()
	a := age.Get()
	exp := true
	if a != exp {
		t.Fatalf(`Bool.Get() = %t, wanted %t`, a, exp)
	}
}

func TestBool_Invert(t *testing.T) {
	age := &Bool{}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		age.Invert()
		wg.Done()
	}()
	go func() {
		age.Invert()
		wg.Done()
	}()
	wg.Wait()
	a := age.Get()
	exp := false
	if a != exp {
		t.Fatalf(`Bool.Get() = %t, wanted %t`, a, exp)
	}
}

func TestBool_InvertR(t *testing.T) {
	age := &Bool{}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		age.InvertR()
		wg.Done()
	}()
	go func() {
		age.InvertR()
		wg.Done()
	}()
	wg.Wait()
	a := age.Get()
	exp := false
	if a != exp {
		t.Fatalf(`Bool.Get() = %t, wanted %t`, a, exp)
	}
}
