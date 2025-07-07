// Package dbms ...
package dbms

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

// Strings ...
type Strings struct {
	strings map[string]string
	*sync.RWMutex
}

func newStr() *Strings {
	return &Strings{map[string]string{}, &sync.RWMutex{}}
}

func (s *Strings) set(cmd string) string {
	c := strings.Split(cmd, " ")
	if len(c) < 3 {
		return "no update"
	}
	s.Lock()
	defer s.Unlock()
	s.strings[c[1]] = c[2]
	return c[2]
}

func (s *Strings) get(cmd string) string {
	c := strings.Split(cmd, " ")
	if len(c) < 3 {
		return "no value"
	}
	s.RLock()
	defer s.RUnlock()
	return s.strings[c[1]]
}
func (s *Strings) append(cmd string) string {
	c := strings.Split(cmd, " ")
	if len(c) < 3 {
		return "no value to update"
	}
	s.Lock()
	defer s.Unlock()
	s.strings[c[1]] += c[2]
	return s.strings[c[1]]
}
func (s *Strings) strlen(cmd string) string {
	c := strings.Split(cmd, " ")
	if len(c) < 3 {
		return "0"
	}
	s.RLock()
	defer s.RUnlock()
	return fmt.Sprint(len(s.strings[c[1]]))
}
func (s *Strings) incr(cmd string) string {
	c := strings.Split(cmd, " ")
	if len(c) < 3 {
		return "no update"
	}
	s.Lock()
	defer s.Unlock()
	val, err := strconv.Atoi(s.strings[c[1]])
	if err != nil {
		return "cannot update"
	}
	s.strings[c[1]] = fmt.Sprintf("%v", val+1)
	return s.strings[c[1]]
}

func (s *Strings) decr(cmd string) string {
	c := strings.Split(cmd, " ")
	if len(c) < 3 {
		return "no update"
	}
	s.Lock()
	defer s.Unlock()
	val, err := strconv.Atoi(s.strings[c[1]])
	if err != nil {
		return "cannot update"
	}
	s.strings[c[1]] = fmt.Sprintf("%v", val-1)
	return s.strings[c[1]]
}

func (s *Strings) incrby(cmd string) string {
	c := strings.Split(cmd, " ")
	if len(c) < 3 {
		return "no update"
	}
	s.Lock()
	defer s.Unlock()
	inc, err := strconv.Atoi(c[2])
	if err != nil {
		return "cannot update"
	}
	val, err := strconv.Atoi(s.strings[c[1]])
	if err != nil {
		return "cannot update"
	}
	s.strings[c[1]] = fmt.Sprintf("%v", val+inc)
	return s.strings[c[1]]
}

func (s *Strings) setnx(cmd string) string {
	c := strings.Split(cmd, " ")
	if len(c) < 3 {
		return "no update"
	}
	s.Lock()
	defer s.Unlock()
	if _, ok := s.strings[c[1]]; !ok {
		s.strings[c[1]] = c[2]
	}
	return "no update"
}
func (s *Strings) getset(cmd string) (val string) {
	c := strings.Split(cmd, " ")
	if len(c) < 3 {
		return "no update"
	}
	s.Lock()
	defer s.Unlock()
	defer func() {
		val = s.strings[c[1]]
	}()
	s.strings[c[1]] = c[2]

	return val
}
