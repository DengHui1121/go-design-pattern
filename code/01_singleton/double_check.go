package _1_singleton

import "sync"

var (
	double *Singleton
	once   = &sync.Once{}
)

func GetDoubleCheck() *Singleton {
	once.Do(func() {
		double = &Singleton{}
	})
	return double
}
