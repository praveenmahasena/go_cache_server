// Package dbms ...
package dbms

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

// List ...
type List struct {
	head, tail *Node
	leng       uint
}

// Node ...
type Node struct {
	prev, next *Node
	val        string
}

type listCollection struct {
	*sync.RWMutex
	list map[string]*List
}

func newList() *listCollection {
	return &listCollection{&sync.RWMutex{}, map[string]*List{}}
}

func (l *listCollection) lpush(cmd string) string {
	l.Lock()
	defer l.Unlock()
	c := strings.Split(cmd, " ")

	if l.list[c[1]] == nil {
		l.list[c[1]] = &List{leng: 1}
	}
	pushHead(l.list[c[1]], c[2])
	return c[2]
}

func (l *listCollection) rpush(cmd string) string {
	l.Lock()
	defer l.Unlock()
	c := strings.Split(cmd, " ")

	if l.list[c[1]] == nil {
		l.list[c[1]] = &List{leng: 1}
	}
	pushTail(l.list[c[1]], c[2])
	return c[2]
}

func (l *listCollection) lpop(cmd string) string {
	l.Lock()
	defer l.Unlock()
	c := strings.Split(cmd, " ")

	if l.list[c[1]] == nil {
		return ""
	}
	l.list[c[1]].leng--
	return headpop(l.list[c[1]])
}

func (l *listCollection) rpop(cmd string) string {
	l.Lock()
	defer l.Unlock()

	c := strings.Split(cmd, " ")

	if l.list[c[1]] == nil {
		return ""
	}
	l.list[c[1]].leng--
	return tailpop(l.list[c[1]])
}

func (l *listCollection) lpushx(cmd string) string {
	l.Lock()
	defer l.Unlock()
	c := strings.Split(cmd, " ")

	if l.list[c[1]] == nil {
		return ""
	}
	l.list[c[1]].leng++
	pushHead(l.list[c[1]], c[2])
	return c[2]
}

func (l *listCollection) rpushx(cmd string) string {
	l.Lock()
	defer l.Unlock()
	c := strings.Split(cmd, " ")

	if l.list[c[1]] == nil {
		return ""
	}
	l.list[c[1]].leng++
	pushTail(l.list[c[1]], c[2])
	return c[2]
}

func (l *listCollection) lrange(cmd string) string {
	l.RLock()
	defer l.RUnlock()
	c := strings.Split(cmd, " ")
	if l.list[c[1]] == nil {
		return ""
	}
	return iter(l.list[c[2]])
}

func (l *listCollection) llen(cmd string) string {
	l.RLock()
	defer l.RUnlock()
	c := strings.Split(cmd, " ")
	if l.list[c[1]] == nil {
		return "0"
	}
	return fmt.Sprint(l.list[c[1]].leng)
}

func (l *listCollection) lindex(cmd string) string {
	l.RLock()
	defer l.RUnlock()

	c := strings.Split(cmd, " ")

	if l.list[c[1]] == nil {
		return ""
	}
	val, err := strconv.Atoi(c[2])
	if err != nil {
		return ""
	}
	if val > int(l.list[c[2]].leng)-1 {
		return ""
	}
	n := l.list[c[1]].head
	for range val {
		n = n.next
	}
	return n.val
}
