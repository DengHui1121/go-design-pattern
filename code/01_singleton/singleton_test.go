package _1_singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSingleton(t *testing.T) {
	assert.Equal(t, GetSingleton(), GetSingleton())
}

func BenchmarkGetSingletonParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetSingleton() != GetSingleton() {
				b.Errorf("test fail")
			}
		}
	})
}
