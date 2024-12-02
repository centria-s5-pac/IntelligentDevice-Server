package lightbrain

import (
	"sync"
)

var (
	sharedVar int
	rwMutex   sync.RWMutex
)

func initValue() {
	sharedVar = 0
}

func SetValue(v int) {
	rwMutex.Lock()
	sharedVar = v
	rwMutex.Unlock()
}

func GetValue() int {
	rwMutex.RLock()
	value := sharedVar
	rwMutex.RUnlock()
	return value
}
