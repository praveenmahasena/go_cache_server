// Package utils ...
package utils

import "net"

// Node ...
type Node struct {
	net.Conn
	Cmd string
}

// NewNode ...
func NewNode(con net.Conn, cmd string) Node {
	return Node{con, cmd}
}
