// Package dbms ...
package dbms

import (
	"fmt"
	"strings"
	"sync"
)

// am doing graph stuff now?
type hashMap map[string]map[string]string

type hash struct {
	hashMap
	*sync.RWMutex
}

func newHash() *hash {
	return &hash{make(hashMap), &sync.RWMutex{}}
}

func (h *hash) hset(cmd string) string {
	h.Lock()
	defer h.Unlock()
	c := strings.Split(cmd, " ")
	if h.hashMap[c[1]] == nil {
		h.hashMap[c[1]] = make(map[string]string)
	}
	h.hashMap[c[1]][c[2]] = c[3]
	return c[2]
}

func (h *hash) hget(cmd string) string {
	h.RLock()
	defer h.RUnlock()
	c := strings.Split(cmd, " ")
	if h.hashMap[c[1]] == nil {
		return ""
	}
	return h.hashMap[c[1]][c[2]]
}

func (h *hash) hdel(cmd string) string {
	h.Lock()
	defer h.Unlock()
	c := strings.Split(cmd, " ")

	if h.hashMap[c[1]] == nil {
		return ""
	}
	val := h.hashMap[c[1]][c[2]]
	delete(h.hashMap[c[1]], c[2])
	return val
}

func (h *hash) hgetall(cmd string) string {
	h.RLock()
	defer h.Unlock()
	c := strings.Split(cmd, " ")
	str := strings.Builder{}
	for _, elem := range h.hashMap[c[2]] {
		str.WriteString(elem)
	}
	return str.String()
}

func (h *hash) hexists(cmd string) string {
	h.RLock()
	defer h.RUnlock()
	c := strings.Split(cmd, " ")
	_, ok := h.hashMap[c[1]][c[2]]
	return fmt.Sprint(ok)
}
