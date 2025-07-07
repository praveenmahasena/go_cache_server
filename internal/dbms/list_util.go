// Package dbms ...
package dbms

import "strings"

func pushHead(l *List, val string) {
	n := &Node{val: val}
	if l.head == nil {
		l.head, l.tail = n, n
		l.leng++
		return
	}
	l.head.prev, n.next = n, l.head
	l.head = n
}

func pushTail(l *List, val string) {
	n := &Node{val: val}
	if l.head == nil {
		l.head, l.tail = n, n
		l.leng++
		return
	}
	l.tail.next, n.prev = n, l.tail
	l.tail = n
}

func headpop(l *List) string {
	t := l.head

	l.head = l.head.next
	l.head.prev = nil
	t.next = nil

	return t.val
}

func tailpop(l *List) string {
	t := l.tail
	l.tail = l.tail.prev
	l.tail.next = nil
	t.prev = nil
	return t.val
}

func iter(l *List) string {
	listStr := strings.Builder{}
	n := l.head
	for n.next != nil {
		listStr.WriteString(n.val)
		n = n.next
	}
	return listStr.String()
}
