package _1_singleton

import "sync"

var (
	lazySingleton *Singleton
	mu            sync.Mutex
)

func newLazySingleton() *Singleton {
	return &Singleton{}
}

func GetLazySingleton() *Singleton {
	mu.Lock()
	if lazySingleton == nil {
		lazySingleton = newLazySingleton()
	}
	mu.Unlock()
	return lazySingleton
}
