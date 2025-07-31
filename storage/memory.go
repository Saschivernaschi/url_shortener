// storage/memory.go
package storage

import "sync"

var urls = make(map[string]string)
var mu sync.RWMutex

func Save(code, url string) {
	mu.Lock()
	defer mu.Unlock()
	urls[code] = url
}

func Get(code string) (string, bool) {
	mu.RLock()
	defer mu.RUnlock()
	url, ok := urls[code]
	return url, ok
}
