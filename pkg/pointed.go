package pkg

import (
	"sync"
)

var (
	ptrsMutex   = &sync.Mutex{}
	ptrsStorage = make(map[uint64]*Object)
)

// Put ...
func Put(o *Object) *Object {
	ptr := o.Ptr
	if ptr == 0 {
		return nil
	}
	ptrsMutex.Lock()
	defer ptrsMutex.Unlock()
	ptrsStorage[ptr] = o
	return o
}

// Get ...
func Get(ptr uint64) *Object {
	if ptr == 0 {
		return nil
	}
	ptrsMutex.Lock()
	defer ptrsMutex.Unlock()
	return ptrsStorage[ptr]
}
