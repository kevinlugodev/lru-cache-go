package testing

import (
	lrucache "lru-cache-golang/library"
	"testing"
	"time"
)

func TestLRUCache(t *testing.T) {
	capacity := uint(3)
	cache := lrucache.NewLRUCache(&capacity)

	// duration := 1 * time.Second
	cache.Set("key1", "value1", nil)
	cache.Set("key2", "value2", nil)
	cache.Set("key3", "value3", nil)

	if cache.Get("key1") != "value1" {
		t.Errorf("Valor esperado para key1: value1, Valor obtenido: %v", cache.Get("key1"))
	}
	if cache.Get("key2") != "value2" {
		t.Errorf("Valor esperado para key2: value2, Valor obtenido: %v", cache.Get("key2"))
	}
	if cache.Get("key3") != "value3" {
		t.Errorf("Valor esperado para key3: value3, Valor obtenido: %v", cache.Get("key3"))
	}

	time.Sleep(2 * time.Second)

	if cache.Get("key1") != nil {
		t.Errorf("Se esperaba que key1 haya expirado, pero se encontró un valor")
	}
	if cache.Get("key2") != nil {
		t.Errorf("Se esperaba que key2 haya expirado, pero se encontró un valor")
	}
	if cache.Get("key3") != nil {
		t.Errorf("Se esperaba que key3 haya expirado, pero se encontró un valor")
	}
}
