package safe

import (
	"testing"
)

func TestMapString_GetUnsafe(t *testing.T) {
	age := NewMapString(map[string]string{"foo": "bar"}, 1)
	a := age.GetUnsafe("foo")
	exp := "bar"
	if a != exp {
		t.Fatalf("MapString.GetUnsafe() = %s, wanted %s", a, exp)
	}
}

func TestMapString_GetOkUnsafe(t *testing.T) {
	age := NewMapString(map[string]string{"foo": "bar"}, 1)
	a, ok := age.GetOkUnsafe("foo")
	exp := "bar"
	if a != exp {
		t.Fatalf("MapString.GetOkUnsafe() = %s, wanted %s", a, exp)
	}
	if !ok {
		t.Fatalf("MapString.GetOkUnsafe() = _, %t, wanted %t", ok, false)
	}
}

func TestMapString_HasUnsafe(t *testing.T) {
	age := NewMapString(map[string]string{"foo": "bar"}, 1)
	ok := age.HasUnsafe("foo")
	if !ok {
		t.Fatalf("MapString.HasUnsafe() = %t, wanted %t", ok, false)
	}
}

func TestMapString_LenUnsafe(t *testing.T) {
	age := NewMapString(map[string]string{"foo": "bar"}, 1)
	a := age.LenUnsafe()
	if a != 1 {
		t.Fatalf("MapString.LenUnsafe() = %d, wanted %d", a, 1)
	}
}

func TestMapString_DeleteUnsafe(t *testing.T) {
	age := NewMapString(map[string]string{"foo": "bar"}, 1)
	age.DeleteUnsafe("foo")
	a := len(age.value)
	if a != 0 {
		t.Fatalf("len(MapString.value) = %d, wanted %d", a, 0)
	}
}

func TestMapString_SetUnsafe(t *testing.T) {
	age := NewMapString(map[string]string{"foo": "bar"}, 1)
	exp := "zoo"
	age.SetUnsafe("foo", exp)
	a := age.value["foo"]
	if a != exp {
		t.Fatalf("age.value['foo'] = %s, wanted %s", a, exp)
	}
}

func TestMapString_SetDefaultUnsafe(t *testing.T) {
	exp := "bar"
	key := "foo"
	age := NewMapString(map[string]string{key: exp}, 1)
	age.SetDefaultUnsafe(key, "zoo")
	a := age.value[key]
	if a != exp {
		t.Fatalf("age.value['foo'] = %s, wanted %s", a, exp)
	}

	age.SetDefaultUnsafe("zoo", "goo")
	a = age.value["zoo"]
	if a != "goo" {
		t.Fatalf("age.value['zoo'] = %s, wanted %s", a, "goo")
	}
}

func TestMapString_SetDefaultRUnsafe(t *testing.T) {
	exp := "bar"
	key := "foo"
	age := NewMapString(map[string]string{key: exp}, 1)
	age.SetDefaultRUnsafe(key, "zoo")
	a := age.value[key]
	if a != exp {
		t.Fatalf("age.value['foo'] = %s, wanted %s", a, exp)
	}

	age.SetDefaultRUnsafe("zoo", "goo")
	a = age.value["zoo"]
	if a != "goo" {
		t.Fatalf("age.value['zoo'] = %s, wanted %s", a, "goo")
	}
}

func TestMapString_RangeUnsafe(t *testing.T) {
	age := NewMapString(map[string]string{"foo": "bar"}, 1)
	age.RangeUnsafe(func(k, v string) {
		if k != "foo" {
			t.Fatalf("k = %s, wanted %s", k, "foo")
		}
		if v != "bar" {
			t.Fatalf("v = %s, wanted %s", v, "bar")
		}
	})
}

func TestMapString_RangeBUnsafe(t *testing.T) {
	age := NewMapString(map[string]string{"foo": "bar"}, 1)
	age.RangeBUnsafe(func(k, v string) bool {
		if k != "foo" {
			t.Fatalf("k = %s, wanted %s", k, "foo")
		}
		if v != "bar" {
			t.Fatalf("v = %s, wanted %s", v, "bar")
		}
		return false
	})
}

func TestMapString_CopyUnsafe(t *testing.T) {
	age := NewMapString(map[string]string{"foo": "bar"}, 1)
	cp := age.CopyUnsafe()

	exp := 1
	a := len(cp.value)
	if a != exp {
		t.Fatalf("len(cp.value) = %d, wanted %d", a, exp)
	}
}

func TestMapString_CopyDataUnsafe(t *testing.T) {
	age := NewMapString(map[string]string{"foo": "bar"}, 1)
	cp := make(map[string]string, 1)
	age.CopyDataUnsafe(cp)

	exp := 1
	a := len(cp)
	if a != exp {
		t.Fatalf("len(cp.value) = %d, wanted %d", a, exp)
	}
}
