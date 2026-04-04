package cache

import (
	"testing"
	"time"
)

func TestMemoryCache_SetAndGet(t *testing.T) {
	cache := NewMemoryCache(5 * time.Second)

	// Тест сохранения и получения
	key := "test:key"
	value := 25.5

	cache.Set(key, value, 0)

	// Проверяем, что значение сохранилось
	got, found := cache.Get(key)
	if !found {
		t.Errorf("Get() found = false, ожидалось true")
	}
	if got != value {
		t.Errorf("Get() got = %v, ожидалось %v", got, value)
	}
}

func TestMemoryCache_GetNotFound(t *testing.T) {
	cache := NewMemoryCache(5 * time.Second)

	// Тест получения несуществующего ключа
	_, found := cache.Get("nonexistent:key")
	if found {
		t.Errorf("Get() found = true для несуществующего ключа, ожидалось false")
	}
}

func TestMemoryCache_Delete(t *testing.T) {
	cache := NewMemoryCache(5 * time.Second)

	key := "test:delete"
	cache.Set(key, 100, 0)
	cache.Delete(key)

	_, found := cache.Get(key)
	if found {
		t.Errorf("После Delete() ключ всё еще существует")
	}
}

func TestMemoryCache_Clear(t *testing.T) {
	cache := NewMemoryCache(5 * time.Second)

	cache.Set("key1", 1, 0)
	cache.Set("key2", 2, 0)
	cache.Set("key3", 3, 0)

	cache.Clear()

	_, found1 := cache.Get("key1")
	_, found2 := cache.Get("key2")
	_, found3 := cache.Get("key3")

	if found1 || found2 || found3 {
		t.Errorf("После Clear() ключи всё еще существуют")
	}
}

func TestMemoryCache_Expiration(t *testing.T) {
	cache := NewMemoryCache(1 * time.Second)

	key := "test:expire"
	cache.Set(key, 50, 1*time.Second)

	// Сразу после сохранения должно быть
	_, found := cache.Get(key)
	if !found {
		t.Errorf("Сразу после Set() ключ не найден")
	}

	// Ждем истечения TTL
	time.Sleep(1100 * time.Millisecond)

	// После истечения не должно быть
	_, found = cache.Get(key)
	if found {
		t.Errorf("После истечения TTL ключ всё еще существует")
	}
}