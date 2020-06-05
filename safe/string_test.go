package safe

import (
	"encoding/json"
	"sync"
	"testing"
)

func TestString_String(t *testing.T) {
	age := &String{}
	var wg sync.WaitGroup
	wg.Add(2)
	a := ""
	go func() {
		age.Set("foo")
		wg.Done()
	}()
	go func() {
		a = age.String()
		wg.Done()
	}()
	wg.Wait()
	a = age.String()
	exp := "String{foo}"
	if a != exp {
		t.Fatalf("String.String() = %s, wanted %s", a, exp)
	}
}

func TestString_MarshalJSON(t *testing.T) {
	age := &String{}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		age.Set("foo")
		wg.Done()
	}()
	var err error
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
	exp := `"foo"`
	if string(b) != exp {
		t.Fatalf("String.String() = %s, wanted %s", string(b), exp)
	}
}

func TestString_UnmarshalJSON(t *testing.T) {
	age := &String{}
	var wg sync.WaitGroup
	buf := []byte(`"hello"`)
	wg.Add(2)
	go func() {
		age.Set("foo")
		wg.Done()
	}()
	var err error
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
	exp := "hello" //nolint:goconst
	if age.value != exp {
		t.Fatalf("String.String() = %s, wanted %s", age.value, exp)
	}
}

func TestString_Get(t *testing.T) {
	v := "hello"
	age := &String{value: v}
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
	age := &String{value: "hello"}
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
	age := &String{value: "hello"}
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
	age := &String{value: "hello"}
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
	age := &String{value: "hello"}
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
	age := &String{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		age.Add("h")
	}
}

func BenchmarkString_AddR(b *testing.B) {
	age := &String{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		age.AddR("h")
	}
}
