package safe

import (
	"encoding/json"
	"strings"
	"sync"
	"testing"
)

func TestMapString_String(t *testing.T) {
	age := NewMapString(map[string]string{"foo": "bar"})
	var wg sync.WaitGroup
	wg.Add(2)
	a := ""
	go func() {
		age.Set("zoo", "world")
		wg.Done()
	}()
	go func() {
		a = age.String()
		wg.Done()
	}()
	wg.Wait()
	a = age.String()
	if !strings.HasPrefix(a, "MapString") {
		t.Fatalf("MapString.String() = %s, must start with 'MapString'", a)
	}
}

func TestMapString_MarshalJSON(t *testing.T) {
	age := NewMapString(map[string]string{"foo": "bar"})
	var wg sync.WaitGroup
	wg.Add(2)
	var err error
	go func() {
		age.Set("foo", "world")
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
	exp := `{"foo":"world"}`

	if string(b) != exp {
		t.Fatalf("MapString.MarshalJSON() = %s, wanted %s", string(b), exp)
	}
}

func TestMapString_UnmarshalJSON(t *testing.T) {
	age := NewMapString(map[string]string{"foo": "bar"})
	var wg sync.WaitGroup
	buf := []byte(`{"hello":"world"}`)
	wg.Add(2)
	var err error
	go func() {
		age.Set("zoo", "world")
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
	if len(age.value) != 3 {
		t.Fatalf("len(age.value) = %d, wanted 3", len(age.value))
	}
	exp := "world" //nolint:goconst
	if age.value["hello"] != exp {
		t.Fatalf(`age.value["hello"] = %s, wanted "%s"`, age.value["hello"], exp)
	}
}

func TestMapString_Get(t *testing.T) {
	age := NewMapString(map[string]string{"foo": "bar"})

	var wg sync.WaitGroup
	wg.Add(2)
	key := "foo" //nolint:goconst
	go func() {
		age.Get(key)
		wg.Done()
	}()
	go func() {
		age.Set(key, "zoo")
		wg.Done()
	}()
	wg.Wait()

	a := age.Get(key)

	exp := "zoo" //nolint:goconst
	if a != exp {
		t.Fatalf("MapString.Get() = %s, wanted %s", a, exp)
	}
}

func TestMapString_GetOk(t *testing.T) {
	age := NewMapString(map[string]string{"foo": "bar"})

	var wg sync.WaitGroup
	wg.Add(2)
	key := "foo"
	go func() {
		age.GetOk(key)
		wg.Done()
	}()
	go func() {
		age.Set(key, "zoo")
		wg.Done()
	}()
	wg.Wait()

	a, ok := age.GetOk(key)

	exp := "zoo"
	if a != exp {
		t.Fatalf("MapString.GetOk() = %s, wanted %s", a, exp)
	}
	if !ok {
		t.Fatalf("MapString.GetOk() = _, %t, wanted %t", ok, false)
	}
}

func TestMapString_Has(t *testing.T) {
	age := NewMapString(map[string]string{"foo": "bar"})
	key := "foo"

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		age.Has(key)
		wg.Done()
	}()
	go func() {
		age.Set(key, "zoo")
		wg.Done()
	}()
	wg.Wait()

	ok := age.Has(key)

	if !ok {
		t.Fatalf("MapString.Has() = %t, wanted %t", ok, false)
	}
}

func TestMapString_Len(t *testing.T) {
	age := NewMapString(map[string]string{"foo": "bar"})

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		age.Len()
		wg.Done()
	}()
	go func() {
		age.Set("foo", "zoo")
		wg.Done()
	}()
	wg.Wait()

	a := age.Len()
	if a != 1 {
		t.Fatalf("MapString.Len() = %d, wanted %d", a, 1)
	}
}

func TestMapString_Delete(t *testing.T) {
	age := NewMapString(map[string]string{"foo": "bar"})

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		age.Delete("foo")
		wg.Done()
	}()
	go func() {
		age.Delete("foo")
		wg.Done()
	}()
	wg.Wait()

	a := len(age.value)
	if a != 0 {
		t.Fatalf("len(MapString.value) = %d, wanted %d", a, 0)
	}
}

func TestMapString_DeleteR(t *testing.T) {
	exp := "bar" //nolint:goconst
	age := NewMapString(map[string]string{"foo": "bar"})

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		age.DeleteR("zoo")
		wg.Done()
	}()
	go func() {
		age.DeleteR("zoo")
		wg.Done()
	}()
	wg.Wait()

	a := age.DeleteR("foo")
	if a != exp {
		t.Fatalf("a = %s, wanted %s", a, exp)
	}
}

func TestMapString_DeleteROk(t *testing.T) {
	exp := "bar"
	age := NewMapString(map[string]string{"foo": "bar"})

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		age.DeleteROk("zoo")
		wg.Done()
	}()
	go func() {
		age.DeleteROk("zoo")
		wg.Done()
	}()
	wg.Wait()

	a, ok := age.DeleteROk("foo")
	if a != exp {
		t.Fatalf("a = %s, wanted %s", a, exp)
	}
	if !ok {
		t.Fatalf("ok = %t, wanted %t", ok, true)
	}
}

