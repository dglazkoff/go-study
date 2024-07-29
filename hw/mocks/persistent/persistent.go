package persistent

import "gopl.io/hw/mocks/store"

func Lookup(s store.Store, key string) ([]byte, error) {
	// ...
	return s.Get(key)
}
