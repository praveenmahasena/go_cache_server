// Package dbms ...
package dbms

import (
	"fmt"
	"strings"
	"sync"
)

type setMap map[string]struct{}

type sets struct {
	setMap
	*sync.RWMutex
}

func newSets() *sets {
	return &sets{make(setMap), &sync.RWMutex{}}
}

func (s *sets) sadd(cmd string) string {
	s.Lock()
	defer s.Unlock()
	c := strings.Split(cmd, " ")
	for _, elem := range c[1:] {
		s.setMap[elem] = struct{}{}
	}
	return fmt.Sprint(c[1:])
}

func (s *sets) srem(cmd string) string {
	s.Lock()
	defer s.Unlock()
	c := strings.Split(cmd, " ")
	for _, elem := range c[1:] {
		delete(s.setMap, elem)
	}
	return fmt.Sprint(c[1:])
}

func (s *sets) smembers() string {
	s.RLock()
	defer s.RUnlock()
	if len(s.setMap) == 0 {
		return ""
	}
	str := make([]string, len(s.setMap))
	for elem := range s.setMap {
		str = append(str, elem)
	}
	return fmt.Sprint(str)
}

func (s *sets) sismember(cmd string) string {
	s.RLock()
	defer s.RUnlock()
	if len(s.setMap) == 0 {
		return ""
	}
	c := strings.Split(cmd, " ")
	memberStr := strings.Builder{}
	for _, elem := range c[1:] {
		if _, ok := s.setMap[elem]; ok {
			memberStr.WriteString("1 ")
			continue
		}
		memberStr.WriteString("0 ")
	}
	return memberStr.String()
}

func (s *sets) scard() string {
	s.RLock()
	defer s.RUnlock()
	return fmt.Sprint(len(s.setMap))
}