func TestMapString_Set(t *testing.T) {
	age := NewMapString(map[string]string{"foo": "bar"})
	exp := "zoo"

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		age.Set("foo", exp)
		wg.Done()
	}()
	go func() {
		age.Set("foo", exp)
		wg.Done()
	}()
	wg.Wait()

	a := age.value["foo"]
	if a != exp {
		t.Fatalf("age.value['foo'] = %s, wanted %s", a, exp)
	}
}

func TestMapString_SetFunc(t *testing.T) {
	age := NewMapString(map[string]string{"foo": "bar"})
	age.SetFunc("foo", func(v string, ok bool) string {
		return v + " world"
	})
	a := age.value["foo"]
	exp := "bar world"
	if a != exp {
		t.Fatalf("age.value['foo'] = %s, wanted %s", a, exp)
	}
}

func TestMapString_Range(t *testing.T) {
	age := NewMapString(map[string]string{"foo": "bar"})
	age.Range(func(k, v string) {
		if k != "foo" {
			t.Fatalf("k = %s, wanted %s", k, "foo")
		}
		if v != "bar" {
			t.Fatalf("v = %s, wanted %s", v, "bar")
		}
	})
	age.Range(func(k, v string) {
		if k != "foo" {
			t.Fatalf("k = %s, wanted %s", k, "foo")
		}
		if v != "bar" {
			t.Fatalf("v = %s, wanted %s", v, "bar")
		}
	})
}

func TestMapString_RangeB(t *testing.T) {
	age := NewMapString(map[string]string{"foo": "bar"})
	age.RangeB(func(k, v string) bool {
		if k != "foo" {
			t.Fatalf("k = %s, wanted %s", k, "foo")
		}
		if v != "bar" {
			t.Fatalf("v = %s, wanted %s", v, "bar")
		}
		return false
	})
	age.RangeB(func(k, v string) bool {
		if k != "foo" {
			t.Fatalf("k = %s, wanted %s", k, "foo")
		}
		if v != "bar" {
			t.Fatalf("v = %s, wanted %s", v, "bar")
		}
		return false
	})
}

func TestMapString_Copy(t *testing.T) {
	age := NewMapString(map[string]string{"foo": "bar"})

	var wg sync.WaitGroup
	cp := NewMapString(map[string]string{})
	wg.Add(2)
	go func() {
		age.Copy(cp)
		wg.Done()
	}()
	go func() {
		age.Set("foo", "bar")
		wg.Done()
	}()
	wg.Wait()

	exp := 1
	a := len(cp.value)
	if a != exp {
		t.Fatalf("len(cp.value) = %d, wanted %d", a, exp)
	}
}

func TestMapString_CopyData(t *testing.T) {
	age := NewMapString(map[string]string{"foo": "bar"})

	var wg sync.WaitGroup
	cp := map[string]string{}
	wg.Add(2)
	go func() {
		age.CopyData(cp)
		wg.Done()
	}()
	go func() {
		age.Set("foo", "bar")
		wg.Done()
	}()
	wg.Wait()

	exp := 1
	a := len(cp)
	if a != exp {
		t.Fatalf("len(cp.value) = %d, wanted %d", a, exp)
	}
}

func TestMapString_SetDefault(t *testing.T) {
	age := NewMapString(map[string]string{"foo": "bar"})

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		age.SetDefault("foo", "zoo")
		wg.Done()
	}()
	go func() {
		age.SetDefault("bar", "world")
		wg.Done()
	}()
	wg.Wait()

	exp := "bar"
	a := age.GetUnsafe("foo")
	if a != exp {
		t.Fatalf(`age.GetUnsafe("foo") = %s, wanted %s`, a, exp)
	}

	exp = "world"
	a = age.GetUnsafe("bar")
	if a != exp {
		t.Fatalf(`age.GetUnsafe("bar") = %s, wanted %s`, a, exp)
	}
}

func TestMapString_SetDefaultR(t *testing.T) {
	age := NewMapString(map[string]string{"foo": "bar"})

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		age.SetDefaultR("foo", "zoo")
		wg.Done()
	}()
	go func() {
		age.SetDefaultR("bar", "world")
		wg.Done()
	}()
	wg.Wait()

	exp := "bar"
	a := age.GetUnsafe("foo")
	if a != exp {
		t.Fatalf(`age.GetUnsafe("foo") = %s, wanted %s`, a, exp)
	}

	exp = "world"
	a = age.GetUnsafe("bar")
	if a != exp {
		t.Fatalf(`age.GetUnsafe("bar") = %s, wanted %s`, a, exp)
	}
}

func BenchmarkMapString_Set(b *testing.B) {
	key := "foo"
	age := NewMapString(map[string]string{key: "bar"})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a := age.Get(key)
		age.Set(key, a+" world")
	}
}

func BenchmarkMapString_SetFunc(b *testing.B) {
	key := "foo"
	age := NewMapString(map[string]string{key: "bar"})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		age.SetFunc(key, func(v string, ok bool) string {
			return v + " world"
		})
	}
}

func BenchmarkNewMapString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		age := NewMapString(map[string]string{"foo": "bar"})
		age.LenUnsafe()
	}
}
