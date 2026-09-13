package iplimiter

import "sync"

type typedSyncMap[K comparable, V any] struct {
	m sync.Map
}

func (s *typedSyncMap[K, V]) Load(key K) (V, bool) {
	v, ok := s.m.Load(key)
	if !ok {
		var zero V
		return zero, false
	}

	return v.(V), true //nolint:errcheck // err checking not needed in actual realization
}

func (s *typedSyncMap[K, V]) LoadOrStore(key K, value V) (V, bool) {
	a, loaded := s.m.LoadOrStore(key, value)
	return a.(V), loaded //nolint:errcheck // err checking not needed in actual realization
}

func (s *typedSyncMap[K, V]) Range(f func(K, V) bool) {
	s.m.Range(func(key, value any) bool {
		return f(key.(K), value.(V)) //nolint:errcheck // err checking not needed in actual realization
	})
}

func (s *typedSyncMap[K, V]) Delete(key K) {
	s.m.Delete(key)
}
