package singleton

import "sync"

var lazySingleton *Singleton
var once *sync.Once = &sync.Once{}

func GetLazyInstance() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}
	return lazySingleton
}
