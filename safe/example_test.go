package safe_test

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/suzuki-shunsuke/go-thread-safe/safe"
)

func Example() {
	dest := safe.NewMapString(map[string]string{})

	data := map[string]int{
		"foo": 1,
		"zoo": 2,
		"bar": 3,
	}

	var wg sync.WaitGroup
	for k := range data {
		k := k
		wg.Add(1)
		go func() {
			defer wg.Done()
			// dest is updated by multiple go routine in parallel.
			// MapString.Set is thread safe.
			dest.Set(k, k+" world")
		}()
	}
	wg.Wait()
	// MapString.Range is thread safe.
	dest.Range(func(k, v string) {
		fmt.Printf("%s: %s\n", k, v)
	})
	// Unordered output:
	// foo: foo world
	// bar: bar world
	// zoo: zoo world
}

func ExampleBool() {
	flag := &safe.Bool{}

	data := map[string]int{
		"foo": 1,
		"zoo": 2,
		"bar": 3,
	}

	var wg sync.WaitGroup
	for _, v := range data {
		v := v
		wg.Add(1)
		go func() {
			// flag is updated by multiple go routines in parallel.
			defer wg.Done()
			if v%2 == 1 {
				// flag.Set is thread safe.
				flag.Set(true)
			}
		}()
	}
	wg.Wait()
	// Note that GetUnsafe isn't thread safe, so we should use this method carefully.
	// GetUnsafe is faster than Get.
	fmt.Printf("flag: %t\n", flag.GetUnsafe())
	// Output:
	// flag: true
}

func ExampleInt() {
	age := &safe.Int{}
	age.Set(5)

	data := map[string]int{
		"foo": 1,
		"zoo": 2,
		"bar": 3,
	}

	var wg sync.WaitGroup
	for _, v := range data {
		v := v
		wg.Add(1)
		go func() {
			// age is updated by multiple go routines in parallel.
			defer wg.Done()
			// Int.Add is thread safe.
			age.Add(v)
		}()
	}
	wg.Wait()
	fmt.Printf("age: %d\n", age.Get())
	// Note that GetUnsafe isn't thread safe, so we should use this method carefully.
	// GetUnsafe is faster than Get.
	fmt.Printf("age: %d\n", age.GetUnsafe())
	// Output:
	// age: 11
	// age: 11
}

func ExampleString() {
	name := &safe.String{}
	name.Set("foo")

	data := map[string]int{
		"foo": 1,
		"zoo": 2,
		"bar": 3,
	}

	var wg sync.WaitGroup
	for k, v := range data {
		k := k
		v := v
		wg.Add(1)
		go func() {
			// age is updated by multiple go routines in parallel.
			defer wg.Done()
			// String.Get is thread safe.
			if name.Get() == k {
				fmt.Println(v)
				name.Set(strconv.Itoa(v))
			}
		}()
	}
	wg.Wait()
	// Output:
	// 1
}

func ExampleMapString() {
	dest := safe.NewMapString(map[string]string{})

	data := map[string]int{
		"foo": 1,
		"zoo": 2,
		"bar": 3,
	}

	var wg sync.WaitGroup
	for k := range data {
		k := k
		wg.Add(1)
		go func() {
			defer wg.Done()
			// dest is updated by multiple go routine in parallel.
			// MapString.Set is thread safe.
			dest.Set(k, k+" world")
		}()
	}
	wg.Wait()
	// MapString.Range is thread safe.
	dest.Range(func(k, v string) {
		fmt.Printf("%s: %s\n", k, v)
	})
	// Unordered output:
	// foo: foo world
	// bar: bar world
	// zoo: zoo world
}
