package singleton

import (
	"fmt"
	"sync"
)

type singleton struct {
}

var (
	instance *singleton
	once     sync.Once
	mutex    sync.Mutex
)

// Race condition may occur
func GetInstance() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

func GetInstanceWithSyncOnce() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func GetInstanceWithMutex() *singleton {
	mutex.Lock()
	defer mutex.Unlock()
	if instance == nil {
		fmt.Println("init singleton instance")
		instance = &singleton{}
	}
	return instance
}

func GetInstanceWithMutexDoubleCheck() *singleton {
	if instance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if instance == nil {
			fmt.Println("init singleton instance")
			instance = &singleton{}
		}
	}
	return instance
}
