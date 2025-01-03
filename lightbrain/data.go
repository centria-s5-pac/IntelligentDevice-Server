package lightbrain

import (
	"sync"
)

var (
	sharedVar  int
	sharedMode int
	rwMutex    sync.RWMutex
)

func initValue() {
	sharedVar = 0
	sharedMode = 0
}

func SetValue(v int) {
	rwMutex.Lock()
	sharedVar = v
	rwMutex.Unlock()
}

func SetMode(v int) {
	rwMutex.Lock()
	sharedMode = v
	rwMutex.Unlock()
}

func GetValue() int {
	rwMutex.RLock()
	value := sharedVar
	rwMutex.RUnlock()
	return value
}

func GetMode() int {
	rwMutex.RLock()
	value := sharedMode
	rwMutex.RUnlock()
	return value
}
