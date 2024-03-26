package set

import (
	"sync"
)

// Set - потокобезопасное множество
type Set struct {
	lock sync.RWMutex
	data map[interface{}]struct{}
}

// New создает новое множество
func New() *Set {
	return &Set{
		data: make(map[interface{}]struct{}),
	}
}

// Add добавляет элемент в множество
func (s *Set) Add(item interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.data[item] = struct{}{}
}

// Remove удаляет элемент из множества
func (s *Set) Remove(item interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.data, item)
}

// Exists проверяет, существует ли элемент в множестве
func (s *Set) Exists(item interface{}) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	_, exists := s.data[item]
	return exists
}

// List возвращает список всех элементов множества
func (s *Set) List() []interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()
	items := make([]interface{}, 0, len(s.data))
	for item := range s.data {
		items = append(items, item)
	}
	return items
}

// Size возвращает размер множества
func (s *Set) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.data)
}
