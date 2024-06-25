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

import "sync"

// Cache is a cache by key K to value V.
//
// It also uses locking to block getting values that are under computation.
type Cache[K comparable, V any] struct {
	Store map[K]*result[V]
	lock  sync.RWMutex
}

func New[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		Store: make(map[K]*result[V]),
	}
}

type result[V any] struct {
	Value V
	err   error
	lock  sync.RWMutex
}

// GetOrAdd returns the cached value for the key,
// or calls getUncached to get a new value, and then caches the value.
// Calls for the same key are blocked until the computed value is available.
func (c *Cache[K, V]) GetOrAdd(key K, getUncached func(K) (V, error)) (V, error) {
	c.lock.Lock()
	res, ok := c.Store[key]
	if ok {
		res.lock.RLock()
		c.lock.Unlock()
		value, err := res.Value, res.err
		res.lock.RUnlock()
		return value, err
	}

	res = &result[V]{}
	c.Store[key] = res

	res.lock.Lock()
	c.lock.Unlock()
	defer res.lock.Unlock()

	res.Value, res.err = getUncached(key)

	return res.Value, res.err
}

// Values returns all values of the cache.
// This access is not locked.
// The caller has to ensure that no parallel modification to the map is done.
func (c *Cache[K, V]) Values() []V {
	var x []V
	for _, v := range c.Store {
		x = append(x, v.Value)
	}
	return x
}
