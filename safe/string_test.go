package safe

import (
	"sync"
	"testing"
)

func TestString_Get(t *testing.T) {
	v := "hello"
	age := NewString(v)
	var wg sync.WaitGroup
	wg.Add(2)
	a := ""
	b := ""
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
		t.Fatalf("String.Get() = %s, wanted %s", a, v)
	}
	if b != v {
		t.Fatalf("String.Get() = %s, wanted %s", b, v)
	}
}

func TestString_Set(t *testing.T) {
	age := NewString("hello")
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		age.Set("foo")
		wg.Done()
	}()
	go func() {
		age.Set("foo")
		wg.Done()
	}()
	wg.Wait()
	a := age.Get()
	if a != "foo" {
		t.Fatalf(`String.Get() = "%s", wanted "foo"`, a)
	}
}

func TestString_SetFunc(t *testing.T) {
	age := NewString("hello")
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		age.SetFunc(func(v string) string {
			return v + "!"
		})
		wg.Done()
	}()
	go func() {
		age.SetFunc(func(v string) string {
			return v + "!!"
		})
		wg.Done()
	}()
	wg.Wait()
	a := age.Get()
	exp := "hello!!!" //nolint:goconst
	if a != exp {
		t.Fatalf(`String.Get() = "%s", wanted "%s"`, a, exp)
	}
}

func TestString_Add(t *testing.T) {
	age := NewString("hello")
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		age.Add("!")
		wg.Done()
	}()
	go func() {
		age.Add("!!")
		wg.Done()
	}()
	wg.Wait()
	a := age.Get()
	exp := "hello!!!"
	if a != exp {
		t.Fatalf(`String.Get() = "%s", wanted "%s"`, a, exp)
	}
}

func TestString_AddR(t *testing.T) {
	age := NewString("hello")
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		age.AddR("!")
		wg.Done()
	}()
	go func() {
		age.AddR("!!")
		wg.Done()
	}()
	wg.Wait()
	a := age.Get()
	exp := "hello!!!"
	if a != exp {
		t.Fatalf(`String.Get() = "%s", wanted "%s"`, a, exp)
	}
}

func BenchmarkString_Add(b *testing.B) {
	age := NewString("")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		age.Add("h")
	}
}

func BenchmarkString_AddR(b *testing.B) {
	age := NewString("")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		age.AddR("h")
	}
}
