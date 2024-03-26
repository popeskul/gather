package set

import (
	"reflect"
	"sort"
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		elements     []interface{}
		expectedSize int
	}{
		{[]interface{}{"test1"}, 1},
		{[]interface{}{"test1", "test2"}, 2},
		{[]interface{}{"test1", "test2", "test2"}, 2}, // Проверка на уникальность
	}

	for _, c := range cases {
		s := New()
		for _, el := range c.elements {
			s.Add(el)
		}
		if s.Size() != c.expectedSize {
			t.Errorf("Add(%v): expected size %d, got %d", c.elements, c.expectedSize, s.Size())
		}
	}
}

func TestRemove(t *testing.T) {
	cases := []struct {
		elementsToAdd    []interface{}
		elementsToRemove []interface{}
		expectedSize     int
	}{
		{[]interface{}{"test1"}, []interface{}{"test1"}, 0},
		{[]interface{}{"test1", "test2"}, []interface{}{"test1"}, 1},
		{[]interface{}{"test1", "test2", "test3"}, []interface{}{"test4"}, 3}, // Попытка удалить несуществующий элемент
	}

	for _, c := range cases {
		s := New()
		for _, el := range c.elementsToAdd {
			s.Add(el)
		}
		for _, el := range c.elementsToRemove {
			s.Remove(el)
		}
		if s.Size() != c.expectedSize {
			t.Errorf("Remove(%v): expected size %d, got %d", c.elementsToRemove, c.expectedSize, s.Size())
		}
	}
}

func TestExists(t *testing.T) {
	cases := []struct {
		elements []interface{}
		check    interface{}
		expected bool
	}{
		{[]interface{}{"test1"}, "test1", true},
		{[]interface{}{"test1", "test2"}, "test3", false},
	}

	for _, c := range cases {
		s := New()
		for _, el := range c.elements {
			s.Add(el)
		}
		if s.Exists(c.check) != c.expected {
			t.Errorf("Exists(%v): expected %t, got %t", c.check, c.expected, s.Exists(c.check))
		}
	}
}

func TestList(t *testing.T) {
	cases := []struct {
		elements []interface{}
		expected []interface{}
	}{
		{[]interface{}{"test1", "test2"}, []interface{}{"test1", "test2"}},
		{[]interface{}{"test2", "test1"}, []interface{}{"test1", "test2"}}, // Проверка на порядок
	}

	for _, c := range cases {
		s := New()
		for _, el := range c.elements {
			s.Add(el)
		}
		got := s.List()
		sort.Slice(got, func(i, j int) bool { return got[i].(string) < got[j].(string) })
		if !reflect.DeepEqual(got, c.expected) {
			t.Errorf("List(): expected %v, got %v", c.expected, got)
		}
	}
}

func TestSize(t *testing.T) {
	cases := []struct {
		elements []interface{}
		expected int
	}{
		{[]interface{}{}, 0},
		{[]interface{}{"test1"}, 1},
		{[]interface{}{"test1", "test2"}, 2},
		{[]interface{}{"test1", "test2", "test2"}, 2}, // Повторяющиеся элементы не считаются
	}

	for _, c := range cases {
		s := New()
		for _, el := range c.elements {
			s.Add(el)
		}
		if s.Size() != c.expected {
			t.Errorf("Size(): expected %d, got %d", c.expected, s.Size())
		}
	}
}
