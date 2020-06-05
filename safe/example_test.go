package safe_test

import (
	"fmt"
	"sync"

	"github.com/suzuki-shunsuke/go-thread-safe/safe"
)

func Example() {
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
