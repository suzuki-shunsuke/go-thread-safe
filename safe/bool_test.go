package safe

import (
	"sync"
	"testing"
)

func TestBool_Get(t *testing.T) {
	v := true
	age := NewBool(v)
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
	age := NewBool(false)
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
	age := NewBool(true)
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

func TestBool_SetFuncR(t *testing.T) {
	age := NewBool(true)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		age.SetFuncR(func(v bool) bool {
			return !v
		})
		wg.Done()
	}()
	go func() {
		age.SetFuncR(func(v bool) bool {
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
	age := NewBool(false)
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
	age := NewBool(false)
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
