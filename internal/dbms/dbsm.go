// Package dbms ...
package dbms

import (
	"strings"
	"sync"

	utils "github.com/praveenmahasena/gocacheserver/internal/util"
)

// DBMS ...
type DBMS struct {
	*Strings
	*listCollection
	*sets
	*hash
}

// New ...
func New() *DBMS {
	return &DBMS{newStr(), newList(), newSets(), newHash()}
}

// Job ...
func (d *DBMS) Job(messageCh chan utils.Node, responseCh chan utils.Node, wg *sync.WaitGroup) {
	// this is not effective and I'm not happy with all
	// these looping happening but I wanna keep this simple
	// there's a book about building compilers that would give me a good idea of all these
	// I grew up doing leetcode and this is a disaster i've coded yet
	defer wg.Done()
	defer close(responseCh)
	for node := range messageCh {
		switch {
		case strings.HasPrefix(node.Cmd, "SET"):
			val := d.set(node.Cmd)
			responseCh <- utils.NewNode(node.Conn, val)
		case strings.HasPrefix(node.Cmd, "GET"):
			val := d.get(node.Cmd)
			responseCh <- utils.NewNode(node.Conn, val)
		case strings.HasPrefix(node.Cmd, "APPEND"):
			val := d.append(node.Cmd)
			responseCh <- utils.NewNode(node.Conn, val)
		case strings.HasPrefix(node.Cmd, "STRLEN"):
			val := d.strlen(node.Cmd)
			responseCh <- utils.NewNode(node.Conn, val)
		case strings.HasPrefix(node.Cmd, "INCR"):
			val := d.incr(node.Cmd)
			responseCh <- utils.NewNode(node.Conn, val)
		case strings.HasPrefix(node.Cmd, "DECR"):
			val := d.decr(node.Cmd)
			responseCh <- utils.NewNode(node.Conn, val)
		case strings.HasPrefix(node.Cmd, "INCRBY"):
			val := d.incrby(node.Cmd)
			responseCh <- utils.NewNode(node.Conn, val)
		case strings.HasPrefix(node.Cmd, "SETNX"):
			val := d.setnx(node.Cmd)
			responseCh <- utils.NewNode(node.Conn, val)
		case strings.HasPrefix(node.Cmd, "GETSET"):
			val := d.getset(node.Cmd)
			responseCh <- utils.NewNode(node.Conn, val)
		case strings.HasPrefix(node.Cmd, "LPUSH"):
			val := d.lpush(node.Cmd)
			responseCh <- utils.NewNode(node.Conn, val)
		case strings.HasPrefix(node.Cmd, "RPUSH"):
			val := d.rpush(node.Cmd)
			responseCh <- utils.NewNode(node.Conn, val)
		case strings.HasPrefix(node.Cmd, "LPOP"):
			val := d.lpop(node.Cmd)
			responseCh <- utils.NewNode(node.Conn, val)
		case strings.HasPrefix(node.Cmd, "RPOP"):
			val := d.rpop(node.Cmd)
			responseCh <- utils.NewNode(node.Conn, val)
		case strings.HasPrefix(node.Cmd, "LPUSHX"):
			val := d.lpushx(node.Cmd)
			responseCh <- utils.NewNode(node.Conn, val)
		case strings.HasPrefix(node.Cmd, "RPUSHX"):
			val := d.rpushx(node.Cmd)
			responseCh <- utils.NewNode(node.Conn, val)
		case strings.HasPrefix(node.Cmd, "LRANGE"):
			val := d.lrange(node.Cmd)
			responseCh <- utils.NewNode(node.Conn, val)
		case strings.HasPrefix(node.Cmd, "LLEN"):
			val := d.llen(node.Cmd)
			responseCh <- utils.NewNode(node.Conn, val)
		case strings.HasPrefix(node.Cmd, "LINDEX"):
			val := d.lindex(node.Cmd)
			responseCh <- utils.NewNode(node.Conn, val)
		case strings.HasPrefix(node.Cmd, "SADD"):
			val := d.sadd(node.Cmd)
			responseCh <- utils.NewNode(node.Conn, val)
		case strings.HasPrefix(node.Cmd, "SREM"):
			val := d.srem(node.Cmd)
			responseCh <- utils.NewNode(node.Conn, val)
		case strings.HasPrefix(node.Cmd, "SMEMBERS"):
			val := d.smembers()
			responseCh <- utils.NewNode(node.Conn, val)
		case strings.HasPrefix(node.Cmd, "SISMEMBER"):
			val := d.sismember(node.Cmd)
			responseCh <- utils.NewNode(node.Conn, val)
		case strings.HasPrefix(node.Cmd, "SCARD"):
			val := d.scard()
			responseCh <- utils.NewNode(node.Conn, val)
		case strings.HasPrefix(node.Cmd, "HSET"):
			val := d.hset(node.Cmd)
			responseCh <- utils.NewNode(node.Conn, val)
		case strings.HasPrefix(node.Cmd, "HGET"):
			val := d.hget(node.Cmd)
			responseCh <- utils.NewNode(node.Conn, val)
		case strings.HasPrefix(node.Cmd, "HDEL"):
			val := d.hdel(node.Cmd)
			responseCh <- utils.NewNode(node.Conn, val)
		case strings.HasPrefix(node.Cmd, "HGETALL"):
			val := d.hgetall(node.Cmd)
			responseCh <- utils.NewNode(node.Conn, val)
		case strings.HasPrefix(node.Cmd, "HEXISTS"):
			val := d.hexists(node.Cmd)
			responseCh <- utils.NewNode(node.Conn, val)
		}
	}
}
