package _1_singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDoubleCheck(t *testing.T) {
	assert.Equalf(t, GetDoubleCheck(), GetDoubleCheck(), "GetDoubleCheck fail")
}

func BenchmarkGetDoubleCheck(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetDoubleCheck() != GetDoubleCheck() {
				b.Errorf("test fail")
			}
		}
	})
}
