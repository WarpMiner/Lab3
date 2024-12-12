package hashtable

import (
	"testing"
)

func TestInsertAndGet(t *testing.T) {
	ht := NewHashTable()

	ht.Insert("key1", "value1")
	value, found := ht.Get("key1")
	if !found || value != "value1" {
		t.Errorf("Expected value1, got %v", value)
	}

	ht.Insert("key2", "value2")
	value, found = ht.Get("key2")
	if !found || value != "value2" {
		t.Errorf("Expected value2, got %v", value)
	}
}

func TestUpdateValue(t *testing.T) {
	ht := NewHashTable()

	ht.Insert("key1", "value1")
	ht.Insert("key1", "value2") // Обновляем значение
	value, found := ht.Get("key1")
	if !found || value != "value2" {
		t.Errorf("Expected value2 after update, got %v", value)
	}
}

func TestRemove(t *testing.T) {
	ht := NewHashTable()

	ht.Insert("1234526", "value1")
	ht.Insert("54321", "value2")
	if !ht.Remove("54321") {
		t.Error("Expected to remove 54321")
	}
	if _, found := ht.Get("54321"); found {
		t.Error("Expected 54321 to be removed")
	}
}

func TestRemoveNonExistentKey(t *testing.T) {
	ht := NewHashTable()
	if ht.Remove("nonexistent") {
		t.Error("Expected not to remove nonexistent key")
	}
}

func TestCollisionResolution(t *testing.T) {
	ht := NewHashTable()

	// Вставляем два ключа, которые должны хешироваться в один и тот же индекс
	ht.Insert("1234526", "value1")
	ht.Insert("54321", "value2") // Предполагаем, что эти ключи будут коллизией

	value, found := ht.Get("1234526")
	if !found || value != "value1" {
		t.Errorf("Expected value1 for 1234526, got %v", value)
	}

	value, found = ht.Get("54321")
	if !found || value != "value2" {
		t.Errorf("Expected value2 for 54321, got %v", value)
	}
}