package day24

import "testing"

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)

	actual := cache.Get(1)
	expected := 1
	if actual != expected {
		t.Errorf("Expected %d but was %d\n", expected, actual)
	}

	cache.Put(3, 3) // evicts key 2
	actual = cache.Get(2)
	expected = -1
	if actual != expected {
		t.Errorf("Expected %d but was %d\n", expected, actual)
	}

	cache.Put(4, 4) // evicts key 1
	actual = cache.Get(1)
	expected = -1
	if actual != expected {
		t.Errorf("Expected %d but was %d\n", expected, actual)
	}

	actual = cache.Get(3)
	expected = 3
	if actual != expected {
		t.Errorf("Expected %d but was %d\n", expected, actual)
	}

	actual = cache.Get(4)
	expected = 4
	if actual != expected {
		t.Errorf("Expected %d but was %d\n", expected, actual)
	}

	// scenario 2
	// cache := Constructor(2)
	// cache.Put(2, 1)
	// cache.Put(1, 1)
	// cache.Put(2, 3)
	// cache.Put(4, 1)
	// actual := cache.Get(1)
	// expected := -1
	// if actual != expected {
	// 	t.Errorf("Expected %d but was %d\n", expected, actual)
	// }
	// actual = cache.Get(2)
	// expected = 3
	// if actual != expected {
	// 	t.Errorf("Expected %d but was %d\n", expected, actual)
	// }
}
