// Package server ...
package server

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"

	utils "github.com/praveenmahasena/gocacheserver/internal/util"
)

// Server holds all that needs to start radis...
type Server struct {
	port string
	ln   net.Listener
}

// NewServer to init server
func NewServer(port string) *Server {
	return &Server{port, nil}
}

// Listen func to start the server
func (s *Server) Listen() error {
	ln, lnErr := net.Listen("tcp", s.port)
	if lnErr != nil {
		return fmt.Errorf("error during spwaning server on port %v with error value %#v", s.port, lnErr)
	}
	s.ln = ln
	return nil
}

// Accept to accept connection requests
func (s *Server) Accept(ctx context.Context, messageCh chan utils.Node, wg *sync.WaitGroup) {
	// defer wg.Done()
	// defer fmt.Println("accept")
	// defer s.ln.Close()
	// defer close(messageCh)

	for !errors.Is(ctx.Err(), context.Canceled) {
		con, conErr := s.ln.Accept()
		if conErr != nil {
			continue
		}
		go handleDB(ctx, con, messageCh)
	}
}

func handleDB(ctx context.Context, con net.Conn, messageCh chan utils.Node) {
	buff := bufio.NewScanner(con)
	defer con.Close()
	defer fmt.Println("handleDB")
	for buff.Scan() && (!errors.Is(ctx.Err(), context.Canceled) || buff.Err() != nil) {
		text := buff.Text()
		if text == "exit" {
			continue
		}
		messageCh <- utils.NewNode(con, text)
	}
}

// Response job
func Response(responseCh chan utils.Node, wg *sync.WaitGroup) {
	defer fmt.Println("Response")
	defer wg.Done()
	for node := range responseCh {
		io.WriteString(node.Conn, node.Cmd)
	}
}
