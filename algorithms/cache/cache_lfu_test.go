package cache

import (
	"testing"
)

func TestCacheLFU(t *testing.T) {
	t.Run("Should return nil for none stored keys", func(t *testing.T) {
		c := New()
		v := c.Get("a")
		if v != nil {
			t.Fatalf("Cache Value %v; want %v", v, nil)
		}
	})

	t.Run("Should return value for store key", func(t *testing.T) {
		c := New()
		c.Set("a", 1)
		v := c.Get("a")
		if v != 1 {
			t.Fatalf("Cache Value %v; want %v", v, 1)
		}
	})
}
