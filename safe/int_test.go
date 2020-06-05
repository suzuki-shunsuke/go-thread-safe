package safe

import (
	"encoding/json"
	"sync"
	"testing"
)

func TestInt_String(t *testing.T) {
	age := &Int{}
	var wg sync.WaitGroup
	wg.Add(2)
	a := ""
	go func() {
		age.Set(5)
		wg.Done()
	}()
	go func() {
		a = age.String()
		wg.Done()
	}()
	wg.Wait()
	a = age.String()
	exp := "Int{5}"
	if a != exp {
		t.Fatalf("Int.String() = %s, wanted %s", a, exp)
	}
}

func TestInt_MarshalJSON(t *testing.T) {
	age := &Int{}
	var wg sync.WaitGroup
	wg.Add(2)
	var err error
	go func() {
		age.Set(5)
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
	exp := "5"

	if string(b) != exp {
		t.Fatalf("Int.MarshalJSON() = %s, wanted %s", string(b), exp)
	}
}

func TestInt_UnmarshalJSON(t *testing.T) {
	age := &Int{}
	var wg sync.WaitGroup
	buf := []byte("10")
	wg.Add(2)
	var err error
	go func() {
		age.Set(5)
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
	exp := 10

	if age.value != exp {
		t.Fatalf("Int.UnmarshalJSON() = %d, wanted %d", age.value, exp)
	}
}

func TestInt_Get(t *testing.T) {
	v := 5
	age := &Int{value: v}
	var wg sync.WaitGroup
	wg.Add(2)
	a := 0
	b := 0
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
		t.Fatalf("Int.Get() = %d, wanted %d", a, v)
	}
	if b != v {
		t.Fatalf("Int.Get() = %d, wanted %d", b, v)
	}
}

func TestInt_Set(t *testing.T) {
	age := &Int{value: 5}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		age.Set(2)
		wg.Done()
	}()
	go func() {
		age.Set(2)
		wg.Done()
	}()
	wg.Wait()
	a := age.Get()
	if a != 2 {
		t.Fatalf("Int.Get() = %d, wanted 2", a)
	}
}

func TestInt_SetFunc(t *testing.T) {
	age := &Int{value: 5}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		age.SetFunc(func(v int) int {
			return v + 1
		})
		wg.Done()
	}()
	go func() {
		age.SetFunc(func(v int) int {
			return v + 2
		})
		wg.Done()
	}()
	wg.Wait()
	a := age.Get()
	if a != 8 {
		t.Fatalf("Int.Get() = %d, wanted 8", a)
	}
}

func TestInt_Add(t *testing.T) {
	age := &Int{value: 5}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		age.Add(1)
		wg.Done()
	}()
	go func() {
		age.Add(2)
		wg.Done()
	}()
	wg.Wait()
	a := age.Get()
	if a != 8 {
		t.Fatalf("Int.Get() = %d, wanted 8", a)
	}
}

func TestInt_AddR(t *testing.T) {
	age := &Int{value: 5}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		age.AddR(1)
		wg.Done()
	}()
	go func() {
		age.AddR(2)
		wg.Done()
	}()
	wg.Wait()
	a := age.Get()
	if a != 8 {
		t.Fatalf("Int.Get() = %d, wanted 8", a)
	}
}

func TestInt_Sub(t *testing.T) {
	age := &Int{value: 5}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		age.Sub(1)
		wg.Done()
	}()
	go func() {
		age.Sub(2)
		wg.Done()
	}()
	wg.Wait()
	a := age.Get()
	if a != 2 {
		t.Fatalf("Int.Get() = %d, wanted 2", a)
	}
}

func TestInt_SubR(t *testing.T) {
	age := &Int{value: 5}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		age.SubR(1)
		wg.Done()
	}()
	go func() {
		age.SubR(2)
		wg.Done()
	}()
	wg.Wait()
	a := age.Get()
	if a != 2 {
		t.Fatalf("Int.Get() = %d, wanted 2", a)
	}
}

func TestInt_Mul(t *testing.T) {
	age := &Int{value: 5}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		age.Mul(2)
		wg.Done()
	}()
	go func() {
		age.Mul(3)
		wg.Done()
	}()
	wg.Wait()
	a := age.Get()
	if a != 30 {
		t.Fatalf("Int.Get() = %d, wanted 30", a)
	}
}

func TestInt_MulR(t *testing.T) {
	age := &Int{value: 5}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		age.MulR(2)
		wg.Done()
	}()
	go func() {
		age.MulR(3)
		wg.Done()
	}()
	wg.Wait()
	a := age.Get()
	if a != 30 {
		t.Fatalf("Int.Get() = %d, wanted 30", a)
	}
}

func TestInt_Div(t *testing.T) {
	age := &Int{value: 20}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		age.Div(2)
		wg.Done()
	}()
	go func() {
		age.Div(5)
		wg.Done()
	}()
	wg.Wait()
	a := age.Get()
	if a != 2 {
		t.Fatalf("Int.Get() = %d, wanted 2", a)
	}
}

func TestInt_DivR(t *testing.T) {
	age := &Int{value: 20}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		age.DivR(2)
		wg.Done()
	}()
	go func() {
		age.DivR(5)
		wg.Done()
	}()
	wg.Wait()
	a := age.Get()
	if a != 2 {
		t.Fatalf("Int.Get() = %d, wanted 2", a)
	}
}

func BenchmarkInt_Add(b *testing.B) {
	age := &Int{value: 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		age.Add(1)
	}
}

func BenchmarkInt_AddR(b *testing.B) {
	age := &Int{value: 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		age.AddR(1)
	}
}
