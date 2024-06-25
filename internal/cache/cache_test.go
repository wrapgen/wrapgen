// Copyright 2024 Wrapgen authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cache

import (
	"context"
	"fmt"
	"slices"
	"sort"
	"sync"
	"testing"
)

func TestCache_GetOrAdd(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	m := new(sync.Map)
	c := New[int, struct{}]()
	getter := func(key int) (struct{}, error) {
		_, ok := m.Load(key)
		if ok {
			cancel(fmt.Errorf("this value was already computed: %v", key))
			return struct{}{}, nil
		}
		m.Store(key, struct{}{})
		return struct{}{}, nil
	}
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10_000; i++ {
				_, err := c.GetOrAdd(i, getter)
				if err != nil {
					cancel(fmt.Errorf("unexpected error: %w", err))
				}
			}
		}()
	}

	wg.Wait()
	err := context.Cause(ctx)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestCache_Values(t *testing.T) {
	c := New[int, int]()
	for i := 0; i < 10; i++ {
		v, err := c.GetOrAdd(i, func(i int) (int, error) { return i, nil })
		if v != i {
			t.Fatalf("wrong value %v expected %v", v, i)
		}
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
	}

	v := c.Values()
	sort.Ints(v)
	if !slices.Equal(v, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Fatalf("wrong values: %v", v)
	}
}
